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
)

type Document struct {
	ent.Schema
}

func (Document) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UUIDMixin{},
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
	const edgeRef = "document"

	return []ent.Edge{
		edge.To("annotations", Annotation.Type).
			StructTag(`json:"-"`).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("metadata", Metadata.Type).
			Ref(edgeRef).
			Unique().
			Immutable().
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Field("metadata_id"),
		edge.From("node_list", NodeList.Type).
			Ref(edgeRef).
			Unique().
			Immutable().
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Field("node_list_id"),
	}
}
