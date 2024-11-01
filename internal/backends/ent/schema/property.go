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

type Property struct {
	ent.Schema
}

func (Property) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DocumentMixin{},
		ProtoMessageMixin[*sbom.Property]{},
		UUIDMixin{},
	}
}

func (Property) Fields() []ent.Field {
	return []ent.Field{
		field.String("node_id").Optional(),
		field.String("name"),
		field.String("data"),
	}
}

func (Property) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("node", Node.Type).
			Ref("properties").
			Unique().
			Field("node_id"),
	}
}

func (Property) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "data").
			Unique().
			StorageKey("idx_property"),
	}
}
