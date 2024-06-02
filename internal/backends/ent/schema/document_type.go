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

type DocumentType struct {
	ent.Schema
}

func (DocumentType) Fields() []ent.Field {
	return []ent.Field{
		field.String("metadata_id").Optional(),
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
		edge.From("metadata", Metadata.Type).Ref("document_types").Unique().Field("metadata_id"),
	}
}

func (DocumentType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("metadata_id", "type", "name", "description").Unique(),
	}
}
