// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

// Package objectstore implements protobom document storage on top of a generic
// object store (blob storage keyed by name). It is provider agnostic: concrete
// backends such as GCS or S3 supply an ObjectStore adapter and reuse the shared
// document store, key scheme and indexer defined here.
package objectstore

import (
	"context"
	"errors"
)

// ErrObjectNotExist is returned by ObjectStore.Get when the key does not exist.
// Adapters must translate their provider's not-found error into this sentinel.
var ErrObjectNotExist = errors.New("object does not exist")

// ObjectStore is the minimal object-storage abstraction the backend builds on.
// Keys are opaque, slash-delimited names; List matches by key prefix.
type ObjectStore interface {
	// Put writes data at key, overwriting any existing object.
	Put(ctx context.Context, key string, data []byte) error

	// Get returns the object at key, or ErrObjectNotExist if it does not exist.
	Get(ctx context.Context, key string) ([]byte, error)

	// Exists reports whether an object exists at key.
	Exists(ctx context.Context, key string) (bool, error)

	// List returns the keys of every object whose name starts with prefix.
	List(ctx context.Context, prefix string) ([]string, error)
}
