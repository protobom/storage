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

	"github.com/protobom/storage/internal/backends/ent/schema/mixin"
)

type Tool struct {
	ent.Schema
}

func (Tool) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ProtoMessage[*sbom.Tool]{},
	}
}

func (Tool) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("version"),
		field.String("vendor"),
	}
}

func (Tool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("documents", Document.Type).
			Ref("tools").
			Required().
			Immutable(),
		edge.From("metadata", Metadata.Type).
			Ref("tools").
			Required(),
	}
}

func (Tool) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "version", "vendor").
			Unique().
			StorageKey("idx_tools"),
	}
}
