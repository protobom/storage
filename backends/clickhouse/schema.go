// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package clickhouse

import (
	"context"
	"fmt"
	"strings"

	"github.com/protobom/storage/schemas"
)

// createTables applies the embedded ClickHouse DDL, creating the backend's
// tables if they do not already exist.
func (backend *Backend) createTables(ctx context.Context) error {
	for _, stmt := range splitStatements(schemas.ClickHouse) {
		if err := backend.conn.Exec(ctx, stmt); err != nil {
			return fmt.Errorf("creating clickhouse schema: %w", err)
		}
	}

	return nil
}

// splitStatements breaks a multi-statement DDL string into individual,
// executable statements. Comments are stripped first (a `--` comment runs to the
// end of its line) so that a semicolon inside a comment cannot split a
// statement; the comment-free text is then split on semicolons.
func splitStatements(ddl string) []string {
	lines := strings.Split(ddl, "\n")
	for idx, line := range lines {
		if before, _, found := strings.Cut(line, "--"); found {
			lines[idx] = before
		}
	}

	parts := strings.Split(strings.Join(lines, "\n"), ";")
	statements := make([]string, 0, len(parts))

	for _, raw := range parts {
		if stmt := strings.TrimSpace(raw); stmt != "" {
			statements = append(statements, stmt)
		}
	}

	return statements
}
