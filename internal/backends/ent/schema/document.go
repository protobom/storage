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

type Document struct {
	ent.Schema
}

func (Document) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Strings("tags").Default([]string{}).Optional(),
	}
}

func (Document) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metadata", Metadata.Type).
			Unique().
			Immutable().
			StorageKey(edge.Column("id")),
		edge.To("node_list", NodeList.Type).
			Unique().
			Immutable(),
	}
}

func (Document) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tags").
			Unique().
			StorageKey("idx_document_tags"),
	}
}
