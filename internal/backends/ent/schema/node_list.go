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
)

type NodeList struct {
	ent.Schema
}

func (NodeList) Fields() []ent.Field {
	return []ent.Field{
		field.String("document_id").Unique().Immutable(),
		field.Strings("root_elements"),
	}
}

func (NodeList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nodes", Node.Type),
		edge.From("document", Document.Type).
			Ref("node_list").
			Required().
			Unique().
			Immutable().
			Field("document_id"),
	}
}
