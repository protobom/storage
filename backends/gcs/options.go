// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package gcs

import (
	"errors"

	"google.golang.org/api/option"
)

var (
	// errUninitializedClient is returned when the backend is used before its
	// client has been initialized with InitClient.
	errUninitializedClient = errors.New("backend client must be initialized")

	// errBucketRequired is returned by InitClient when no bucket is configured.
	errBucketRequired = errors.New("a bucket name is required")
)

type (
	// BackendOptions contains options specific to the protobom GCS backend.
	// Field order is chosen to satisfy the fieldalignment linter.
	BackendOptions struct {
		// Bucket is the name of the GCS bucket that holds the objects.
		Bucket string

		// Prefix is an optional object-key prefix applied to every key.
		Prefix string

		// ProjectID is the GCP project used when CreateBucket is set.
		ProjectID string

		// ClientOptions are passed through to the underlying storage client
		// (credentials, endpoints, ...). Leave empty to use default credentials.
		ClientOptions []option.ClientOption

		// CreateBucket creates the bucket on InitClient if it does not exist.
		CreateBucket bool
	}

	// Option represents a single configuration option for the GCS backend.
	Option func(*Backend)
)

// NewBackendOptions creates a new, empty BackendOptions.
func NewBackendOptions() *BackendOptions {
	return &BackendOptions{}
}

// WithBucket sets the GCS bucket that holds the objects.
func WithBucket(name string) Option {
	return func(backend *Backend) {
		backend.Options.Bucket = name
	}
}

// WithPrefix sets an object-key prefix applied to every stored key.
func WithPrefix(prefix string) Option {
	return func(backend *Backend) {
		backend.Options.Prefix = prefix
	}
}

// WithProjectID sets the GCP project used when creating the bucket.
func WithProjectID(projectID string) Option {
	return func(backend *Backend) {
		backend.Options.ProjectID = projectID
	}
}

// WithCreateBucket configures whether InitClient creates the bucket if missing.
func WithCreateBucket(create bool) Option {
	return func(backend *Backend) {
		backend.Options.CreateBucket = create
	}
}

// WithClientOptions sets the options passed to the underlying storage client.
func WithClientOptions(opts ...option.ClientOption) Option {
	return func(backend *Backend) {
		backend.Options.ClientOptions = opts
	}
}
