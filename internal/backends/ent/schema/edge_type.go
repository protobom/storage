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

type EdgeType struct {
	ent.Schema
}

func (EdgeType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DocumentMixin{},
	}
}

func (EdgeType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values(enumValues(new(sbom.Edge_Type))...),
		field.String("node_id"),
		field.String("to_node_id"),
	}
}

func (EdgeType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("from", Node.Type).Required().Unique().Field("node_id"),
		edge.To("to", Node.Type).Required().Unique().Field("to_node_id"),
	}
}

func (EdgeType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type", "node_id", "to_node_id").
			Unique().
			StorageKey("idx_edge_types"),
	}
}
