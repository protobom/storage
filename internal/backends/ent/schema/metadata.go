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
	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"

	"github.com/protobom/storage/internal/backends/ent/schema/mixin"
)

type Metadata struct {
	ent.Schema
}

func (Metadata) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ProtoMessage[*sbom.Metadata]{},
	}
}

func (Metadata) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("source_data_id", uuid.UUID{}).Optional(),
		field.String("native_id").
			NotEmpty().
			Immutable().
			StructTag(`json:"id"`),
		field.String("version"),
		field.String("name"),
		field.Time("date"),
		field.String("comment"),
	}
}

func (Metadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tools", Tool.Type).
			StorageKey(edge.Table("metadata_tools"), edge.Columns("metadata_id", "tool_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("authors", Person.Type).
			StorageKey(edge.Table("metadata_authors"), edge.Columns("metadata_id", "person_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("document_types", DocumentType.Type).
			StorageKey(edge.Table("metadata_document_types"), edge.Columns("metadata_id", "document_type_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("source_data", SourceData.Type).
			Unique().
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Field("source_data_id"),
		edge.From("documents", Document.Type).
			Ref("metadata").
			Required().
			Immutable(),
	}
}
