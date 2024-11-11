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
	"github.com/google/uuid"
)

type Document struct {
	ent.Schema
}

func (Document) Mixin() []ent.Mixin {
	return []ent.Mixin{
		OnDeleteCascadeMixin{},
		UUIDMixin{},
	}
}

func (Document) Fields() []ent.Field {
	return []ent.Field{
		field.String("metadata_id").
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
		edge.From("metadata", Metadata.Type).
			Ref("document").
			Unique().
			Immutable().
			Field("metadata_id"),
		edge.From("node_list", NodeList.Type).
			Ref("document").
			Unique().
			Immutable().
			Field("node_list_id"),
	}
}

func (Document) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("metadata_id", "node_list_id").
			Unique().
			StorageKey("idx_documents"),
	}
}
