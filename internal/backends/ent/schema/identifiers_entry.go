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
	"github.com/protobom/protobom/pkg/sbom"
)

type IdentifiersEntry struct {
	ent.Schema
}

func (IdentifiersEntry) Fields() []ent.Field {
	values := []string{}
	for idx := range len(sbom.SoftwareIdentifierType_name) {
		values = append(values, sbom.SoftwareIdentifierType_name[int32(idx)])
	}

	return []ent.Field{
		field.String("node_id").Optional(),
		field.Enum("software_identifier_type").Values(values...),
		field.String("software_identifier_value"),
	}
}

func (IdentifiersEntry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("node", Node.Type).
			Ref("identifiers").
			Unique().
			Field("node_id"),
	}
}

func (IdentifiersEntry) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("node_id", "software_identifier_type", "software_identifier_value").
			Unique().
			StorageKey("idx_identifiers_entries"),
	}
}
