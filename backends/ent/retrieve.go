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
)

var (
	errMultipleDocuments = errors.New("multiple documents matching ID")
	errMissingDocument   = errors.New("no documents matching IDs")
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
		err = fmt.Errorf("%w %s", errMissingDocument, id)
	case len(documents) > 1:
		err = fmt.Errorf("%w %s", errMultipleDocuments, id)
	default:
		doc = documents[0]
	}

	return
}

func (backend *Backend) GetDocumentsByID(ids ...string) ([]*sbom.Document, error) {
	documents := []*sbom.Document{}
	query := backend.client.Document.Query()

	if len(ids) > 0 {
		query.Where(document.MetadataIDIn(ids...))
	}

	results, err := query.All(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying documents table: %w", err)
	}

	for idx := range results {
		documents = append(documents, results[idx].ProtoMessage)
	}

	return documents, nil
}

func (backend *Backend) GetExternalReferencesByDocumentID(
	id string, types ...string,
) ([]*sbom.ExternalReference, error) {
	query := backend.client.ExternalReference.Query().
		Select("proto_message").
		Where(externalreference.HasDocumentWith(document.MetadataIDEQ(id)))

	typeValues := []any{}
	for idx := range types {
		typeValues = append(typeValues, sbom.ExternalReference_ExternalReferenceType_value[types[idx]])
	}

	if len(typeValues) > 0 {
		query.Where(
			func(s *sql.Selector) {
				s.Where(sqljson.ValueIn("proto_message", typeValues, sqljson.Path("type")))
			},
		)
	}

	results, err := query.All(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying external references: %w", err)
	}

	extRefs := []*sbom.ExternalReference{}
	for _, result := range results {
		extRefs = append(extRefs, result.ProtoMessage)
	}

	return extRefs, nil
}
