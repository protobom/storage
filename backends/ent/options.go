// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"errors"
)

// Enable SQLite foreign key support.
const dsnParams string = "?_pragma=foreign_keys(1)"

var (
	errInvalidEntOptions   = errors.New("invalid ent backend options")
	errUninitializedClient = errors.New("backend client must be initialized")
)

type (
	// BackendOptions contains options specific to the protobom ent backend.
	BackendOptions struct {
		// DatabaseFile is the file path of the SQLite database to be created.
		DatabaseFile string

		// Debug configures the ent client to output all SQL statements during execution.
		Debug bool
	}

	// Option represents a single configuration option for the ent backend.
	Option func(*Backend)
)

// NewBackendOptions creates a new BackendOptions for the backend.
func NewBackendOptions() *BackendOptions {
	return &BackendOptions{
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
