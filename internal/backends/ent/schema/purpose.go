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
	"github.com/protobom/protobom/pkg/sbom"
)

type Purpose struct {
	ent.Schema
}

func (Purpose) Fields() []ent.Field {
	values := []string{}
	for idx := range len(sbom.Purpose_name) {
		values = append(values, sbom.Purpose_name[int32(idx)])
	}

	return []ent.Field{
		field.String("node_id").Optional(),
		field.Enum("primary_purpose").Values(values...),
	}
}

func (Purpose) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("node", Node.Type).
			Ref("primary_purpose").
			Unique().
			Field("node_id"),
	}
}

func (Purpose) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("node_id", "primary_purpose").
			Unique().
			StorageKey("idx_purposes"),
	}
}
