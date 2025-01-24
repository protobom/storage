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

type DocumentType struct {
	ent.Schema
}

func (DocumentType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ProtoMessage[*sbom.DocumentType]{},
	}
}

func (DocumentType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values(enumValues(new(sbom.DocumentType_SBOMType))...).
			Optional().
			Nillable(),
		field.String("name").Optional().Nillable(),
		field.String("description").Optional().Nillable(),
	}
}

func (DocumentType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("documents", Document.Type).
			Ref("document_types").
			Required().
			Immutable(),
		edge.From("metadata", Metadata.Type).
			Ref("document_types").
			Required(),
	}
}

func (DocumentType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type", "name", "description").
			Unique().
			StorageKey("idx_document_types"),
	}
}
