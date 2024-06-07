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
		field.String("metadata_id").Immutable(),
		field.Int("node_list_id").Immutable(),
	}
}

func (Document) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metadata", Metadata.Type).
			Ref("document").
			Required().
			Unique().
			Immutable().
			Field("metadata_id"),
		edge.From("node_list", NodeList.Type).
			Ref("document").
			Required().
			Unique().
			Immutable().
			Field("node_list_id"),
	}
}

func (Document) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("metadata_id").Unique(),
	}
}
