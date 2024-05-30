// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"context"

	"github.com/protobom/storage/internal/backends/ent"
)

// Enable SQLite foreign key support.
const dsnParams string = "?_pragma=foreign_keys(1)"

type (
	config struct {
		client *ent.Client
		ctx    context.Context
		debug  bool
	}

	// BackendOptions contains options specific to the protobom ent backend.
	BackendOptions struct {
		*config
		DatabaseFile string
	}

	// Option represents a single configuration option for the ent backend.
	Option func(*Backend)
)

// NewBackendOptions creates a new BackendOptions for the backend.
func NewBackendOptions() *BackendOptions {
	return &BackendOptions{
		config:       &config{},
		DatabaseFile: ":memory:",
	}
}

func WithBackendOptions(opts *BackendOptions) Option {
	return func(backend *Backend) {
		backend.WithBackendOptions(opts)
	}
}

func WithDatabaseFile(file string) Option {
	return func(backend *Backend) {
		backend.WithDatabaseFile(file)
	}
}

func Debug() Option {
	return func(backend *Backend) {
		backend.Debug()
	}
}
