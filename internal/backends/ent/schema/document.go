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
	"github.com/protobom/protobom/pkg/sbom"
)

type Document struct {
	ent.Schema
}

func (Document) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ProtoMessageMixin{ProtoMessageType: &sbom.Document{}},
	}
}

func (Document) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
	}
}

func (Document) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metadata", Metadata.Type).
			Unique().
			Immutable().
			StorageKey(edge.Column("id")),
		edge.From("node_list", NodeList.Type).
			Ref("document").
			Unique().
			Immutable(),
	}
}
