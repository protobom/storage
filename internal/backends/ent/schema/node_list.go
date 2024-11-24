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
	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
)

type NodeList struct {
	ent.Schema
}

func (NodeList) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ProtoMessageMixin[*sbom.NodeList]{},
		UUIDMixin{},
	}
}

func (NodeList) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("document_id", uuid.UUID{}).Unique().Immutable(),
		field.Strings("root_elements"),
	}
}

func (NodeList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nodes", Node.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("document", Document.Type).
			Ref("node_list").
			Required().
			Unique().
			Immutable().
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Field("document_id"),
	}
}
