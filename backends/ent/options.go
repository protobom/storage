// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent

import (
	"errors"

	"github.com/protobom/storage/internal/backends/ent"
)

// Enable SQLite foreign key support.
const dsnParams string = "?_pragma=foreign_keys(1)"

var (
	errInvalidEntOptions   = errors.New("invalid ent backend options")
	errUninitializedClient = errors.New("backend client must be initialized")
	errUnsupportedDialect  = errors.New("unsupported database dialect")
)

type (
	// Annotation is the model entity for the Annotation schema.
	Annotation = ent.Annotation

	// Annotations is a parsable slice of Annotation.
	Annotations = ent.Annotations

	// DatabaseDialect represents the database dialect to use.
	DatabaseDialect string

	// BackendOptions contains options specific to the protobom ent backend.
	BackendOptions struct {
		// DatabaseURL is the database connection string or file path.
		// For SQLite: file path (e.g., ":memory:" or "path/to/file.db")
		// For PostgreSQL: connection string (e.g., "postgres://user:password@host:port/dbname")
		DatabaseURL string

		// Dialect specifies the database dialect to use (sqlite or postgres)
		Dialect DatabaseDialect

		// Annotations is a slice of annotations to apply to stored document.
		Annotations

		// Debug configures the ent client to output all SQL statements during execution.
		Debug bool
	}

	// Option represents a single configuration option for the ent backend.
	Option func(*Backend)
)

const (
	// SQLiteDialect represents SQLite database dialect.
	SQLiteDialect DatabaseDialect = "sqlite"
	// PostgresDialect represents PostgreSQL database dialect.
	PostgresDialect DatabaseDialect = "postgres"
)

// NewBackendOptions creates a new BackendOptions for the backend.
func NewBackendOptions() *BackendOptions {
	return &BackendOptions{
		DatabaseURL: ":memory:",
		Dialect:     SQLiteDialect,
	}
}

func WithBackendOptions(opts *BackendOptions) Option {
	return func(backend *Backend) {
		backend.WithBackendOptions(opts)
	}
}

func WithDatabaseFile(file string) Option {
	return func(backend *Backend) {
		backend.WithDatabaseURL(file)
	}
}

func WithDatabaseURL(url string) Option {
	return func(backend *Backend) {
		backend.WithDatabaseURL(url)
	}
}

func WithDialect(dialect DatabaseDialect) Option {
	return func(backend *Backend) {
		backend.WithDialect(dialect)
	}
}

func WithPostgresConnection(connectionString string) Option {
	return func(backend *Backend) {
		backend.WithDialect(PostgresDialect)
		backend.WithDatabaseURL(connectionString)
	}
}

func Debug() Option {
	return func(backend *Backend) {
		backend.Debug()
	}
}
