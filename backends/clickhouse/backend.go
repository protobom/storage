// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package clickhouse

import (
	"context"
	"errors"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/protobom/protobom/pkg/storage"
)

// bootstrapDatabase is the always-present ClickHouse database used to create the
// backend's target database before connecting to it.
const bootstrapDatabase = "default"

// Backend implements the protobom.pkg.storage.Backend interface backed by ClickHouse.
type Backend struct {
	conn driver.Conn
	ctx  context.Context

	// Options is the set of options for the clickhouse Backend.
	Options *BackendOptions
}

var _ storage.Backend = (*Backend)(nil)

// NewBackend creates a new clickhouse Backend, applying the given options on top
// of the defaults.
func NewBackend(opts ...Option) *Backend {
	backend := &Backend{
		Options: NewBackendOptions(),
	}

	for _, opt := range opts {
		opt(backend)
	}

	return backend
}

// Conn exposes the underlying ClickHouse connection for advanced queries.
func (backend *Backend) Conn() driver.Conn {
	return backend.conn
}

// clickhouseOptions builds the driver options for a connection to the given
// database using the backend's configured address, credentials and settings.
func (backend *Backend) clickhouseOptions(database string) *clickhouse.Options {
	return &clickhouse.Options{
		Addr: backend.Options.Addr,
		Auth: clickhouse.Auth{
			Database: database,
			Username: backend.Options.Username,
			Password: backend.Options.Password,
		},
		TLS:      backend.Options.TLS,
		Settings: clickhouse.Settings(backend.Options.Settings),
		Debug:    backend.Options.Debug,
	}
}

// InitClient connects to ClickHouse, ensures the target database exists and
// applies the backend's schema.
func (backend *Backend) InitClient() error {
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
	}

	ctx := context.Background()

	if err := backend.ensureDatabase(ctx); err != nil {
		return err
	}

	conn, err := clickhouse.Open(backend.clickhouseOptions(backend.Options.Database))
	if err != nil {
		return fmt.Errorf("connecting to clickhouse database %q: %w", backend.Options.Database, err)
	}

	if err := conn.Ping(ctx); err != nil {
		return errors.Join(fmt.Errorf("pinging clickhouse: %w", err), closeConn(conn))
	}

	backend.conn = conn
	backend.ctx = ctx

	return backend.createTables(ctx)
}

// ensureDatabase creates the backend's target database if it does not yet exist,
// connecting through the always-present bootstrap database.
func (backend *Backend) ensureDatabase(ctx context.Context) (err error) {
	conn, err := clickhouse.Open(backend.clickhouseOptions(bootstrapDatabase))
	if err != nil {
		return fmt.Errorf("connecting to clickhouse: %w", err)
	}

	defer func() {
		err = errors.Join(err, closeConn(conn))
	}()

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", backend.Options.Database)
	if execErr := conn.Exec(ctx, query); execErr != nil {
		return fmt.Errorf("creating database %q: %w", backend.Options.Database, execErr)
	}

	return nil
}

// CloseClient closes the underlying ClickHouse connection.
func (backend *Backend) CloseClient() error {
	if backend.conn == nil {
		return nil
	}

	return closeConn(backend.conn)
}

// closeConn closes a ClickHouse connection, wrapping any resulting error.
func closeConn(conn driver.Conn) error {
	if err := conn.Close(); err != nil {
		return fmt.Errorf("closing clickhouse connection: %w", err)
	}

	return nil
}

// Debug enables debug output on the backend's client.
func (backend *Backend) Debug() *Backend {
	backend.Options.Debug = true

	return backend
}

// WithBackendOptions replaces the backend's options with opts.
func (backend *Backend) WithBackendOptions(opts *BackendOptions) *Backend {
	backend.Options = opts

	return backend
}

// WithAddr sets the ClickHouse node addresses to connect to.
func (backend *Backend) WithAddr(addr ...string) *Backend {
	backend.Options.Addr = addr

	return backend
}

// WithDatabase sets the ClickHouse database that holds the protobom tables.
func (backend *Backend) WithDatabase(database string) *Backend {
	backend.Options.Database = database

	return backend
}

// WithCredentials sets the username and password used to authenticate.
func (backend *Backend) WithCredentials(username, password string) *Backend {
	backend.Options.Username = username
	backend.Options.Password = password

	return backend
}
