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

type Tool struct {
	ent.Schema
}

func (Tool) Fields() []ent.Field {
	return []ent.Field{
		field.String("metadata_id").Optional(),
		field.String("name"),
		field.String("version"),
		field.String("vendor"),
	}
}

func (Tool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metadata", Metadata.Type).Ref("tools").Unique().Field("metadata_id"),
	}
}

func (Tool) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("metadata_id", "name", "version", "vendor").Unique(),
	}
}
