// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/protobom/protobom/pkg/reader"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/stretchr/testify/suite"

	"github.com/protobom/storage/backends/ent"
)

type annotationsSuite struct {
	suite.Suite
	*ent.Backend
	documents []*sbom.Document
}

// SetupSuite runs before the tests in the suite are run.
func (as *annotationsSuite) SetupSuite() {
	cwd, err := os.Getwd()
	as.Require().NoError(err)

	rdr := reader.New()
	testdataDir := filepath.Join(cwd, "testdata")

	entries, err := os.ReadDir(testdataDir)
	as.Require().NoError(err)

	for idx := range entries {
		document, err := rdr.ParseFile(filepath.Join(testdataDir, entries[idx].Name()))
		as.Require().NoError(err)

		as.documents = append(as.documents, document)
	}
}

// SetupTest runs before each test in the suite.
func (as *annotationsSuite) SetupTest() {
	as.Backend = ent.NewBackend(ent.WithDatabaseFile(":memory:"))
	as.Require().NoError(as.Backend.InitClient())

	for _, document := range as.documents {
		as.Require().NoError(as.Backend.Store(document, nil))
	}
}

// TearDownTest runs after each test in the suite.
func (as *annotationsSuite) TearDownTest() {
	as.Backend.CloseClient()
	as.Backend = nil
}

func (as *annotationsSuite) TestBackend_AddAnnotations() {
	// juice-shop-11.1.2.cdx.json serial number
	id := "urn:uuid:1f860713-54b9-4253-ba5a-9554851904af"
	annotationName := "add_annotation_test"

	as.Require().NoError(as.Backend.AddAnnotations(id, annotationName, "test-value-1", "test-value-2", "test-value-3"))

	annotations := as.getTestResult(annotationName)

	as.Len(annotations, 3)

	for idx, annotation := range annotations {
		as.Equal(annotationName, annotation.Name)
		as.Equal("test-value-"+strconv.Itoa(idx+1), annotation.Value)
	}
}

func (as *annotationsSuite) TestBackend_AddAnnotationToDocuments() {
	annotationName := "add_annotation_to_documents_test"
	documentIDs := []string{}

	for _, document := range as.documents {
		documentIDs = append(documentIDs, document.GetMetadata().GetId())
	}

	as.Require().NoError(as.Backend.AddAnnotationToDocuments(annotationName, "test-value", documentIDs...))

	annotations := as.getTestResult(annotationName)

	as.Len(annotations, 3)

	for _, annotation := range annotations {
		as.Equal(annotationName, annotation.Name)
		as.Equal("test-value", annotation.Value)
	}
}

func (as *annotationsSuite) TestBackend_ClearAnnotations() {
	annotationName := "clear_annotations_test"
	documentIDs := []string{}

	as.Require().NoError(as.Backend.AddAnnotationToDocuments(annotationName, "test-value", documentIDs...))
	as.Require().NoError(as.Backend.ClearAnnotations(documentIDs...))

	annotations := as.getTestResult(annotationName)

	as.Empty(annotations)
}

func (as *annotationsSuite) TestBackend_GetDocumentAnnotations() {
	// juice-shop-11.1.2.cdx.json serial number
	id := "urn:uuid:1f860713-54b9-4253-ba5a-9554851904af"
	annotationName := "get_document_annotations_test"

	as.Require().NoError(as.Backend.AddAnnotations(id, annotationName, "test-value-1", "test-value-2", "test-value-3"))

	annotations, err := as.Backend.GetDocumentAnnotations(id, annotationName)
	as.Require().NoError(err)

	as.Len(annotations, 3)

	for idx, annotation := range annotations {
		as.Equal(annotationName, annotation.Name)
		as.Equal("test-value-"+strconv.Itoa(idx+1), annotation.Value)
	}
}

func (as *annotationsSuite) TestBackend_GetDocumentsByAnnotation() {
	annotationName := "get_documents_by_annotation_test"
	query := strings.Join([]string{
		"INSERT INTO annotations (document_id, is_unique, name, value)",
		"VALUES (?, ?, ?, ?)",
		"ON CONFLICT DO NOTHING",
	}, " ")

	for idx, document := range as.documents {
		uniqueID, err := ent.GenerateUUID(document)
		as.Require().NoError(err)

		for _, value := range []string{
			"test-value-" + strconv.Itoa(idx+1),
			"test-value-" + strconv.Itoa(idx+2),
			"test-value-" + strconv.Itoa(idx+3),
		} {
			_, err = as.Backend.Client().ExecContext(as.Backend.Context(), query, uniqueID, false, annotationName, value)
			as.Require().NoError(err)
		}
	}

	spdxID := as.documents[0].GetMetadata().GetId()
	juiceShopID := as.documents[1].GetMetadata().GetId()
	cdxID := as.documents[2].GetMetadata().GetId()

	subtests := []struct{ values, expected []string }{
		{values: []string{"test-value-1"}, expected: []string{spdxID}},
		{values: []string{"test-value-2"}, expected: []string{spdxID, juiceShopID}},
		{values: []string{"test-value-3"}, expected: []string{spdxID, juiceShopID, cdxID}},
		{values: []string{"test-value-4"}, expected: []string{juiceShopID, cdxID}},
		{values: []string{"test-value-5"}, expected: []string{cdxID}},
		{values: []string{"test-value-1", "test-value-5"}, expected: []string{spdxID, cdxID}},
		{values: []string{}, expected: []string{spdxID, juiceShopID, cdxID}},
		{values: []string{"invalid-value"}, expected: []string{}},
	}

	for _, subtest := range subtests {
		prefix := strings.Join(subtest.values, "+")
		if prefix == "" {
			prefix = "no-values"
		}

		name := fmt.Sprintf("%s--expecting-%d-document(s)", prefix, len(subtest.expected))

		as.Run(name, func() {
			documents, err := as.Backend.GetDocumentsByAnnotation(annotationName, subtest.values...)
			as.Require().NoError(err)

			as.Len(documents, len(subtest.expected))

			gotIDs := []string{}

			for idx := range documents {
				gotIDs = append(gotIDs, documents[idx].GetMetadata().GetId())
			}

			as.ElementsMatch(gotIDs, subtest.expected)
		})
	}
}

func (as *annotationsSuite) TestBackend_GetDocumentUniqueAnnotation() {
	id, err := ent.GenerateUUID(as.documents[0])
	as.Require().NoError(err)

	annotationName := "get_document_unique_annotation_test"
	annotationValue := "unique-value"
	query := "INSERT INTO annotations (document_id, is_unique, name, value) VALUES (?, ?, ?, ?)"

	_, err = as.Backend.Client().ExecContext(as.Backend.Context(), query, id, true, annotationName, annotationValue)
	as.Require().NoError(err)

	got, err := as.Backend.GetDocumentUniqueAnnotation(as.documents[0].GetMetadata().GetId(), annotationName)
	as.Require().NoError(err)

	as.Equal(annotationValue, got)
}

func (as *annotationsSuite) TestBackend_RemoveAnnotations() {
	documentID := as.documents[0].GetMetadata().GetId()
	uniqueID, err := ent.GenerateUUID(as.documents[0])
	as.Require().NoError(err)

	annotationName := "remove_annotations_test"
	query := strings.Join([]string{
		"INSERT INTO annotations (document_id, is_unique, name, value)",
		"VALUES (?, ?, ?, ?)",
		"ON CONFLICT DO NOTHING",
	}, " ")

	subtests := []struct{ values, expected []string }{
		{values: []string{"test-value-2", "test-value-3"}, expected: []string{"test-value-1"}},
		{values: []string{"test-value-1"}, expected: []string{"test-value-2", "test-value-3"}},
		{values: []string{"unknown-value"}, expected: []string{"test-value-1", "test-value-2", "test-value-3"}},
		{values: []string{}, expected: []string{}},
	}

	for _, subtest := range subtests {
		prefix := strings.Join(subtest.values, "+")
		if prefix == "" {
			prefix = "no-values"
		}

		name := fmt.Sprintf("%s--expecting-%d-values(s)", prefix, len(subtest.expected))
		ctx := as.Backend.Context()

		as.Run(name, func() {
			for _, value := range []string{"test-value-1", "test-value-2", "test-value-3"} {
				_, err = as.Backend.Client().ExecContext(ctx, query, uniqueID, false, annotationName, value)
				as.Require().NoError(err)
			}

			as.Require().NoError(as.Backend.RemoveAnnotations(documentID, annotationName, subtest.values...))

			values := []string{}

			for _, annotation := range as.getTestResult(annotationName) {
				values = append(values, annotation.Value)
			}

			as.ElementsMatch(values, subtest.expected)
		})
	}
}

func (as *annotationsSuite) TestBackend_SetAnnotations() {
	documentID := as.documents[0].GetMetadata().GetId()

	annotationName := "set_annotations_test"
	updatedName := "set_annotations_test--updated"

	validateResults := func(name string, expected []string) {
		annotations := as.getTestResult(name)
		values := []string{}

		for _, annotation := range annotations {
			values = append(values, annotation.Value)
		}

		as.ElementsMatch(values, expected)
	}

	as.Require().NoError(as.Backend.SetAnnotations(documentID, annotationName, "test-value-1", "test-value-2"))
	validateResults(annotationName, []string{"test-value-1", "test-value-2"})

	// Replace annotation with different name, same values.
	// Verify previous annotation name is absent.
	as.Require().NoError(as.Backend.SetAnnotations(documentID, updatedName, "test-value-1", "test-value-2"))
	validateResults(updatedName, []string{"test-value-1", "test-value-2"})
	validateResults(annotationName, []string{})

	// Replace annotation with original name, different values.
	// Verify previous annotation name and previous values are absent.
	as.Require().NoError(as.Backend.SetAnnotations(documentID, annotationName, "test-value-3"))
	validateResults(annotationName, []string{"test-value-3"})
	validateResults(updatedName, []string{})
}

func (as *annotationsSuite) TestBackend_SetUniqueAnnotation() {
	annotationName := "set_unique_annotation_test"

	subtests := []struct {
		value                    string
		documentIdx, expectedLen int
	}{
		{documentIdx: 0, value: "unique-value", expectedLen: 1},
		{documentIdx: 1, value: "unique-value", expectedLen: 2},
		{documentIdx: 2, value: "unique-value", expectedLen: 3},
		{documentIdx: 0, value: "changed-value", expectedLen: 3},
		{documentIdx: 1, value: "changed-value", expectedLen: 3},
		{documentIdx: 2, value: "changed-value", expectedLen: 3},
	}

	for _, subtest := range subtests {
		name := fmt.Sprintf("document-%d-%s-%d-values(s)", subtest.documentIdx, subtest.value, subtest.expectedLen)

		as.Run(name, func() {
			documentID := as.documents[subtest.documentIdx].GetMetadata().GetId()
			uniqueID, err := ent.GenerateUUID(as.documents[subtest.documentIdx])
			as.Require().NoError(err)

			as.Require().NoError(as.Backend.SetUniqueAnnotation(documentID, annotationName, subtest.value))

			annotations := as.getTestResult(annotationName)

			as.Require().Len(annotations, subtest.expectedLen)
			as.Equal(uniqueID, annotations[subtest.documentIdx].DocumentID)
			as.Equal(annotationName, annotations[subtest.documentIdx].Name)
			as.Equal(subtest.value, annotations[subtest.documentIdx].Value)
		})
	}
}

func (as *annotationsSuite) getTestResult(annotationName string) ent.Annotations {
	as.T().Helper()

	result, err := as.Backend.Client().QueryContext(
		as.Backend.Context(),
		"SELECT * FROM annotations WHERE name == ?",
		annotationName,
	)
	as.Require().NoError(err)

	defer result.Close()

	annotations := ent.Annotations{}

	for result.Next() {
		annotation := &ent.Annotation{}

		as.Require().NoError(result.Scan(
			&annotation.ID,
			&annotation.Name,
			&annotation.Value,
			&annotation.IsUnique,
			&annotation.DocumentID,
		))

		annotations = append(annotations, annotation)
	}

	as.Require().NoError(result.Err())

	return annotations
}

func TestAnnotationsSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(annotationsSuite))
}
