// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"

	"github.com/protobom/storage/internal/backends/ent/schema/mixin"
)

type EdgeType struct {
	ent.Schema
}

func (EdgeType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ProtoMessage[*sbom.Edge]{},
	}
}

func (EdgeType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values(enumValues(new(sbom.Edge_Type))...),
		field.UUID("node_id", uuid.UUID{}),
		field.UUID("to_node_id", uuid.UUID{}),
	}
}

func (EdgeType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("from", Node.Type).
			Required().
			Unique().
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Field("node_id"),
		edge.To("to", Node.Type).
			Required().
			Unique().
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Field("to_node_id"),
		edge.From("documents", Document.Type).
			Ref("edge_types").
			Required().
			Immutable(),
		edge.From("node_lists", NodeList.Type).
			Ref("edge_types"),
	}
}

func (EdgeType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type", "node_id", "to_node_id").
			Unique().
			StorageKey("idx_edge_types"),
	}
}
