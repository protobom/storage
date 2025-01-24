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

	"github.com/protobom/storage/internal/backends/ent/schema/mixin"
)

type Document struct {
	ent.Schema
}

func (Document) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.UUID{},
	}
}

func (Document) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("metadata_id", uuid.UUID{}).
			Unique().
			Immutable().
			Optional(),
		field.UUID("node_list_id", uuid.UUID{}).
			Unique().
			Immutable().
			Optional(),
	}
}

func (Document) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("annotations", Annotation.Type).
			StructTag(`json:"-"`).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("metadata", Metadata.Type).
			Unique().
			Immutable().
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Field("metadata_id"),
		edge.To("node_list", NodeList.Type).
			Unique().
			Immutable().
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Field("node_list_id"),
		edge.To("document_types", DocumentType.Type),
		edge.To("edge_types", EdgeType.Type),
		edge.To("external_references", ExternalReference.Type),
		edge.To("hashes", HashesEntry.Type),
		edge.To("identifiers", IdentifiersEntry.Type),
		edge.To("nodes", Node.Type),
		edge.To("persons", Person.Type),
		edge.To("properties", Property.Type),
		edge.To("purposes", Purpose.Type),
		edge.To("source_data", SourceData.Type),
		edge.To("tools", Tool.Type),
	}
}

func (Document) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("metadata").StorageKey("idx_documents_metadata_id"),
		index.Edges("node_list").StorageKey("idx_documents_node_list_id"),
	}
}
