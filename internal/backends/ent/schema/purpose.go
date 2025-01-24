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

type Purpose struct {
	ent.Schema
}

func (Purpose) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("primary_purpose").Values(enumValues(new(sbom.Purpose))...),
	}
}

func (Purpose) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("documents", Document.Type).
			Ref("purposes").
			Required().
			Immutable(),
		edge.From("nodes", Node.Type).
			Ref("primary_purpose").
			Required(),
	}
}
