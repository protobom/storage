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

type Metadata struct {
	ent.Schema
}

func (Metadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Immutable(),
		field.String("version"),
		field.String("name"),
		field.Time("date"),
		field.String("comment"),
	}
}

func (Metadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tools", Tool.Type),
		edge.To("authors", Person.Type),
		edge.To("document_types", DocumentType.Type),
		edge.From("document", Document.Type).Ref("metadata").Required().Unique().Immutable(),
	}
}

func (Metadata) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "version", "name").Unique(),
	}
}
