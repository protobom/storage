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
	"io"

	gstorage "cloud.google.com/go/storage"
	"google.golang.org/api/iterator"

	"github.com/protobom/storage/backends/objectstore"
)

// gcsStore adapts a GCS bucket to the objectstore.ObjectStore interface.
type gcsStore struct {
	bucket *gstorage.BucketHandle
}

var _ objectstore.ObjectStore = (*gcsStore)(nil)

// Put writes data to the object at key, overwriting any existing object.
func (s *gcsStore) Put(ctx context.Context, key string, data []byte) error {
	writer := s.bucket.Object(key).NewWriter(ctx)

	if _, err := writer.Write(data); err != nil {
		return errors.Join(fmt.Errorf("writing object %q: %w", key, err), writer.Close())
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("closing object %q: %w", key, err)
	}

	return nil
}

// Get returns the object at key, or objectstore.ErrObjectNotExist if missing.
func (s *gcsStore) Get(ctx context.Context, key string) ([]byte, error) {
	reader, err := s.bucket.Object(key).NewReader(ctx)
	if err != nil {
		if errors.Is(err, gstorage.ErrObjectNotExist) {
			return nil, fmt.Errorf("object %q: %w", key, objectstore.ErrObjectNotExist)
		}

		return nil, fmt.Errorf("opening object %q: %w", key, err)
	}

	data, readErr := io.ReadAll(reader)
	if closeErr := reader.Close(); closeErr != nil && readErr == nil {
		readErr = closeErr
	}

	if readErr != nil {
		return nil, fmt.Errorf("reading object %q: %w", key, readErr)
	}

	return data, nil
}

// Exists reports whether an object exists at key.
func (s *gcsStore) Exists(ctx context.Context, key string) (bool, error) {
	_, err := s.bucket.Object(key).Attrs(ctx)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, gstorage.ErrObjectNotExist) {
		return false, nil
	}

	return false, fmt.Errorf("checking object %q: %w", key, err)
}

// List returns the keys of every object whose name starts with prefix.
func (s *gcsStore) List(ctx context.Context, prefix string) ([]string, error) {
	iter := s.bucket.Objects(ctx, &gstorage.Query{Prefix: prefix})
	keys := []string{}

	for {
		attrs, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("listing objects under %q: %w", prefix, err)
		}

		keys = append(keys, attrs.Name)
	}

	return keys, nil
}
