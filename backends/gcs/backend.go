// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package gcs

import (
	"context"
	"errors"
	"fmt"

	gstorage "cloud.google.com/go/storage"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"

	"github.com/protobom/storage/backends/objectstore"
)

// Backend implements the protobom/storage.Backend interface backed by Google
// Cloud Storage. It adapts a GCS bucket to the shared objectstore.Backend.
type Backend struct {
	client *gstorage.Client
	inner  *objectstore.Backend
	ctx    context.Context

	// Options is the set of options for the GCS Backend.
	Options *BackendOptions
}

var _ storage.Backend = (*Backend)(nil)

// NewBackend creates a new GCS Backend, applying the given options on top of the
// defaults.
func NewBackend(opts ...Option) *Backend {
	backend := &Backend{
		Options: NewBackendOptions(),
	}

	for _, opt := range opts {
		opt(backend)
	}

	return backend
}

// InitClient connects to Google Cloud Storage, optionally creates the target
// bucket and wires up the shared object-store backend.
func (backend *Backend) InitClient() error {
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
	}

	if backend.Options.Bucket == "" {
		return errBucketRequired
	}

	ctx := context.Background()

	client, err := gstorage.NewClient(ctx, backend.Options.ClientOptions...)
	if err != nil {
		return fmt.Errorf("creating storage client: %w", err)
	}

	bucket := client.Bucket(backend.Options.Bucket)

	if backend.Options.CreateBucket {
		if err := ensureBucket(ctx, bucket, backend.Options.ProjectID); err != nil {
			return errors.Join(err, closeClient(client))
		}
	}

	backend.client = client
	backend.ctx = ctx
	backend.inner = objectstore.NewBackend(&gcsStore{bucket: bucket}, backend.Options.Prefix)

	return nil
}

// CloseClient closes the underlying storage client.
func (backend *Backend) CloseClient() error {
	if backend.client == nil {
		return nil
	}

	return closeClient(backend.client)
}

// Store implements the storage.Storer interface.
func (backend *Backend) Store(doc *sbom.Document, opts *storage.StoreOptions) error {
	if backend.inner == nil {
		return errUninitializedClient
	}

	if err := backend.inner.Store(backend.ctx, doc, opts); err != nil {
		return fmt.Errorf("gcs store: %w", err)
	}

	return nil
}

// Retrieve implements the storage.Retriever interface.
func (backend *Backend) Retrieve(id string, _ *storage.RetrieveOptions) (*sbom.Document, error) {
	if backend.inner == nil {
		return nil, errUninitializedClient
	}

	doc, err := backend.inner.Retrieve(backend.ctx, id)
	if err != nil {
		return nil, fmt.Errorf("gcs retrieve: %w", err)
	}

	return doc, nil
}

// ensureBucket creates the bucket if it does not already exist.
func ensureBucket(ctx context.Context, bucket *gstorage.BucketHandle, projectID string) error {
	_, err := bucket.Attrs(ctx)
	if err == nil {
		return nil
	}

	if !errors.Is(err, gstorage.ErrBucketNotExist) {
		return fmt.Errorf("checking bucket: %w", err)
	}

	if err := bucket.Create(ctx, projectID, nil); err != nil {
		return fmt.Errorf("creating bucket: %w", err)
	}

	return nil
}

// closeClient closes a storage client, wrapping any resulting error.
func closeClient(client *gstorage.Client) error {
	if err := client.Close(); err != nil {
		return fmt.Errorf("closing storage client: %w", err)
	}

	return nil
}
