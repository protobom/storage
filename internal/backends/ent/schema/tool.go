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
	"github.com/protobom/protobom/pkg/sbom"
)

type Tool struct {
	ent.Schema
}

func (Tool) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DocumentMixin{},
		ProtoMessageMixin[*sbom.Tool]{},
	}
}

func (Tool) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("metadata_id", uuid.UUID{}),
		field.String("name"),
		field.String("version"),
		field.String("vendor"),
	}
}

func (Tool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metadata", Metadata.Type).
			Ref("tools").
			Required().
			Unique().
			Field("metadata_id"),
	}
}

func (Tool) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("metadata_id", "name", "version", "vendor").
			Unique().
			StorageKey("idx_tools"),
	}
}
