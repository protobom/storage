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
	"testing"
	"time"

	"github.com/protobom/protobom/pkg/reader"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"google.golang.org/protobuf/proto"

	"github.com/protobom/storage/backends/ent"
)

// PostgreSQL Integration Test Suite
// Run with: go test -tags=integration -v ./backends/ent -run TestPostgresSuite
type postgresSuite struct {
	suite.Suite
	backend   *ent.Backend
	container testcontainers.Container
	documents []*sbom.Document
}

func (ps *postgresSuite) SetupSuite() {
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
	ps.Require().NoError(err)
	ps.container = pgContainer

	// Get connection string
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	ps.Require().NoError(err)

	// Initialize backend
	ps.backend = ent.NewBackend(
		ent.WithPostgresConnection(connStr),
	)
	ps.Require().NoError(ps.backend.InitClient())

	// Load test documents
	cwd, err := os.Getwd()
	ps.Require().NoError(err)

	rdr := reader.New()
	testdataDir := filepath.Join(cwd, "testdata")

	entries, err := os.ReadDir(testdataDir)
	ps.Require().NoError(err)

	for idx := range entries {
		document, err := rdr.ParseFile(filepath.Join(testdataDir, entries[idx].Name()))
		ps.Require().NoError(err)
		ps.documents = append(ps.documents, document)
	}
}

func (ps *postgresSuite) TearDownSuite() {
	if ps.backend != nil {
		ps.backend.CloseClient()
	}
	if ps.container != nil {
		ctx := context.Background()
		ps.Require().NoError(ps.container.Terminate(ctx))
	}
}

func (ps *postgresSuite) TestPostgresBackendStore() {
	tests := []struct {
		name           string
		document       *sbom.Document
		expectError    bool
		validateRetry  bool
	}{}

	// Build test cases from loaded documents
	for idx, doc := range ps.documents {
		documentID := doc.GetMetadata().GetId()
		testName := documentID
		if testName == "" {
			testName = fmt.Sprintf("document_%d", idx)
		}

		tests = append(tests, struct {
			name           string
			document       *sbom.Document
			expectError    bool
			validateRetry  bool
		}{
			name:          testName,
			document:      doc,
			expectError:   false,
			validateRetry: true,
		})
	}

	for _, tt := range tests {
		ps.T().Run(tt.name, func(t *testing.T) {
			// Test initial storage
			err := ps.backend.Store(tt.document, nil)
			if tt.expectError {
				ps.Require().Error(err)
				return
			}
			ps.Require().NoError(err)

			// Get retrieval ID
			retrievalID := tt.document.GetMetadata().GetId()
			if retrievalID == "" {
				generatedUUID, err := ent.GenerateUUID(tt.document.GetMetadata())
				ps.Require().NoError(err)
				retrievalID = generatedUUID.String()
			}

			// Test retrieval
			retrieved, err := ps.backend.Retrieve(retrievalID, nil)
			ps.Require().NoError(err)

			// Validate content equality
			retrievedCopy := proto.Clone(retrieved).(*sbom.Document)
			originalCopy := proto.Clone(tt.document).(*sbom.Document)
			retrievedCopy.GetMetadata().GetSourceData().Uri = nil
			originalCopy.GetMetadata().GetSourceData().Uri = nil

			ps.Require().True(proto.Equal(originalCopy, retrievedCopy))

			// Test duplicate storage (should not error)
			if tt.validateRetry {
				err = ps.backend.Store(tt.document, nil)
				ps.Require().NoError(err, "Duplicate storage should not fail")
			}
		})
	}
}


func (ps *postgresSuite) TestPostgresConflictHandling() {
	tests := []struct {
		name            string
		document        *sbom.Document
		conflictCount   int
		expectError     bool
		validateContent bool
	}{
		{
			name:            "single_document_conflicts",
			document:        ps.documents[0],
			conflictCount:   15,
			expectError:     false,
			validateContent: true,
		},
	}

	// Add cross-document conflicts if we have multiple documents
	if len(ps.documents) > 2 {
		tests = append(tests, struct {
			name            string
			document        *sbom.Document
			conflictCount   int
			expectError     bool
			validateContent bool
		}{
			name:            "cross_document_conflicts",
			document:        ps.documents[1],
			conflictCount:   5,
			expectError:     false,
			validateContent: true,
		})
	}

	for _, tt := range tests {
		ps.T().Run(tt.name, func(t *testing.T) {
			// Initial storage
			err := ps.backend.Store(tt.document, nil)
			ps.Require().NoError(err, "Initial storage should succeed")

			// Test multiple conflicts
			for i := 0; i < tt.conflictCount; i++ {
				err = ps.backend.Store(tt.document, nil)
				if tt.expectError {
					ps.Require().Error(err, fmt.Sprintf("Conflict test %d should fail", i+1))
				} else {
					ps.Require().NoError(err, fmt.Sprintf("Conflict test %d should succeed", i+1))
				}
			}

			// Validate content integrity if requested
			if tt.validateContent {
				retrieved, err := ps.backend.Retrieve(tt.document.GetMetadata().GetId(), nil)
				ps.Require().NoError(err)
				ps.Require().NotNil(retrieved)

				// Compare content
				retrievedCopy := proto.Clone(retrieved).(*sbom.Document)
				testDocCopy := proto.Clone(tt.document).(*sbom.Document)
				retrievedCopy.GetMetadata().GetSourceData().Uri = nil
				testDocCopy.GetMetadata().GetSourceData().Uri = nil

				ps.Require().True(proto.Equal(testDocCopy, retrievedCopy), "Retrieved document should match original after conflict resolution")
			}
		})
	}
}

func (ps *postgresSuite) TestPostgresAnnotations() {
	testDoc := ps.documents[0]

	// Store the document first
	err := ps.backend.Store(testDoc, nil)
	ps.Require().NoError(err, "Initial document storage should succeed")

	docID := testDoc.GetMetadata().GetId()

	tests := []struct {
		name           string
		setupFunc      func() error
		testFunc       func() error
		expectedValues []string
		expectedCount  int
		expectError    bool
	}{
		{
			name: "basic_document_annotations",
			setupFunc: func() error {
				return ps.backend.AddDocumentAnnotations(docID, "test-annotation-1", "test-value-1", "test-value-2")
			},
			testFunc: func() error {
				annotations, err := ps.backend.GetDocumentAnnotations(docID, "test-annotation-1")
				if err != nil {
					return err
				}
				if len(annotations) != 2 {
					return fmt.Errorf("expected 2 annotations, got %d", len(annotations))
				}
				return nil
			},
			expectedCount: 2,
			expectError:   false,
		},
		{
			name: "update_document_annotations",
			setupFunc: func() error {
				return ps.backend.SetDocumentAnnotations(docID, "test-annotation-1", "updated-value-1", "updated-value-2", "updated-value-3")
			},
			testFunc: func() error {
				annotations, err := ps.backend.GetDocumentAnnotations(docID, "test-annotation-1")
				if err != nil {
					return err
				}
				if len(annotations) != 3 {
					return fmt.Errorf("expected 3 annotations, got %d", len(annotations))
				}
				return nil
			},
			expectedCount: 3,
			expectError:   false,
		},
		{
			name: "unique_document_annotation",
			setupFunc: func() error {
				return ps.backend.SetDocumentUniqueAnnotation(docID, "unique-test", "unique-value")
			},
			testFunc: func() error {
				value, err := ps.backend.GetDocumentUniqueAnnotation(docID, "unique-test")
				if err != nil {
					return err
				}
				if value != "unique-value" {
					return fmt.Errorf("expected 'unique-value', got '%s'", value)
				}
				return nil
			},
			expectedValues: []string{"unique-value"},
			expectError:    false,
		},
		{
			name: "update_unique_document_annotation",
			setupFunc: func() error {
				return ps.backend.SetDocumentUniqueAnnotation(docID, "unique-test", "updated-unique-value")
			},
			testFunc: func() error {
				value, err := ps.backend.GetDocumentUniqueAnnotation(docID, "unique-test")
				if err != nil {
					return err
				}
				if value != "updated-unique-value" {
					return fmt.Errorf("expected 'updated-unique-value', got '%s'", value)
				}
				return nil
			},
			expectedValues: []string{"updated-unique-value"},
			expectError:    false,
		},
		{
			name: "bulk_document_annotations",
			setupFunc: func() error {
				bulkValues := make([]string, 100)
				for i := 0; i < 100; i++ {
					bulkValues[i] = fmt.Sprintf("bulk-value-%d", i)
				}
				return ps.backend.AddDocumentAnnotations(docID, "bulk-annotation", bulkValues...)
			},
			testFunc: func() error {
				annotations, err := ps.backend.GetDocumentAnnotations(docID, "bulk-annotation")
				if err != nil {
					return err
				}
				if len(annotations) != 100 {
					return fmt.Errorf("expected 100 annotations, got %d", len(annotations))
				}
				return nil
			},
			expectedCount: 100,
			expectError:   false,
		},
	}

	for _, tt := range tests {
		ps.T().Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setupFunc != nil {
				err := tt.setupFunc()
				if tt.expectError {
					ps.Require().Error(err)
					return
				}
				ps.Require().NoError(err)
			}

			// Test
			if tt.testFunc != nil {
				err := tt.testFunc()
				if tt.expectError {
					ps.Require().Error(err)
				} else {
					ps.Require().NoError(err)
				}
			}
		})
	}
}

func (ps *postgresSuite) TestPostgresNodeAnnotations() {
	testDoc := ps.documents[0]

	// Store the document first
	err := ps.backend.Store(testDoc, nil)
	ps.Require().NoError(err, "Initial document storage should succeed")

	// Skip if no nodes available
	if len(testDoc.GetNodeList().GetNodes()) == 0 {
		ps.T().Skip("No nodes available for testing")
		return
	}

	firstNode := testDoc.GetNodeList().GetNodes()[0]
	nodeID := firstNode.GetId()

	tests := []struct {
		name          string
		setupFunc     func() error
		testFunc      func() error
		expectedCount int
		expectedValue string
		expectError   bool
	}{
		{
			name: "basic_node_annotations",
			setupFunc: func() error {
				return ps.backend.AddNodeAnnotations(nodeID, "node-annotation", "node-value-1", "node-value-2")
			},
			testFunc: func() error {
				annotations, err := ps.backend.GetNodeAnnotations(nodeID, "node-annotation")
				if err != nil {
					return err
				}
				if len(annotations) != 2 {
					return fmt.Errorf("expected 2 node annotations, got %d", len(annotations))
				}
				return nil
			},
			expectedCount: 2,
			expectError:   false,
		},
		{
			name: "unique_node_annotation",
			setupFunc: func() error {
				return ps.backend.SetNodeUniqueAnnotation(nodeID, "unique-node-annotation", "unique-node-value")
			},
			testFunc: func() error {
				value, err := ps.backend.GetNodeUniqueAnnotation(nodeID, "unique-node-annotation")
				if err != nil {
					return err
				}
				if value != "unique-node-value" {
					return fmt.Errorf("expected 'unique-node-value', got '%s'", value)
				}
				return nil
			},
			expectedValue: "unique-node-value",
			expectError:   false,
		},
		{
			name: "update_unique_node_annotation",
			setupFunc: func() error {
				return ps.backend.SetNodeUniqueAnnotation(nodeID, "unique-node-annotation", "updated-node-value")
			},
			testFunc: func() error {
				value, err := ps.backend.GetNodeUniqueAnnotation(nodeID, "unique-node-annotation")
				if err != nil {
					return err
				}
				if value != "updated-node-value" {
					return fmt.Errorf("expected 'updated-node-value', got '%s'", value)
				}
				return nil
			},
			expectedValue: "updated-node-value",
			expectError:   false,
		},
	}

	for _, tt := range tests {
		ps.T().Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setupFunc != nil {
				err := tt.setupFunc()
				if tt.expectError {
					ps.Require().Error(err)
					return
				}
				ps.Require().NoError(err)
			}

			// Test
			if tt.testFunc != nil {
				err := tt.testFunc()
				if tt.expectError {
					ps.Require().Error(err)
				} else {
					ps.Require().NoError(err)
				}
			}
		})
	}
}

func (ps *postgresSuite) TestPostgresAnnotationManagement() {
	testDoc := ps.documents[0]

	// Store the document first
	err := ps.backend.Store(testDoc, nil)
	ps.Require().NoError(err, "Initial document storage should succeed")

	docID := testDoc.GetMetadata().GetId()

	tests := []struct {
		name          string
		setupFunc     func() error
		actionFunc    func() error
		validateFunc  func() error
		expectError   bool
	}{
		{
			name: "remove_specific_annotation",
			setupFunc: func() error {
				return ps.backend.AddDocumentAnnotations(docID, "remove-test", "value-1", "value-2", "value-3")
			},
			actionFunc: func() error {
				return ps.backend.RemoveDocumentAnnotations(docID, "remove-test", "value-2")
			},
			validateFunc: func() error {
				annotations, err := ps.backend.GetDocumentAnnotations(docID, "remove-test")
				if err != nil {
					return err
				}
				if len(annotations) != 2 {
					return fmt.Errorf("expected 2 remaining annotations, got %d", len(annotations))
				}
				return nil
			},
			expectError: false,
		},
		{
			name: "clear_all_annotations",
			setupFunc: func() error {
				return ps.backend.AddDocumentAnnotations(docID, "clear-test", "value-1", "value-2", "value-3")
			},
			actionFunc: func() error {
				return ps.backend.ClearDocumentAnnotations(docID, "clear-test")
			},
			validateFunc: func() error {
				annotations, err := ps.backend.GetDocumentAnnotations(docID, "clear-test")
				if err != nil {
					return err
				}
				if len(annotations) != 0 {
					return fmt.Errorf("expected 0 annotations after clear, got %d", len(annotations))
				}
				return nil
			},
			expectError: false,
		},
		{
			name: "bulk_annotation_conflict_resolution",
			setupFunc: func() error {
				bulkValues := make([]string, 50)
				for i := 0; i < 50; i++ {
					bulkValues[i] = fmt.Sprintf("bulk-value-%d", i)
				}
				return ps.backend.AddDocumentAnnotations(docID, "conflict-test", bulkValues...)
			},
			actionFunc: func() error {
				// Add the same values again - should handle conflicts
				bulkValues := make([]string, 50)
				for i := 0; i < 50; i++ {
					bulkValues[i] = fmt.Sprintf("bulk-value-%d", i)
				}
				return ps.backend.AddDocumentAnnotations(docID, "conflict-test", bulkValues...)
			},
			validateFunc: func() error {
				annotations, err := ps.backend.GetDocumentAnnotations(docID, "conflict-test")
				if err != nil {
					return err
				}
				if len(annotations) < 50 {
					return fmt.Errorf("expected at least 50 annotations after conflict resolution, got %d", len(annotations))
				}
				return nil
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		ps.T().Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setupFunc != nil {
				err := tt.setupFunc()
				ps.Require().NoError(err)
			}

			// Action
			if tt.actionFunc != nil {
				err := tt.actionFunc()
				if tt.expectError {
					ps.Require().Error(err)
					return
				}
				ps.Require().NoError(err)
			}

			// Validate
			if tt.validateFunc != nil {
				err := tt.validateFunc()
				ps.Require().NoError(err)
			}
		})
	}
}

func (ps *postgresSuite) TestPostgresConfiguration() {
	ctx := context.Background()
	connStr, err := ps.container.(*postgres.PostgresContainer).ConnectionString(ctx, "sslmode=disable")
	ps.Require().NoError(err)

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
		ps.T().Run(tt.name, func(t *testing.T) {
			backend := tt.backendSetup()

			// Validate configuration
			ps.Equal(tt.expectedDialect, backend.Options.Dialect)
			ps.Equal(tt.expectedURL, backend.Options.DatabaseURL)

			// Test initialization if required
			if tt.shouldInit {
				err := backend.InitClient()
				ps.Require().NoError(err)
				backend.CloseClient()
			}
		})
	}
}

func TestPostgresSuite(t *testing.T) {
	suite.Run(t, new(postgresSuite))
}
