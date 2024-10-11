// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent

import (
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"

	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/externalreference"
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
	documents := []*sbom.Document{}

	predicates := []predicate.Document{}
	if len(ids) > 0 {
		predicates = append(predicates, document.MetadataIDIn(ids...))
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
		externalreference.HasDocumentWith(document.MetadataIDEQ(id)),
	}

	typeValues := []any{}
	for idx := range types {
		typeValues = append(typeValues, sbom.ExternalReference_ExternalReferenceType_value[types[idx]])
	}

	if len(typeValues) > 0 {
		predicates = append(predicates, func(s *sql.Selector) {
			s.Where(sqljson.ValueIn("proto_message", typeValues, sqljson.Path("type")))
		})
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
