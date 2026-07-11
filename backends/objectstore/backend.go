// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package objectstore

import (
	"context"
	"errors"
	"fmt"

	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"
	"google.golang.org/protobuf/proto"
)

// ErrMissingDocument is returned by Retrieve when no document matches the id.
var ErrMissingDocument = errors.New("document not found")

// Backend implements protobom document storage on top of an ObjectStore. It is
// provider agnostic; concrete backends (GCS, S3, ...) supply the ObjectStore.
type Backend struct {
	store  ObjectStore
	prefix string
}

// NewBackend returns a Backend that reads and writes objects through store,
// placing every key under the given prefix (which may be empty).
func NewBackend(store ObjectStore, prefix string) *Backend {
	return &Backend{store: store, prefix: prefix}
}

// Store serializes doc and writes it as a single object keyed by its document id.
// When opts.NoClobber is set and a document with that id already exists, it is
// left unchanged.
func (backend *Backend) Store(ctx context.Context, doc *sbom.Document, opts *storage.StoreOptions) error {
	id, err := documentID(doc)
	if err != nil {
		return err
	}

	key := documentKey(backend.prefix, id)

	if opts != nil && opts.NoClobber {
		exists, existsErr := backend.store.Exists(ctx, key)
		if existsErr != nil {
			return fmt.Errorf("checking document %q: %w", id, existsErr)
		}

		if exists {
			return nil
		}
	}

	blob, err := proto.MarshalOptions{Deterministic: true}.Marshal(doc)
	if err != nil {
		return fmt.Errorf("marshaling document %q: %w", id, err)
	}

	if err := backend.store.Put(ctx, key, blob); err != nil {
		return fmt.Errorf("storing document %q: %w", id, err)
	}

	return nil
}

// Retrieve reads and deserializes the document stored under id.
func (backend *Backend) Retrieve(ctx context.Context, id string) (*sbom.Document, error) {
	blob, err := backend.store.Get(ctx, documentKey(backend.prefix, id))
	if err != nil {
		if errors.Is(err, ErrObjectNotExist) {
			return nil, fmt.Errorf("%w: %q", ErrMissingDocument, id)
		}

		return nil, fmt.Errorf("retrieving document %q: %w", id, err)
	}

	doc := &sbom.Document{}
	if err := proto.Unmarshal(blob, doc); err != nil {
		return nil, fmt.Errorf("unmarshaling document %q: %w", id, err)
	}

	return doc, nil
}
