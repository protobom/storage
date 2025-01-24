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
	"github.com/protobom/protobom/pkg/sbom"

	"github.com/protobom/storage/internal/backends/ent/schema/mixin"
)

type SourceData struct {
	ent.Schema
}

func (SourceData) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ProtoMessage[*sbom.SourceData]{},
	}
}

func (SourceData) Fields() []ent.Field {
	return []ent.Field{
		field.String("format"),
		field.Int64("size"),
		field.String("uri").Nillable().Optional(),
	}
}

func (SourceData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hashes", HashesEntry.Type).
			StorageKey(edge.Table("source_data_hashes"), edge.Columns("source_data_id", "hash_entry_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("documents", Document.Type).
			Ref("source_data").
			Required().
			Immutable(),
		edge.From("metadata", Metadata.Type).
			Ref("source_data").
			Required(),
	}
}

func (SourceData) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("format", "size", "uri").
			Unique().
			StorageKey("idx_source_data"),
	}
}
