// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent_test

import (
	"testing"

	"github.com/protobom/storage/backends/ent"
	"github.com/stretchr/testify/assert"
)

func TestDialectConfiguration(t *testing.T) {
	tests := []struct {
		name        string
		dialect     ent.DatabaseDialect
		databaseURL string
		expectInit  bool
	}{
		{
			name:        "SQLite in-memory",
			dialect:     ent.SQLiteDialect,
			databaseURL: ":memory:",
			expectInit:  true,
		},
		{
			name:        "SQLite file",
			dialect:     ent.SQLiteDialect,
			databaseURL: "/tmp/test.db",
			expectInit:  true,
		},
		{
			name:        "PostgreSQL invalid connection",
			dialect:     ent.PostgresDialect,
			databaseURL: "invalid-connection-string",
			expectInit:  false,
		},
		{
			name:        "Unsupported dialect",
			dialect:     ent.DatabaseDialect("mysql"),
			databaseURL: "mysql://localhost/test",
			expectInit:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backend := ent.NewBackend(
				ent.WithDialect(tt.dialect),
				ent.WithDatabaseURL(tt.databaseURL),
			)

			err := backend.InitClient()
			if tt.expectInit {
				assert.NoError(t, err, "Expected no error for dialect %s with URL %s", tt.dialect, tt.databaseURL)
				if err == nil {
					backend.CloseClient()
				}
			} else {
				assert.Error(t, err, "Expected error for dialect %s with URL %s", tt.dialect, tt.databaseURL)
			}
		})
	}
}

func TestBackendOptions(t *testing.T) {
	t.Run("Default options", func(t *testing.T) {
		opts := ent.NewBackendOptions()
		assert.Equal(t, ent.SQLiteDialect, opts.Dialect)
		assert.Equal(t, ":memory:", opts.DatabaseURL)
		assert.False(t, opts.Debug)
	})

	t.Run("PostgreSQL helper", func(t *testing.T) {
		backend := ent.NewBackend(
			ent.WithPostgresConnection("postgres://localhost/test"),
		)
		assert.Equal(t, ent.PostgresDialect, backend.Options.Dialect)
		assert.Equal(t, "postgres://localhost/test", backend.Options.DatabaseURL)
	})

	t.Run("Individual options", func(t *testing.T) {
		backend := ent.NewBackend(
			ent.WithDialect(ent.PostgresDialect),
			ent.WithDatabaseURL("postgres://localhost/mydb"),
			ent.Debug(),
		)
		assert.Equal(t, ent.PostgresDialect, backend.Options.Dialect)
		assert.Equal(t, "postgres://localhost/mydb", backend.Options.DatabaseURL)
		assert.True(t, backend.Options.Debug)
	})

	t.Run("Backward compatibility", func(t *testing.T) {
		backend := ent.NewBackend(
			ent.WithDatabaseFile("/tmp/test.db"),
		)
		assert.Equal(t, ent.SQLiteDialect, backend.Options.Dialect)
		assert.Equal(t, "/tmp/test.db", backend.Options.DatabaseURL)
	})
}
