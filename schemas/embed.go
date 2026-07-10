// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

// Package schemas embeds the portable protobom SQL schema files so backends can
// apply them without a runtime file dependency. The .sql files under sql/ remain
// the single source of truth and are also published for use with other engines.
package schemas

import _ "embed"

// ClickHouse is the ClickHouse DDL used by the clickhouse backend to create its
// tables. It is a series of `CREATE TABLE IF NOT EXISTS ...` statements
// separated by semicolons.
//
//go:embed sql/schema.clickhouse.sql
var ClickHouse string
