// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

type Document struct {
	ent.Schema
}

func (Document) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UUIDMixin{},
	}
}

func (Document) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metadata", Metadata.Type).
			Unique().
			Immutable(),
		edge.To("node_list", NodeList.Type).
			Unique().
			Immutable(),
	}
}
