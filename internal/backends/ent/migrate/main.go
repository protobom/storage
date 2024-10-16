//go:build ignore
// +build ignore

// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	sqlite "github.com/glebarez/go-sqlite"

	"github.com/protobom/storage/internal/backends/ent/migrate"
)

const dsn = "sqlite://:memory:?_pragma=foreign_keys(1)"

func main() {
	output, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		log.Fatal("failed getting current Git branch")
	}

	output = regexp.MustCompile(`^\d+-|\n`).ReplaceAll(output, []byte(""))
	migrationName := string(regexp.MustCompile(`^\d+-|[^\w]+`).ReplaceAll(output, []byte("_")))

	// Register the SQLite driver as "sqlite3".
	if !slices.Contains(sql.Drivers(), "sqlite3") {
		sqlite.RegisterAsSQLITE3()
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("getting working directory: %v", err)
	}

	// Normalize behavior whether running directly or with `go generate`.
	relPath := filepath.Join("internal", "backends", "ent")
	cwd = filepath.Join(strings.TrimSuffix(cwd, relPath), relPath)

	ctx := context.Background()
	migrationDir := filepath.Join(cwd, "migrate", "migrations")

	// Create a local migration directory able to understand Atlas migration file format for replay.
	if err := os.MkdirAll(migrationDir, 0755); err != nil {
		log.Fatalf("failed creating migration directory: %v", err)
	}

	localDir, err := atlas.NewLocalDir(migrationDir)
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDialect(dialect.SQLite),
		schema.WithDir(localDir),
		schema.WithFormatter(atlas.DefaultFormatter),
		schema.WithIndent("  "),
		schema.WithMigrationMode(schema.ModeReplay),
	}

	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, dsn, string(migrationName), opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
