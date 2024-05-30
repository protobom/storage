// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type IdentifiersEntry struct {
	ent.Schema
}

func (IdentifiersEntry) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("software_identifier_type").Values(
			"UNKNOWN_IDENTIFIER_TYPE",
			"PURL",
			"CPE22",
			"CPE23",
			"GITOID",
		),
		field.String("software_identifier_value"),
	}
}

func (IdentifiersEntry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("node", Node.Type).Ref("identifiers").Unique(),
	}
}

func (IdentifiersEntry) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("software_identifier_type", "software_identifier_value").Unique(),
	}
}
