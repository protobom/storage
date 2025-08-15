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

	"github.com/protobom/storage/backends/ent"
)

// Annotations Integration Test Suite
// Tests annotation functionality with PostgreSQL database
type annotationsIntegrationSuite struct {
	suite.Suite
	backend   *ent.Backend
	container testcontainers.Container
	documents []*sbom.Document
}

func (ais *annotationsIntegrationSuite) SetupSuite() {
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
	ais.Require().NoError(err)
	ais.container = pgContainer

	// Get connection string
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	ais.Require().NoError(err)

	// Initialize backend
	ais.backend = ent.NewBackend(
		ent.WithPostgresConnection(connStr),
	)
	ais.Require().NoError(ais.backend.InitClient())

	// Load test documents
	cwd, err := os.Getwd()
	ais.Require().NoError(err)

	rdr := reader.New()
	testdataDir := filepath.Join(cwd, "testdata")

	entries, err := os.ReadDir(testdataDir)
	ais.Require().NoError(err)

	for idx := range entries {
		document, err := rdr.ParseFile(filepath.Join(testdataDir, entries[idx].Name()))
		ais.Require().NoError(err)
		ais.documents = append(ais.documents, document)
	}

	// Store all test documents for annotation testing
	for _, doc := range ais.documents {
		err := ais.backend.Store(doc, nil)
		ais.Require().NoError(err)
	}
}

func (ais *annotationsIntegrationSuite) TearDownSuite() {
	if ais.backend != nil {
		ais.backend.CloseClient()
	}
	if ais.container != nil {
		ctx := context.Background()
		ais.Require().NoError(ais.container.Terminate(ctx))
	}
}

func (ais *annotationsIntegrationSuite) TestDocumentAnnotations() {
	testDoc := ais.documents[0]
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
				return ais.backend.AddDocumentAnnotations(docID, "test-annotation-1", "test-value-1", "test-value-2")
			},
			testFunc: func() error {
				annotations, err := ais.backend.GetDocumentAnnotations(docID, "test-annotation-1")
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
				return ais.backend.SetDocumentAnnotations(docID, "test-annotation-1", "updated-value-1", "updated-value-2", "updated-value-3")
			},
			testFunc: func() error {
				annotations, err := ais.backend.GetDocumentAnnotations(docID, "test-annotation-1")
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
				return ais.backend.SetDocumentUniqueAnnotation(docID, "unique-test", "unique-value")
			},
			testFunc: func() error {
				value, err := ais.backend.GetDocumentUniqueAnnotation(docID, "unique-test")
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
				return ais.backend.SetDocumentUniqueAnnotation(docID, "unique-test", "updated-unique-value")
			},
			testFunc: func() error {
				value, err := ais.backend.GetDocumentUniqueAnnotation(docID, "unique-test")
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
				return ais.backend.AddDocumentAnnotations(docID, "bulk-annotation", bulkValues...)
			},
			testFunc: func() error {
				annotations, err := ais.backend.GetDocumentAnnotations(docID, "bulk-annotation")
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
		ais.T().Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setupFunc != nil {
				err := tt.setupFunc()
				if tt.expectError {
					ais.Require().Error(err)
					return
				}
				ais.Require().NoError(err)
			}

			// Test
			if tt.testFunc != nil {
				err := tt.testFunc()
				if tt.expectError {
					ais.Require().Error(err)
				} else {
					ais.Require().NoError(err)
				}
			}
		})
	}
}

func (ais *annotationsIntegrationSuite) TestNodeAnnotations() {
	testDoc := ais.documents[0]

	// Skip if no nodes available
	if len(testDoc.GetNodeList().GetNodes()) == 0 {
		ais.T().Skip("No nodes available for testing")
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
				return ais.backend.AddNodeAnnotations(nodeID, "node-annotation", "node-value-1", "node-value-2")
			},
			testFunc: func() error {
				annotations, err := ais.backend.GetNodeAnnotations(nodeID, "node-annotation")
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
				return ais.backend.SetNodeUniqueAnnotation(nodeID, "unique-node-annotation", "unique-node-value")
			},
			testFunc: func() error {
				value, err := ais.backend.GetNodeUniqueAnnotation(nodeID, "unique-node-annotation")
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
				return ais.backend.SetNodeUniqueAnnotation(nodeID, "unique-node-annotation", "updated-node-value")
			},
			testFunc: func() error {
				value, err := ais.backend.GetNodeUniqueAnnotation(nodeID, "unique-node-annotation")
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
		ais.T().Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setupFunc != nil {
				err := tt.setupFunc()
				if tt.expectError {
					ais.Require().Error(err)
					return
				}
				ais.Require().NoError(err)
			}

			// Test
			if tt.testFunc != nil {
				err := tt.testFunc()
				if tt.expectError {
					ais.Require().Error(err)
				} else {
					ais.Require().NoError(err)
				}
			}
		})
	}
}

func (ais *annotationsIntegrationSuite) TestAnnotationManagement() {
	testDoc := ais.documents[0]
	docID := testDoc.GetMetadata().GetId()

	tests := []struct {
		name         string
		setupFunc    func() error
		actionFunc   func() error
		validateFunc func() error
		expectError  bool
	}{
		{
			name: "remove_specific_annotation",
			setupFunc: func() error {
				return ais.backend.AddDocumentAnnotations(docID, "remove-test", "value-1", "value-2", "value-3")
			},
			actionFunc: func() error {
				return ais.backend.RemoveDocumentAnnotations(docID, "remove-test", "value-2")
			},
			validateFunc: func() error {
				annotations, err := ais.backend.GetDocumentAnnotations(docID, "remove-test")
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
				return ais.backend.AddDocumentAnnotations(docID, "clear-test", "value-1", "value-2", "value-3")
			},
			actionFunc: func() error {
				return ais.backend.ClearDocumentAnnotations(docID, "clear-test")
			},
			validateFunc: func() error {
				annotations, err := ais.backend.GetDocumentAnnotations(docID, "clear-test")
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
				return ais.backend.AddDocumentAnnotations(docID, "conflict-test", bulkValues...)
			},
			actionFunc: func() error {
				// Add the same values again - should handle conflicts
				bulkValues := make([]string, 50)
				for i := 0; i < 50; i++ {
					bulkValues[i] = fmt.Sprintf("bulk-value-%d", i)
				}
				return ais.backend.AddDocumentAnnotations(docID, "conflict-test", bulkValues...)
			},
			validateFunc: func() error {
				annotations, err := ais.backend.GetDocumentAnnotations(docID, "conflict-test")
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
		ais.T().Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setupFunc != nil {
				err := tt.setupFunc()
				ais.Require().NoError(err)
			}

			// Action
			if tt.actionFunc != nil {
				err := tt.actionFunc()
				if tt.expectError {
					ais.Require().Error(err)
					return
				}
				ais.Require().NoError(err)
			}

			// Validate
			if tt.validateFunc != nil {
				err := tt.validateFunc()
				ais.Require().NoError(err)
			}
		})
	}
}

func (ais *annotationsIntegrationSuite) TestConcurrentAnnotations() {
	// Test concurrent annotation operations
	testDoc := ais.documents[0]

	// Store the document first
	err := ais.backend.Store(testDoc, nil)
	ais.Require().NoError(err)

	docID := testDoc.GetMetadata().GetId()

	tests := []struct {
		name            string
		concurrentCount int
		testFunc        func(index int) error
		expectErrors    bool
		description     string
	}{
		{
			name:            "concurrent_document_annotations",
			concurrentCount: 10,
			testFunc: func(index int) error {
				return ais.backend.AddDocumentAnnotations(docID, "concurrent-test",
					fmt.Sprintf("value-%d", index))
			},
			expectErrors: false,
			description:  "Concurrent document annotation additions should succeed",
		},
		{
			name:            "concurrent_unique_annotations",
			concurrentCount: 10,
			testFunc: func(index int) error {
				return ais.backend.SetDocumentUniqueAnnotation(docID,
					fmt.Sprintf("unique-key-%d", index), fmt.Sprintf("value-%d", index))
			},
			expectErrors: false,
			description:  "Concurrent unique annotation sets should succeed",
		},
		{
			name:            "concurrent_same_unique_annotation",
			concurrentCount: 10,
			testFunc: func(index int) error {
				return ais.backend.SetDocumentUniqueAnnotation(docID,
					"same-unique-key", fmt.Sprintf("value-%d", index))
			},
			expectErrors: false, // Last one wins, no errors expected
			description:  "Concurrent same unique annotation sets should handle conflicts",
		},
	}

	for _, tt := range tests {
		ais.T().Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			resultsChan := make(chan error, tt.concurrentCount)

			// Launch concurrent annotation operations
			for i := 0; i < tt.concurrentCount; i++ {
				wg.Add(1)
				go func(index int) {
					defer wg.Done()
					err := tt.testFunc(index)
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
					ais.T().Logf("Concurrent annotation error: %v", err)
				}
			}

			if tt.expectErrors {
				ais.Greater(errorCount, 0, "Expected some errors")
			} else {
				ais.Equal(0, errorCount, "Expected no errors in concurrent annotations")
			}
		})
	}
}

func (ais *annotationsIntegrationSuite) TestConcurrentAnnotationOperations() {
	// Test mixed concurrent annotation operations (add, get, remove, clear)
	testDoc := ais.documents[0]

	// Store the document first
	err := ais.backend.Store(testDoc, nil)
	ais.Require().NoError(err)

	docID := testDoc.GetMetadata().GetId()

	// Pre-populate some annotations
	err = ais.backend.AddDocumentAnnotations(docID, "mixed-test",
		"initial-1", "initial-2", "initial-3")
	ais.Require().NoError(err)

	tests := []struct {
		name         string
		addOps       int
		getOps       int
		removeOps    int
		expectErrors bool
		description  string
	}{
		{
			name:         "mixed_annotation_operations",
			addOps:       5,
			getOps:       10,
			removeOps:    2,
			expectErrors: false,
			description:  "Mixed annotation operations should work concurrently",
		},
		{
			name:         "heavy_read_light_write",
			addOps:       2,
			getOps:       20,
			removeOps:    1,
			expectErrors: false,
			description:  "Heavy read operations with light write operations",
		},
	}

	for _, tt := range tests {
		ais.T().Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			totalOps := tt.addOps + tt.getOps + tt.removeOps
			resultsChan := make(chan error, totalOps)

			// Launch add operations
			for i := 0; i < tt.addOps; i++ {
				wg.Add(1)
				go func(index int) {
					defer wg.Done()
					err := ais.backend.AddDocumentAnnotations(docID, "concurrent-mixed",
						fmt.Sprintf("add-value-%d-%s", index, uuid.New().String()))
					resultsChan <- err
				}(i)
			}

			// Launch get operations
			for i := 0; i < tt.getOps; i++ {
				wg.Add(1)
				go func(index int) {
					defer wg.Done()
					_, err := ais.backend.GetDocumentAnnotations(docID, "mixed-test")
					resultsChan <- err
				}(i)
			}

			// Launch remove operations
			for i := 0; i < tt.removeOps; i++ {
				wg.Add(1)
				go func(index int) {
					defer wg.Done()
					err := ais.backend.RemoveDocumentAnnotations(docID, "mixed-test",
						fmt.Sprintf("initial-%d", index+1))
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
					ais.T().Logf("Mixed annotation operation error: %v", err)
				}
			}

			if tt.expectErrors {
				ais.Greater(errorCount, 0, "Expected some errors")
			} else {
				ais.Equal(0, errorCount, "Expected no errors in mixed annotation operations")
			}
		})
	}
}

func TestAnnotationsIntegrationSuite(t *testing.T) {
	suite.Run(t, new(annotationsIntegrationSuite))
}
