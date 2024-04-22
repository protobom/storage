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

type NodeList struct {
	ent.Schema
}

func (NodeList) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SourceDataMixin{},
	}
}

func (NodeList) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("root_elements"),
	}
}

func (NodeList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nodes", Node.Type),
		edge.To("document", Document.Type).Unique(),
	}
}

func (NodeList) Annotations() []schema.Annotation { return nil }
