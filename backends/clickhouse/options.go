// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package clickhouse

import (
	"crypto/tls"
	"errors"
)

var (
	// errUninitializedClient is returned when the backend is used before its
	// client has been initialized with InitClient.
	errUninitializedClient = errors.New("backend client must be initialized")

	// errNotImplemented is a placeholder returned by stubs until they are
	// implemented in a following step.
	errNotImplemented = errors.New("not implemented")
)

// Default connection settings used when a BackendOptions field is left unset.
const (
	defaultAddr     = "localhost:9000"
	defaultDatabase = "protobom"
	defaultUsername = "default"
)

type (
	// BackendOptions contains options specific to the protobom clickhouse backend.
	// Field order is chosen to satisfy the fieldalignment linter.
	BackendOptions struct {
		// Database is the ClickHouse database that holds the protobom tables.
		Database string

		// Username and Password are the credentials used to authenticate.
		Username string
		Password string

		// Settings are additional ClickHouse settings passed through on connect.
		Settings map[string]any

		// TLS, when set, enables a secure connection to the server.
		TLS *tls.Config

		// Addr is the list of ClickHouse node addresses (host:port) to connect to.
		Addr []string

		// Debug configures the client to output debug information.
		Debug bool
	}

	// Option represents a single configuration option for the clickhouse backend.
	Option func(*Backend)
)

// NewBackendOptions creates a new BackendOptions with sensible defaults for a
// local single-node ClickHouse server.
func NewBackendOptions() *BackendOptions {
	return &BackendOptions{
		Addr:     []string{defaultAddr},
		Database: defaultDatabase,
		Username: defaultUsername,
	}
}

func WithBackendOptions(opts *BackendOptions) Option {
	return func(backend *Backend) {
		backend.WithBackendOptions(opts)
	}
}

func WithAddr(addr ...string) Option {
	return func(backend *Backend) {
		backend.WithAddr(addr...)
	}
}

func WithDatabase(database string) Option {
	return func(backend *Backend) {
		backend.WithDatabase(database)
	}
}

func WithCredentials(username, password string) Option {
	return func(backend *Backend) {
		backend.WithCredentials(username, password)
	}
}

func Debug() Option {
	return func(backend *Backend) {
		backend.Debug()
	}
}
