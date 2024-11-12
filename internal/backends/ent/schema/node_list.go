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
	"github.com/protobom/protobom/pkg/sbom"
)

type NodeList struct {
	ent.Schema
}

func (NodeList) Mixin() []ent.Mixin {
	return []ent.Mixin{
		OnDeleteCascadeMixin{},
		ProtoMessageMixin[*sbom.NodeList]{},
		UUIDMixin{},
	}
}

func (NodeList) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("root_elements"),
	}
}

func (NodeList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nodes", Node.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("document", Document.Type).
			Required().
			Unique().
			Immutable().
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
