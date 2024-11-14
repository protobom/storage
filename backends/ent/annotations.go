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
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// AddAnnotationToDocuments applies a single named annotation value to multiple documents.
func (backend *Backend) AddAnnotationToDocuments(name, value string, documentIDs ...string) error {
	data := ent.Annotations{}

	for _, documentID := range documentIDs {
		documentUUID, err := backend.client.Document.Query().
			Where(document.MetadataIDEQ(documentID)).
			OnlyID(backend.ctx)
		if err != nil {
			return fmt.Errorf("querying documents: %w", err)
		}

		data = append(data, &ent.Annotation{
			DocumentID: documentUUID,
			Name:       name,
			Value:      value,
		})
	}

	return backend.withTx(backend.saveAnnotations(data...))
}

// AddAnnotationToNodes applies a single named annotation value to multiple nodes.
func (backend *Backend) AddAnnotationToNodes(name, value string, nodeIDs ...string) error {
	data := ent.Annotations{}

	for _, nodeID := range nodeIDs {
		data = append(data, &ent.Annotation{
			NodeID: &nodeID,
			Name:   name,
			Value:  value,
		})
	}

	return backend.withTx(backend.saveAnnotations(data...))
}

// AddDocumentAnnotations applies multiple named annotation values to a single document.
func (backend *Backend) AddDocumentAnnotations(documentID, name string, values ...string) error {
	data := ent.Annotations{}

	documentUUID, err := backend.client.Document.Query().
		Where(document.MetadataIDEQ(documentID)).
		OnlyID(backend.ctx)
	if err != nil {
		return fmt.Errorf("querying documents: %w", err)
	}

	for _, value := range values {
		data = append(data, &ent.Annotation{
			DocumentID: documentUUID,
			Name:       name,
			Value:      value,
		})
	}

	return backend.withTx(backend.saveAnnotations(data...))
}

// AddNodeAnnotations applies multiple named annotation values to a single node.
func (backend *Backend) AddNodeAnnotations(nodeID, name string, values ...string) error {
	data := ent.Annotations{}

	for _, value := range values {
		data = append(data, &ent.Annotation{
			NodeID: &nodeID,
			Name:   name,
			Value:  value,
		})
	}

	return backend.withTx(backend.saveAnnotations(data...))
}

// ClearDocumentAnnotations removes all annotations from the specified documents.
func (backend *Backend) ClearDocumentAnnotations(documentIDs ...string) error {
	if len(documentIDs) == 0 {
		return nil
	}

	return backend.withTx(func(tx *ent.Tx) error {
		if _, err := tx.Annotation.Delete().
			Where(annotation.HasDocumentWith(document.MetadataIDIn(documentIDs...))).
			Exec(backend.ctx); err != nil {
			return fmt.Errorf("clearing annotations: %w", err)
		}

		return nil
	})
}

// ClearNodeAnnotations removes all annotations from the specified nodes.
func (backend *Backend) ClearNodeAnnotations(nodeIDs ...string) error {
	if len(nodeIDs) == 0 {
		return nil
	}

	return backend.withTx(func(tx *ent.Tx) error {
		if _, err := tx.Annotation.Delete().
			Where(annotation.HasNodeWith(node.IDIn(nodeIDs...))).
			Exec(backend.ctx); err != nil {
			return fmt.Errorf("clearing annotations: %w", err)
		}

		return nil
	})
}

// GetDocumentAnnotations gets all annotations for the specified
// document, limited to a set of annotation names if specified.
func (backend *Backend) GetDocumentAnnotations(documentID string, names ...string) (ent.Annotations, error) {
	if backend.client == nil {
		return nil, errUninitializedClient
	}

	predicates := []predicate.Annotation{
		annotation.HasDocumentWith(document.MetadataIDEQ(documentID)),
	}

	if len(names) > 0 {
		predicates = append(predicates, annotation.NameIn(names...))
	}

	annotations, err := backend.client.Annotation.Query().Where(predicates...).All(backend.ctx)
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

	predicates := []predicate.Annotation{annotation.NameEQ(name)}

	if len(values) > 0 {
		predicates = append(predicates, annotation.ValueIn(values...))
	}

	ids, err := backend.client.Annotation.Query().
		Where(predicates...).
		QueryDocument().
		QueryMetadata().
		IDs(backend.ctx)

	if err != nil {
		return nil, fmt.Errorf("querying documents table: %w", err)
	}

	if len(ids) == 0 {
		return []*sbom.Document{}, nil
	}

	return backend.GetDocumentsByID(ids...)
}

// GetDocumentUniqueAnnotation gets the value for a unique annotation.
func (backend *Backend) GetDocumentUniqueAnnotation(documentID, name string) (string, error) {
	if backend.client == nil {
		return "", errUninitializedClient
	}

	result, err := backend.client.Annotation.Query().
		Where(
			annotation.HasDocumentWith(document.MetadataIDEQ(documentID)),
			annotation.NameEQ(name),
			annotation.IsUniqueEQ(true),
		).
		Only(backend.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return "", nil
		}

		return "", fmt.Errorf("retrieving unique annotation for document: %w", err)
	}

	return result.Value, nil
}

// GetNodeAnnotations gets all annotations for the specified
// node, limited to a set of annotation names if specified.
func (backend *Backend) GetNodeAnnotations(nodeID string, names ...string) (ent.Annotations, error) {
	if backend.client == nil {
		return nil, errUninitializedClient
	}

	predicates := []predicate.Annotation{
		annotation.HasNodeWith(node.IDEQ(nodeID)),
	}

	if len(names) > 0 {
		predicates = append(predicates, annotation.NameIn(names...))
	}

	annotations, err := backend.client.Annotation.Query().Where(predicates...).All(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying annotations: %w", err)
	}

	return annotations, nil
}

// GetNodesByAnnotation gets all nodes having the specified named
// annotation, limited to a set of annotation values if specified.
func (backend *Backend) GetNodesByAnnotation(name string, values ...string) ([]*sbom.Node, error) {
	if backend.client == nil {
		return nil, errUninitializedClient
	}

	predicates := []predicate.Node{
		node.HasAnnotationsWith(annotation.NameEQ(name)),
	}

	if len(values) > 0 {
		predicates = append(predicates, node.HasAnnotationsWith(annotation.ValueIn(values...)))
	}

	ids, err := backend.client.Node.Query().Where(predicates...).QueryNodes().IDs(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying nodes table: %w", err)
	}

	if len(ids) == 0 {
		return []*sbom.Node{}, nil
	}

	return backend.GetNodesByID(ids...)
}

// GetNodeUniqueAnnotation gets the value for a unique annotation.
func (backend *Backend) GetNodeUniqueAnnotation(nodeID, name string) (string, error) {
	if backend.client == nil {
		return "", errUninitializedClient
	}

	result, err := backend.client.Annotation.Query().
		Where(
			annotation.HasNodeWith(node.IDEQ(nodeID)),
			annotation.NameEQ(name),
			annotation.IsUniqueEQ(true),
		).
		Only(backend.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return "", nil
		}

		return "", fmt.Errorf("retrieving unique annotation for node: %w", err)
	}

	return result.Value, nil
}

// RemoveDocumentAnnotations removes all annotations with the specified name from
// the document, limited to a set of annotation values if specified.
func (backend *Backend) RemoveDocumentAnnotations(documentID, name string, values ...string) error {
	return backend.withTx(
		func(tx *ent.Tx) error {
			predicates := []predicate.Annotation{
				annotation.HasDocumentWith(document.MetadataIDEQ(documentID)),
				annotation.NameEQ(name),
			}

			if len(values) > 0 {
				predicates = append(predicates, annotation.ValueIn(values...))
			}

			if _, err := tx.Annotation.Delete().Where(predicates...).Exec(backend.ctx); err != nil {
				return fmt.Errorf("removing annotations: %w", err)
			}

			return nil
		})
}

// RemoveNodeAnnotations removes all annotations with the specified name from
// the node, limited to a set of annotation values if specified.
func (backend *Backend) RemoveNodeAnnotations(nodeID, name string, values ...string) error {
	return backend.withTx(
		func(tx *ent.Tx) error {
			predicates := []predicate.Annotation{
				annotation.HasNodeWith(node.IDEQ(nodeID)),
				annotation.NameEQ(name),
			}

			if len(values) > 0 {
				predicates = append(predicates, annotation.ValueIn(values...))
			}

			if _, err := tx.Annotation.Delete().Where(predicates...).Exec(backend.ctx); err != nil {
				return fmt.Errorf("removing annotations: %w", err)
			}

			return nil
		})
}

// SetDocumentAnnotations explicitly sets the named annotations for the specified document.
func (backend *Backend) SetDocumentAnnotations(documentID, name string, values ...string) error {
	if err := backend.ClearDocumentAnnotations(documentID); err != nil {
		return err
	}

	return backend.AddDocumentAnnotations(documentID, name, values...)
}

// SetDocumentUniqueAnnotation sets a named annotation value that is unique to the specified document.
func (backend *Backend) SetDocumentUniqueAnnotation(documentID, name, value string) error {
	documentUUID, err := backend.client.Document.Query().
		Where(document.MetadataIDEQ(documentID)).
		OnlyID(backend.ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return backend.withTx(
		backend.saveAnnotations(&ent.Annotation{
			DocumentID: documentUUID,
			Name:       name,
			Value:      value,
			IsUnique:   true,
		}),
	)
}

// SetNodeAnnotations explicitly sets the named annotations for the specified node.
func (backend *Backend) SetNodeAnnotations(nodeID, name string, values ...string) error {
	if err := backend.ClearNodeAnnotations(nodeID); err != nil {
		return err
	}

	return backend.AddNodeAnnotations(nodeID, name, values...)
}

// SetNodeUniqueAnnotation sets a named annotation value that is unique to the specified node.
func (backend *Backend) SetNodeUniqueAnnotation(nodeID, name, value string) error {
	return backend.withTx(
		backend.saveAnnotations(&ent.Annotation{
			NodeID:   &nodeID,
			Name:     name,
			Value:    value,
			IsUnique: true,
		}),
	)
}
