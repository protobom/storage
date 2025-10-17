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

// Store Integration Test Suite
// Tests complex storage scenarios with PostgreSQL database
type storeIntegrationSuite struct {
	suite.Suite
	backend   *ent.Backend
	container testcontainers.Container
	documents []*sbom.Document
}

func (sis *storeIntegrationSuite) SetupSuite() {
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
	sis.Require().NoError(err)
	sis.container = pgContainer

	// Get connection string
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	sis.Require().NoError(err)

	// Initialize backend
	sis.backend = ent.NewBackend(
		ent.WithPostgresConnection(connStr),
	)
	sis.Require().NoError(sis.backend.InitClient())

	// Load test documents
	cwd, err := os.Getwd()
	sis.Require().NoError(err)

	rdr := reader.New()
	testdataDir := filepath.Join(cwd, "testdata")

	entries, err := os.ReadDir(testdataDir)
	sis.Require().NoError(err)

	for idx := range entries {
		document, err := rdr.ParseFile(filepath.Join(testdataDir, entries[idx].Name()))
		sis.Require().NoError(err)
		sis.documents = append(sis.documents, document)
	}
}

func (sis *storeIntegrationSuite) TearDownSuite() {
	if sis.backend != nil {
		sis.backend.CloseClient()
	}
	if sis.container != nil {
		ctx := context.Background()
		sis.Require().NoError(sis.container.Terminate(ctx))
	}
}

func (sis *storeIntegrationSuite) TestConflictHandling() {
	tests := []struct {
		name            string
		document        *sbom.Document
		conflictCount   int
		expectError     bool
		validateContent bool
	}{
		{
			name:            "single_document_conflicts",
			document:        sis.documents[0],
			conflictCount:   15,
			expectError:     false,
			validateContent: true,
		},
	}

	// Add cross-document conflicts if we have multiple documents
	if len(sis.documents) > 2 {
		tests = append(tests, struct {
			name            string
			document        *sbom.Document
			conflictCount   int
			expectError     bool
			validateContent bool
		}{
			name:            "cross_document_conflicts",
			document:        sis.documents[1],
			conflictCount:   5,
			expectError:     false,
			validateContent: true,
		})
	}

	for _, tt := range tests {
		sis.T().Run(tt.name, func(t *testing.T) {
			// Initial storage
			err := sis.backend.Store(tt.document, nil)
			sis.Require().NoError(err, "Initial storage should succeed")

			// Test multiple conflicts
			for i := 0; i < tt.conflictCount; i++ {
				err = sis.backend.Store(tt.document, nil)
				if tt.expectError {
					sis.Require().Error(err, fmt.Sprintf("Conflict test %d should fail", i+1))
				} else {
					sis.Require().NoError(err, fmt.Sprintf("Conflict test %d should succeed", i+1))
				}
			}

			// Validate content integrity if requested
			if tt.validateContent {
				retrieved, err := sis.backend.Retrieve(tt.document.GetMetadata().GetId(), nil)
				sis.Require().NoError(err)
				sis.Require().NotNil(retrieved)

				// Compare content
				retrievedCopy := proto.Clone(retrieved).(*sbom.Document)
				testDocCopy := proto.Clone(tt.document).(*sbom.Document)
				retrievedCopy.GetMetadata().GetSourceData().Uri = nil
				testDocCopy.GetMetadata().GetSourceData().Uri = nil

				sis.Require().True(proto.Equal(testDocCopy, retrievedCopy), "Retrieved document should match original after conflict resolution")
			}
		})
	}
}

func (sis *storeIntegrationSuite) TestComplexStorageScenarios() {
	tests := []struct {
		name        string
		setupFunc   func() error
		testFunc    func() error
		cleanupFunc func() error
		expectError bool
	}{
		{
			name: "multiple_documents_same_metadata",
			setupFunc: func() error {
				// Store first document
				return sis.backend.Store(sis.documents[0], nil)
			},
			testFunc: func() error {
				// Try to store the same document again
				return sis.backend.Store(sis.documents[0], nil)
			},
			expectError: false, // Should not error on duplicate
		},
		{
			name: "interleaved_storage_and_retrieval",
			setupFunc: func() error {
				return nil // No setup needed
			},
			testFunc: func() error {
				// Store, retrieve, store again pattern
				for i, doc := range sis.documents {
					// Store
					err := sis.backend.Store(doc, nil)
					if err != nil {
						return fmt.Errorf("store %d failed: %w", i, err)
					}

					// Retrieve
					docID := doc.GetMetadata().GetId()
					if docID == "" {
						generatedUUID, err := ent.GenerateUUID(doc.GetMetadata())
						if err != nil {
							return fmt.Errorf("generate UUID %d failed: %w", i, err)
						}
						docID = generatedUUID.String()
					}

					_, err = sis.backend.Retrieve(docID, nil)
					if err != nil {
						return fmt.Errorf("retrieve %d failed: %w", i, err)
					}

					// Store again
					err = sis.backend.Store(doc, nil)
					if err != nil {
						return fmt.Errorf("re-store %d failed: %w", i, err)
					}
				}
				return nil
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		sis.T().Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setupFunc != nil {
				err := tt.setupFunc()
				sis.Require().NoError(err)
			}

			// Test
			if tt.testFunc != nil {
				err := tt.testFunc()
				if tt.expectError {
					sis.Require().Error(err)
				} else {
					sis.Require().NoError(err)
				}
			}

			// Cleanup
			if tt.cleanupFunc != nil {
				err := tt.cleanupFunc()
				sis.Require().NoError(err)
			}
		})
	}
}

func (sis *storeIntegrationSuite) TestEdgeValidation() {
	// Test our edge validation fix - create SBOM with invalid edges
	tests := []struct {
		name        string
		createDoc   func() *sbom.Document
		expectError bool
		description string
	}{
		{
			name: "valid_edges",
			createDoc: func() *sbom.Document {
				return &sbom.Document{
					Metadata: &sbom.Metadata{
						Id:      "test-valid-edges",
						Version: "1",
						Name:    "Test Valid Edges",
					},
					NodeList: &sbom.NodeList{
						Nodes: []*sbom.Node{
							{
								Id:   "node-1",
								Type: sbom.Node_PACKAGE,
								Name: "package-1",
							},
							{
								Id:   "node-2",
								Type: sbom.Node_PACKAGE,
								Name: "package-2",
							},
						},
						Edges: []*sbom.Edge{
							{
								Type: sbom.Edge_dependsOn,
								From: "node-1",
								To:   []string{"node-2"},
							},
						},
					},
				}
			},
			expectError: false,
			description: "Valid edges should be stored successfully",
		},
		{
			name: "invalid_edges_missing_nodes",
			createDoc: func() *sbom.Document {
				return &sbom.Document{
					Metadata: &sbom.Metadata{
						Id:      "test-invalid-edges",
						Version: "1",
						Name:    "Test Invalid Edges",
					},
					NodeList: &sbom.NodeList{
						Nodes: []*sbom.Node{
							{
								Id:   "node-1",
								Type: sbom.Node_PACKAGE,
								Name: "package-1",
							},
						},
						Edges: []*sbom.Edge{{
							Type: sbom.Edge_dependsOn,
							From: "node-1",
							To:   []string{"00000000-0000-0000-0000-000000000000"},
						}},
					},
				}
			},
			expectError: true,
			description: "Invalid edges should error and abort storage due to missing node",
		},
	}

	for _, tt := range tests {
		sis.T().Run(tt.name, func(t *testing.T) {
			doc := tt.createDoc()

			err := sis.backend.Store(doc, nil)
			if tt.expectError {
				sis.Require().Error(err, tt.description)
			} else {
				sis.Require().NoError(err, tt.description)

				// Verify document can be retrieved
				retrieved, err := sis.backend.Retrieve(doc.GetMetadata().GetId(), nil)
				sis.Require().NoError(err)
				sis.Require().NotNil(retrieved)
			}
		})
	}
}

func (sis *storeIntegrationSuite) TestConcurrentStore() {
	// Test concurrent storage operations to validate transaction isolation
	tests := []struct {
		name               string
		concurrentCount    int
		useFixedID         bool
		expectAllSucceed   bool
		expectMinSuccesses int
		description        string
	}{
		{
			name:               "concurrent_different_documents",
			concurrentCount:    10,
			useFixedID:         false,
			expectAllSucceed:   true,
			expectMinSuccesses: 10,
			description:        "All concurrent operations with different IDs should succeed",
		},
		{
			name:               "concurrent_same_document",
			concurrentCount:    10,
			useFixedID:         true,
			expectAllSucceed:   false,
			expectMinSuccesses: 1,
			description:        "Only first concurrent operation with same ID should succeed",
		},
		{
			name:               "high_stress_concurrent",
			concurrentCount:    20,
			useFixedID:         false,
			expectAllSucceed:   true,
			expectMinSuccesses: 20,
			description:        "High stress test with 20 concurrent operations",
		},
	}

	for _, tt := range tests {
		sis.T().Run(tt.name, func(t *testing.T) {
			testDoc := sis.documents[0] // Use first document as template

			// Channel to collect results
			resultsChan := make(chan error, tt.concurrentCount)

			// WaitGroup to wait for all goroutines
			var wg sync.WaitGroup

			// Launch concurrent storage operations
			for i := 0; i < tt.concurrentCount; i++ {
				wg.Add(1)
				go func(index int) {
					defer wg.Done()

					// Create document copy
					doc := proto.Clone(testDoc).(*sbom.Document)

					if tt.useFixedID {
						// Use same ID for all - should cause conflicts
						doc.Metadata.Id = "concurrent-test-fixed-id"
					} else {
						// Use unique ID for each
						doc.Metadata.Id = fmt.Sprintf("concurrent-test-%d-%s", index, uuid.New().String())
					}

					// Store the document
					err := sis.backend.Store(doc, nil)
					resultsChan <- err
				}(i)
			}

			// Wait for all goroutines to complete
			wg.Wait()
			close(resultsChan)

			// Collect and analyze results
			var successCount, errorCount int
			var errors []error

			for err := range resultsChan {
				if err != nil {
					errorCount++
					errors = append(errors, err)
				} else {
					successCount++
				}
			}

			sis.T().Logf("Concurrent test %s: %d successes, %d errors out of %d operations",
				tt.name, successCount, errorCount, tt.concurrentCount)

			// Validate results based on expectations
			if tt.expectAllSucceed {
				sis.Equal(tt.concurrentCount, successCount, "All operations should succeed")
				sis.Equal(0, errorCount, "No operations should fail")
			} else {
				sis.GreaterOrEqual(successCount, tt.expectMinSuccesses,
					"Should have at least minimum expected successes")
				// For fixed ID tests, we expect constraint violations (not transaction errors)
				for _, err := range errors {
					sis.NotContains(err.Error(), "failed transaction",
						"Should not see transaction abort errors")
				}
			}
		})
	}
}

func (sis *storeIntegrationSuite) TestConcurrentMixedOperations() {
	// Test mixed concurrent operations (store + retrieve)
	testDoc := sis.documents[0]
	baseID := "mixed-concurrent-test"

	// Pre-store a document for retrieval testing
	doc := proto.Clone(testDoc).(*sbom.Document)
	doc.Metadata.Id = baseID
	err := sis.backend.Store(doc, nil)
	sis.Require().NoError(err)

	tests := []struct {
		name         string
		storeOps     int
		retrieveOps  int
		expectErrors bool
		description  string
	}{
		{
			name:         "mixed_store_retrieve",
			storeOps:     5,
			retrieveOps:  10,
			expectErrors: false,
			description:  "Mixed store and retrieve operations should work concurrently",
		},
		{
			name:         "heavy_retrieve_light_store",
			storeOps:     2,
			retrieveOps:  20,
			expectErrors: false,
			description:  "Heavy retrieve load with light store operations",
		},
	}

	for _, tt := range tests {
		sis.T().Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			resultsChan := make(chan error, tt.storeOps+tt.retrieveOps)

			// Launch store operations
			for i := 0; i < tt.storeOps; i++ {
				wg.Add(1)
				go func(index int) {
					defer wg.Done()

					doc := proto.Clone(testDoc).(*sbom.Document)
					doc.Metadata.Id = fmt.Sprintf("%s-store-%d-%s", baseID, index, uuid.New().String())

					err := sis.backend.Store(doc, nil)
					resultsChan <- err
				}(i)
			}

			// Launch retrieve operations
			for i := 0; i < tt.retrieveOps; i++ {
				wg.Add(1)
				go func(index int) {
					defer wg.Done()

					_, err := sis.backend.Retrieve(baseID, nil)
					resultsChan <- err
				}(i)
			}

			wg.Wait()
			close(resultsChan)

			// Analyze results
			var errorCount int
			for err := range resultsChan {
				if err != nil {
					errorCount++
					sis.T().Logf("Mixed operation error: %v", err)
				}
			}

			if tt.expectErrors {
				sis.Greater(errorCount, 0, "Expected some errors")
			} else {
				sis.Equal(0, errorCount, "Expected no errors in mixed operations")
			}
		})
	}
}

func TestStoreIntegrationSuite(t *testing.T) {
	suite.Run(t, new(storeIntegrationSuite))
}
