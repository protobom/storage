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

type Metadata struct {
	ent.Schema
}

func (Metadata) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ProtoMessageMixin{ProtoMessageType: &sbom.Metadata{}},
	}
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
		edge.To("document", Document.Type).
			Required().
			Unique().
			Immutable(),
	}
}

func (Metadata) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "version", "name").
			Unique().
			StorageKey("idx_metadata"),
	}
}
