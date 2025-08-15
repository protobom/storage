// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

//go:build integration

package ent_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/reader"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"google.golang.org/protobuf/proto"

	"github.com/protobom/storage/backends/ent"
)

// Backend Integration Test Suite
// Tests backend functionality with PostgreSQL database
type backendIntegrationSuite struct {
	suite.Suite
	backend   *ent.Backend
	container testcontainers.Container
	documents []*sbom.Document
}

func (bis *backendIntegrationSuite) SetupSuite() {
	ctx := context.Background()

	// Start PostgreSQL container
	pgContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpass"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second)),
	)
	bis.Require().NoError(err)
	bis.container = pgContainer

	// Get connection string
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	bis.Require().NoError(err)

	// Initialize backend
	bis.backend = ent.NewBackend(
		ent.WithPostgresConnection(connStr),
	)
	bis.Require().NoError(bis.backend.InitClient())

	// Load test documents
	cwd, err := os.Getwd()
	bis.Require().NoError(err)

	rdr := reader.New()
	testdataDir := filepath.Join(cwd, "testdata")

	entries, err := os.ReadDir(testdataDir)
	bis.Require().NoError(err)

	for idx := range entries {
		document, err := rdr.ParseFile(filepath.Join(testdataDir, entries[idx].Name()))
		bis.Require().NoError(err)
		bis.documents = append(bis.documents, document)
	}
}

func (bis *backendIntegrationSuite) TearDownSuite() {
	if bis.backend != nil {
		bis.backend.CloseClient()
	}
	if bis.container != nil {
		ctx := context.Background()
		bis.Require().NoError(bis.container.Terminate(ctx))
	}
}

func (bis *backendIntegrationSuite) TestBackendConfiguration() {
	ctx := context.Background()
	connStr, err := bis.container.(*postgres.PostgresContainer).ConnectionString(ctx, "sslmode=disable")
	bis.Require().NoError(err)

	tests := []struct {
		name            string
		backendSetup    func() *ent.Backend
		expectedDialect ent.DatabaseDialect
		expectedURL     string
		shouldInit      bool
	}{
		{
			name: "postgres_connection_helper",
			backendSetup: func() *ent.Backend {
				return ent.NewBackend(ent.WithPostgresConnection(connStr))
			},
			expectedDialect: ent.PostgresDialect,
			expectedURL:     connStr,
			shouldInit:      true,
		},
		{
			name: "separate_options",
			backendSetup: func() *ent.Backend {
				return ent.NewBackend(
					ent.WithDialect(ent.PostgresDialect),
					ent.WithDatabaseURL(connStr),
				)
			},
			expectedDialect: ent.PostgresDialect,
			expectedURL:     connStr,
			shouldInit:      true,
		},
		{
			name: "debug_enabled",
			backendSetup: func() *ent.Backend {
				return ent.NewBackend(
					ent.WithPostgresConnection(connStr),
					ent.Debug(),
				)
			},
			expectedDialect: ent.PostgresDialect,
			expectedURL:     connStr,
			shouldInit:      true,
		},
	}

	for _, tt := range tests {
		bis.T().Run(tt.name, func(t *testing.T) {
			backend := tt.backendSetup()

			// Validate configuration
			bis.Equal(tt.expectedDialect, backend.Options.Dialect)
			bis.Equal(tt.expectedURL, backend.Options.DatabaseURL)

			// Test initialization if required
			if tt.shouldInit {
				err := backend.InitClient()
				bis.Require().NoError(err)
				backend.CloseClient()
			}
		})
	}
}

func (bis *backendIntegrationSuite) TestBackendStoreAndRetrieve() {
	tests := []struct {
		name          string
		document      *sbom.Document
		expectError   bool
		validateRetry bool
	}{}

	// Build test cases from loaded documents
	for idx, doc := range bis.documents {
		documentID := doc.GetMetadata().GetId()
		testName := documentID
		if testName == "" {
			testName = fmt.Sprintf("document_%d", idx)
		}

		tests = append(tests, struct {
			name          string
			document      *sbom.Document
			expectError   bool
			validateRetry bool
		}{
			name:          testName,
			document:      doc,
			expectError:   false,
			validateRetry: true,
		})
	}

	for _, tt := range tests {
		bis.T().Run(tt.name, func(t *testing.T) {
			// Test initial storage
			err := bis.backend.Store(tt.document, nil)
			if tt.expectError {
				bis.Require().Error(err)
				return
			}
			bis.Require().NoError(err)

			// Get retrieval ID
			retrievalID := tt.document.GetMetadata().GetId()
			if retrievalID == "" {
				generatedUUID, err := ent.GenerateUUID(tt.document.GetMetadata())
				bis.Require().NoError(err)
				retrievalID = generatedUUID.String()
			}

			// Test retrieval
			retrieved, err := bis.backend.Retrieve(retrievalID, nil)
			bis.Require().NoError(err)

			// Validate content equality
			retrievedCopy := proto.Clone(retrieved).(*sbom.Document)
			originalCopy := proto.Clone(tt.document).(*sbom.Document)
			retrievedCopy.GetMetadata().GetSourceData().Uri = nil
			originalCopy.GetMetadata().GetSourceData().Uri = nil

			bis.Require().True(proto.Equal(originalCopy, retrievedCopy))

			// Test duplicate storage (should not error)
			if tt.validateRetry {
				err = bis.backend.Store(tt.document, nil)
				bis.Require().NoError(err, "Duplicate storage should not fail")
			}
		})
	}
}

func (bis *backendIntegrationSuite) TestBackendClientManagement() {
	ctx := context.Background()
	connStr, err := bis.container.(*postgres.PostgresContainer).ConnectionString(ctx, "sslmode=disable")
	bis.Require().NoError(err)

	tests := []struct {
		name        string
		setupFunc   func() *ent.Backend
		expectError bool
	}{
		{
			name: "valid_initialization",
			setupFunc: func() *ent.Backend {
				return ent.NewBackend(ent.WithPostgresConnection(connStr))
			},
			expectError: false,
		},
		{
			name: "invalid_connection_string",
			setupFunc: func() *ent.Backend {
				return ent.NewBackend(ent.WithPostgresConnection("invalid-connection"))
			},
			expectError: true,
		},
		{
			name: "double_initialization",
			setupFunc: func() *ent.Backend {
				backend := ent.NewBackend(ent.WithPostgresConnection(connStr))
				// Initialize once
				err := backend.InitClient()
				bis.Require().NoError(err)
				return backend
			},
			expectError: false, // Should not error on double init
		},
	}

	for _, tt := range tests {
		bis.T().Run(tt.name, func(t *testing.T) {
			backend := tt.setupFunc()

			err := backend.InitClient()
			if tt.expectError {
				bis.Require().Error(err)
			} else {
				bis.Require().NoError(err)
				backend.CloseClient()
			}
		})
	}
}

func (bis *backendIntegrationSuite) TestConcurrentBackendOperations() {
	// Test concurrent backend operations using the same backend instance
	tests := []struct {
		name            string
		concurrentCount int
		testFunc        func(index int, backend *ent.Backend) error
		expectErrors    bool
		description     string
	}{
		{
			name:            "concurrent_backend_store_operations",
			concurrentCount: 8,
			testFunc: func(index int, backend *ent.Backend) error {
				testDoc := bis.documents[0]
				doc := proto.Clone(testDoc).(*sbom.Document)
				doc.Metadata.Id = fmt.Sprintf("concurrent-backend-%d-%s", index, uuid.New().String())
				return backend.Store(doc, nil)
			},
			expectErrors: false,
			description:  "Concurrent backend store operations should succeed",
		},
		{
			name:            "concurrent_backend_retrieve_operations",
			concurrentCount: 10,
			testFunc: func(index int, backend *ent.Backend) error {
				// Small delay to ensure document is stored before retrieval attempts
				time.Sleep(50 * time.Millisecond)

				_, err := backend.Retrieve("shared-retrieval-doc", nil)
				return err
			},
			expectErrors: false,
			description:  "Concurrent backend retrieve operations should succeed",
		},
	}

	for _, tt := range tests {
		bis.T().Run(tt.name, func(t *testing.T) {
			// Pre-store the document for retrieval tests
			if tt.name == "concurrent_backend_retrieve_operations" {
				testDoc := bis.documents[0]
				doc := proto.Clone(testDoc).(*sbom.Document)
				doc.Metadata.Id = "shared-retrieval-doc"
				err := bis.backend.Store(doc, nil)
				bis.Require().NoError(err)
			}

			var wg sync.WaitGroup
			resultsChan := make(chan error, tt.concurrentCount)

			// Use the shared backend instance for concurrent operations
			for i := 0; i < tt.concurrentCount; i++ {
				wg.Add(1)
				go func(index int) {
					defer wg.Done()

					// Execute the test function with the shared backend
					err := tt.testFunc(index, bis.backend)
					resultsChan <- err
				}(i)
			}

			wg.Wait()
			close(resultsChan)

			// Collect results
			var errorCount int
			for err := range resultsChan {
				if err != nil {
					errorCount++
					bis.T().Logf("Concurrent backend operation error: %v", err)
				}
			}

			if tt.expectErrors {
				bis.Greater(errorCount, 0, "Expected some errors")
			} else {
				bis.Equal(0, errorCount, "Expected no errors in concurrent backend operations")
			}
		})
	}
}

func TestBackendIntegrationSuite(t *testing.T) {
	suite.Run(t, new(backendIntegrationSuite))
}
