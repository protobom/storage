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

func (backend *Backend) AddAnnotations(documentID, name string, values ...string) error {
	tx, err := backend.txClient()
	if err != nil {
		return err
	}

	builders := []*ent.AnnotationCreate{}

	for _, value := range values {
		builder := tx.Annotation.Create().
			SetDocumentID(documentID).
			SetName(name).
			SetValue(value)

		builders = append(builders, builder)
	}

	err = tx.Annotation.CreateBulk(builders...).
		OnConflict().
		Ignore().
		Exec(backend.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return rollback(tx, fmt.Errorf("creating annotations: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return rollback(tx, err)
	}

	return nil
}

func (backend *Backend) ClearAnnotations(documentIDs ...string) error {
	tx, err := backend.txClient()
	if err != nil {
		return err
	}

	_, err = tx.Annotation.Delete().
		Where(annotation.DocumentIDIn(documentIDs...)).
		Exec(backend.ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("clearing annotations: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return rollback(tx, err)
	}

	return nil
}

func (backend *Backend) GetDocumentAnnotations(documentID string, names ...string) (ent.Annotations, error) {
	if backend.client == nil {
		return nil, errUninitializedClient
	}

	query := backend.client.Annotation.Query().
		Where(annotation.HasDocumentWith(document.IDEQ(documentID)))

	if len(names) > 0 {
		query.Where(annotation.NameIn(names...))
	}

	annotations, err := query.All(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying annotations: %w", err)
	}

	return annotations, nil
}

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
		IDs(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying documents table: %w", err)
	}

	return backend.GetDocumentsByID(ids...)
}

func (backend *Backend) RemoveAnnotations(documentID, name string, values ...string) error {
	tx, err := backend.txClient()
	if err != nil {
		return err
	}

	_, err = tx.Annotation.Delete().
		Where(
			annotation.And(
				annotation.DocumentIDEQ(documentID),
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

func (backend *Backend) SetAnnotations(documentID, name string, values ...string) error {
	if err := backend.ClearAnnotations(documentID); err != nil {
		return err
	}

	if err := backend.AddAnnotations(documentID, name, values...); err != nil {
		return err
	}

	return nil
}
