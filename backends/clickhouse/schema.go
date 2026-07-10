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
// executable statements, dropping blank statements and full-line comments.
// Inline (trailing) comments are preserved; ClickHouse parses them.
func splitStatements(ddl string) []string {
	statements := []string{}

	for _, raw := range strings.Split(ddl, ";") {
		lines := make([]string, 0, strings.Count(raw, "\n")+1)

		for _, line := range strings.Split(raw, "\n") {
			if strings.HasPrefix(strings.TrimSpace(line), "--") {
				continue
			}

			lines = append(lines, line)
		}

		if stmt := strings.TrimSpace(strings.Join(lines, "\n")); stmt != "" {
			statements = append(statements, stmt)
		}
	}

	return statements
}
