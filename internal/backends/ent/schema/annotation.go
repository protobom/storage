// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Annotation struct {
	ent.Schema
}

func (Annotation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DocumentMixin{},
	}
}

func (Annotation) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("value"),
		field.Bool("is_unique").
			Default(false),
	}
}

func (Annotation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("document_id", "name", "value").
			Unique().
			StorageKey("idx_annotations"),
		index.Fields("document_id", "name").
			Unique().
			Annotations(entsql.IndexWhere("is_unique = true")).
			StorageKey("idx_document_unique_annotations"),
	}
}
