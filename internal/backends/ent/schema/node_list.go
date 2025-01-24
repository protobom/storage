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
	"github.com/protobom/protobom/pkg/sbom"

	"github.com/protobom/storage/internal/backends/ent/schema/mixin"
)

type NodeList struct {
	ent.Schema
}

func (NodeList) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ProtoMessage[*sbom.NodeList]{},
	}
}

func (NodeList) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("root_elements"),
	}
}

func (NodeList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("edge_types", EdgeType.Type).
			StorageKey(edge.Table("node_list_edges"), edge.Columns("node_list_id", "edge_type_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("nodes", Node.Type).
			StorageKey(edge.Table("node_list_nodes"), edge.Columns("node_list_id", "node_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("documents", Document.Type).
			Ref("node_list").
			Required().
			Immutable(),
	}
}
