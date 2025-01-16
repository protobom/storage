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
	"github.com/protobom/protobom/pkg/sbom"
)

type Metadata struct {
	ent.Schema
}

func (Metadata) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ProtoMessageMixin[*sbom.Metadata]{},
		UUIDMixin{},
	}
}

func (Metadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("native_id").NotEmpty().Immutable(),
		field.String("version"),
		field.String("name"),
		field.Time("date"),
		field.String("comment"),
	}
}

func (Metadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("document", Document.Type).
			Required().
			Unique().
			Immutable(),
		edge.To("tools", Tool.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("authors", Person.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("document_types", DocumentType.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("source_data", SourceData.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (Metadata) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("native_id", "version", "name").
			Unique().
			StorageKey("idx_metadata"),
	}
}
