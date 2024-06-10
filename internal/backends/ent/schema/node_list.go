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
)

type NodeList struct {
	ent.Schema
}

func (NodeList) Fields() []ent.Field {
	return []ent.Field{
		field.String("document_id"),
		field.Strings("root_elements"),
	}
}

func (NodeList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nodes", Node.Type),
		edge.From("document", Document.Type).Ref("node_list").Required().Unique().Field("document_id"),
	}
}

func (NodeList) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("document_id", "root_elements").
			Unique().
			Annotations(
				entsql.IndexWhere("root_elements IS NOT NULL"),
			),
	}
}
