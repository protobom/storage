// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"context"
	"database/sql"
	"fmt"
	"slices"

	sqlite "github.com/glebarez/go-sqlite"
	"github.com/protobom/protobom/pkg/storage"

	"github.com/protobom/storage/internal/backends/ent"
)

// Backend implements the protobom.pkg.storage.Backend interface.
type Backend struct {
	// Options is the set of options common to all ent Backends.
	Options *BackendOptions
}

var _ storage.Backend = (*Backend)(nil)

func NewBackend(opts ...Option) *Backend {
	backend := &Backend{
		Options: NewBackendOptions(),
	}

	for _, opt := range opts {
		opt(backend)
	}

	backend.init(backend.Options)

	return backend
}

func (backend *Backend) init(opts *BackendOptions) {
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
	}

	// Register the SQLite driver as "sqlite3".
	if !slices.Contains(sql.Drivers(), "sqlite3") {
		sqlite.RegisterAsSQLITE3()
	}

	clientOpts := []ent.Option{}
	if opts.debug {
		clientOpts = append(clientOpts, ent.Debug())
	}

	client, err := ent.Open("sqlite3", opts.DatabaseFile+dsnParams, clientOpts...)
	if err != nil {
		panic(fmt.Errorf("failed opening connection to sqlite: %w", err))
	}

	opts.client = client
	opts.ctx = ent.NewContext(context.Background(), client)

	// Run the auto migration tool.
	if err := opts.client.Schema.Create(opts.ctx); err != nil {
		panic(fmt.Errorf("failed creating schema resources: %w", err))
	}
}

func (backend *Backend) Debug() *Backend {
	backend.Options.debug = true
	backend.init(backend.Options)

	return backend
}

func (backend *Backend) WithBackendOptions(opts *BackendOptions) *Backend {
	backend.Options = opts
	backend.init(backend.Options)

	return backend
}

func (backend *Backend) WithDatabaseFile(file string) *Backend {
	backend.Options.DatabaseFile = file
	backend.init(backend.Options)

	return backend
}
