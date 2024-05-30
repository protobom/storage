// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
)

type Document struct {
	ent.Schema
}

func (Document) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metadata", Metadata.Type).Ref("document").Required().Unique(),
		edge.From("node_list", NodeList.Type).Ref("document").Required().Unique(),
	}
}

func (Document) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("metadata", "node_list").Unique(),
	}
}
