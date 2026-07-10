// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package clickhouse

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"
	"google.golang.org/protobuf/proto"
)

// ErrMissingDocument is returned by Retrieve when no document matches the id.
var ErrMissingDocument = errors.New("document not found")

// Retrieve implements the storage.Retriever interface. It reads the serialized
// document blob stored by Store and unmarshals it, so the returned document is a
// lossless copy of the one originally stored.
//
// FINAL collapses any not-yet-merged ReplacingMergeTree duplicates so the latest
// stored version of the document is returned.
func (backend *Backend) Retrieve(id string, _ *storage.RetrieveOptions) (*sbom.Document, error) {
	if backend.conn == nil {
		return nil, errUninitializedClient
	}

	const query = "SELECT document FROM documents FINAL WHERE id = ? LIMIT 1"

	var blob []byte
	if err := backend.conn.QueryRow(backend.ctx, query, id).Scan(&blob); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: %q", ErrMissingDocument, id)
		}

		return nil, fmt.Errorf("querying document %q: %w", id, err)
	}

	doc := &sbom.Document{}
	if err := proto.Unmarshal(blob, doc); err != nil {
		return nil, fmt.Errorf("unmarshaling document %q: %w", id, err)
	}

	return doc, nil
}
