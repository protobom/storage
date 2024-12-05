// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"

	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/externalreference"
	"github.com/protobom/storage/internal/backends/ent/metadata"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

var (
	ErrMultipleDocuments = errors.New("multiple documents matching ID")
	ErrMissingDocument   = errors.New("no documents matching IDs")
)

// Retrieve implements the storage.Retriever interface.
func (backend *Backend) Retrieve(id string, _ *storage.RetrieveOptions) (doc *sbom.Document, err error) {
	if backend.client == nil {
		return nil, fmt.Errorf("%w", errUninitializedClient)
	}

	if backend.Options == nil {
		backend.Options = NewBackendOptions()
	}

	switch documents, getDocsErr := backend.GetDocumentsByID(id); {
	case getDocsErr != nil:
		err = fmt.Errorf("querying documents: %w", getDocsErr)
	case len(documents) == 0:
		err = fmt.Errorf("%w %s", ErrMissingDocument, id)
	case len(documents) > 1:
		err = fmt.Errorf("%w %s", ErrMultipleDocuments, id)
	default:
		doc = documents[0]
	}

	return doc, err
}

func (backend *Backend) GetDocumentsByID(ids ...string) ([]*sbom.Document, error) {
	predicates := []predicate.Metadata{}

	if len(ids) > 0 {
		predicates = append(predicates, metadata.NativeIDIn(ids...))
	}

	docUUIDs, err := backend.client.Metadata.Query().
		WithDocument().
		Where(predicates...).
		QueryDocument().
		IDs(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying documents IDs: %w", err)
	}

	// If a id string was passed in, but no matches were found, return nothing.
	if len(ids) > 0 && len(docUUIDs) == 0 {
		return nil, nil
	}

	return backend.GetDocumentsByUUID(docUUIDs...)
}

func (backend *Backend) GetDocumentsByUUID(uuids ...uuid.UUID) ([]*sbom.Document, error) {
	documents := []*sbom.Document{}
	predicates := []predicate.Document{}

	if len(uuids) > 0 {
		predicates = append(predicates, document.IDIn(uuids...))
	}

	results, err := backend.client.Document.Query().
		Where(predicates...).
		WithMetadata().
		WithNodeList().
		All(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying documents table: %w", err)
	}

	for idx := range results {
		documents = append(documents, &sbom.Document{
			Metadata: results[idx].Edges.Metadata.ProtoMessage,
			NodeList: results[idx].Edges.NodeList.ProtoMessage,
		})
	}

	return documents, nil
}

func (backend *Backend) GetExternalReferencesByDocumentID(
	id string, types ...string,
) ([]*sbom.ExternalReference, error) {
	predicates := []predicate.ExternalReference{
		externalreference.HasDocumentWith(document.HasMetadataWith(metadata.NativeIDEQ(id))),
	}

	extRefTypes := []externalreference.Type{}

	for idx := range types {
		extRefType := externalreference.Type(types[idx])
		if err := externalreference.TypeValidator(extRefType); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		extRefTypes = append(extRefTypes, extRefType)
	}

	if len(extRefTypes) > 0 {
		predicates = append(predicates, externalreference.TypeIn(extRefTypes...))
	}

	results, err := backend.client.ExternalReference.Query().
		Select("proto_message").
		Where(predicates...).
		All(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying external references: %w", err)
	}

	extRefs := []*sbom.ExternalReference{}
	for _, result := range results {
		extRefs = append(extRefs, result.ProtoMessage)
	}

	return extRefs, nil
}

func (backend *Backend) GetNodesByID(ids ...string) ([]*sbom.Node, error) {
	nodeUUIDS, err := backend.client.Node.Query().
		Where(node.NativeIDIn(ids...)).
		IDs(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying node IDs: %w", err)
	}

	return backend.GetNodesByUUID(nodeUUIDS...)
}

func (backend *Backend) GetNodesByUUID(uuids ...uuid.UUID) ([]*sbom.Node, error) {
	nodes := []*sbom.Node{}
	predicates := []predicate.Node{}

	if len(uuids) > 0 {
		predicates = append(predicates, node.IDIn(uuids...))
	}

	results, err := backend.client.Node.Query().
		Where(predicates...).
		All(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying nodes table: %w", err)
	}

	for idx := range results {
		nodes = append(nodes, results[idx].ProtoMessage)
	}

	return nodes, nil
}
