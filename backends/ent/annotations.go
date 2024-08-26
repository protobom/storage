// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"fmt"

	"github.com/protobom/protobom/pkg/sbom"

	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/annotation"
	"github.com/protobom/storage/internal/backends/ent/document"
)

func (backend *Backend) createAnnotations(data ...*ent.Annotation) error {
	tx, err := backend.txClient()
	if err != nil {
		return err
	}

	builders := []*ent.AnnotationCreate{}

	for idx := range data {
		builder := tx.Annotation.Create().
			SetDocumentID(data[idx].DocumentID).
			SetName(data[idx].Name).
			SetValue(data[idx].Value)

		builders = append(builders, builder)
	}

	err = tx.Annotation.CreateBulk(builders...).
		OnConflict().
		UpdateNewValues().
		Exec(backend.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return rollback(tx, fmt.Errorf("creating annotations: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return rollback(tx, err)
	}

	return nil
}

// AddAnnotations applies multiple named annotation values to a single document.
func (backend *Backend) AddAnnotations(documentID, name string, values ...string) error {
	data := ent.Annotations{}

	doc, err := backend.client.Document.Query().
		Where(document.MetadataIDEQ(documentID)).
		Only(backend.ctx)
	if err != nil {
		return fmt.Errorf("querying documents: %w", err)
	}

	for _, value := range values {
		data = append(data, &ent.Annotation{
			DocumentID: doc.ID,
			Name:       name,
			Value:      value,
		})
	}

	return backend.createAnnotations(data...)
}

// AddAnnotationToDocuments applies a single named annotation value to multiple documents.
func (backend *Backend) AddAnnotationToDocuments(name, value string, documentIDs ...string) error {
	data := ent.Annotations{}

	for _, documentID := range documentIDs {
		doc, err := backend.client.Document.Query().
			Where(document.MetadataIDEQ(documentID)).
			Only(backend.ctx)
		if err != nil {
			return fmt.Errorf("querying documents: %w", err)
		}

		data = append(data, &ent.Annotation{
			DocumentID: doc.ID,
			Name:       name,
			Value:      value,
		})
	}

	return backend.createAnnotations(data...)
}

// ClearAnnotations removes all annotations from the specified documents.
func (backend *Backend) ClearAnnotations(documentIDs ...string) error {
	tx, err := backend.txClient()
	if err != nil {
		return err
	}

	_, err = tx.Annotation.Delete().
		Where(annotation.HasDocumentWith(document.MetadataIDIn(documentIDs...))).
		Exec(backend.ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("clearing annotations: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return rollback(tx, err)
	}

	return nil
}

// GetDocumentAlias gets the value for the annotation named "alias".
func (backend *Backend) GetDocumentAlias(documentID string) (string, error) {
	if backend.client == nil {
		return "", errUninitializedClient
	}

	query := backend.client.Annotation.Query().
		Where(
			annotation.And(
				annotation.HasDocumentWith(document.MetadataIDEQ(documentID)),
				annotation.NameEQ("alias"),
			),
		)

	result, err := query.Only(backend.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return "", nil
		}

		return "", fmt.Errorf("retrieving document alias: %w", err)
	}

	return result.Value, nil
}

// GetDocumentAnnotations gets all annotations for the specified
// document, limited to a set of annotation names if specified.
func (backend *Backend) GetDocumentAnnotations(documentID string, names ...string) (ent.Annotations, error) {
	if backend.client == nil {
		return nil, errUninitializedClient
	}

	query := backend.client.Annotation.Query().
		Where(annotation.HasDocumentWith(document.MetadataIDEQ(documentID)))

	if len(names) > 0 {
		query.Where(annotation.NameIn(names...))
	}

	annotations, err := query.All(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying annotations: %w", err)
	}

	return annotations, nil
}

// GetDocumentsByAnnotation gets all documents having the specified named
// annotation, limited to a set of annotation values if specified.
func (backend *Backend) GetDocumentsByAnnotation(name string, values ...string) ([]*sbom.Document, error) {
	if backend.client == nil {
		return nil, errUninitializedClient
	}

	ids, err := backend.client.Document.Query().
		Where(
			document.HasAnnotationsWith(
				annotation.And(
					annotation.NameEQ(name),
					annotation.ValueIn(values...),
				),
			),
		).
		QueryMetadata().
		IDs(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying documents table: %w", err)
	}

	return backend.GetDocumentsByID(ids...)
}

// RemoveAnnotations removes all annotations with the specified name from
// the document, limited to a set of annotation values if specified.
func (backend *Backend) RemoveAnnotations(documentID, name string, values ...string) error {
	tx, err := backend.txClient()
	if err != nil {
		return err
	}

	_, err = tx.Annotation.Delete().
		Where(
			annotation.And(
				annotation.HasDocumentWith(document.MetadataIDEQ(documentID)),
				annotation.NameEQ(name),
				annotation.ValueIn(values...),
			),
		).
		Exec(backend.ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("removing annotations: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return rollback(tx, err)
	}

	return nil
}

// SetDocumentAlias set the value for the annotation named "alias".
func (backend *Backend) SetDocumentAlias(documentID, value string) error {
	if err := backend.RemoveAnnotations(documentID, "alias"); err != nil {
		return err
	}

	return backend.AddAnnotations(documentID, "alias", value)
}

// SetAnnotations explicitly sets the named annotations for the specified document.
func (backend *Backend) SetAnnotations(documentID, name string, values ...string) error {
	if err := backend.ClearAnnotations(documentID); err != nil {
		return err
	}

	return backend.AddAnnotations(documentID, name, values...)
}
