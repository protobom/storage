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

type SourceData struct {
	ent.Schema
}

func (SourceData) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DocumentMixin{},
		ProtoMessageMixin[*sbom.SourceData]{},
		UUIDMixin{},
	}
}

func (SourceData) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("metadata_id", uuid.UUID{}),
		field.String("format"),
		field.Int64("size"),
		field.String("uri").Nillable().Optional(),
		field.JSON("hashes", map[int32]string{}).Optional(),
	}
}

func (SourceData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metadata", Metadata.Type).
			Ref("source_data").
			Required().
			Unique().
			Field("metadata_id"),
	}
}

func (SourceData) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("format", "size", "uri").
			Unique().
			StorageKey("idx_source_data"),
	}
}
