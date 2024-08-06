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
)

type Annotation struct {
	ent.Schema
}

func (Annotation) Fields() []ent.Field {
	return []ent.Field{
		field.String("document_id"),
		field.String("name"),
		field.String("value"),
	}
}

func (Annotation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("document", Document.Type).
			Ref("annotations").
			Required().
			Unique().
			Field("document_id"),
	}
}

func (Annotation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("document_id", "name", "value").
			Unique().
			StorageKey("idx_annotation"),
		index.Fields("document_id", "name").
			Unique().
			Annotations(entsql.IndexWhere("name = 'alias'")).
			StorageKey("idx_document_alias"),
	}
}
