// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DocumentType struct {
	ent.Schema
}

func (DocumentType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SourceDataMixin{},
	}
}

func (DocumentType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values(
			"OTHER",
			"DESIGN",
			"SOURCE",
			"BUILD",
			"ANALYZED",
			"DEPLOYED",
			"RUNTIME",
			"DISCOVERY",
			"DECOMISSION",
		).Optional().Nillable(),
		field.String("name").Optional().Nillable(),
		field.String("description").Optional().Nillable(),
	}
}

func (DocumentType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metadata", Metadata.Type).Ref("document_types").Unique(),
	}
}

func (DocumentType) Annotations() []schema.Annotation { return nil }
