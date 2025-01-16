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

type IdentifiersEntry struct {
	ent.Schema
}

func (IdentifiersEntry) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DocumentMixin{},
		UUIDMixin{},
	}
}

func (IdentifiersEntry) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values(enumValues(new(sbom.SoftwareIdentifierType))...),
		field.String("value"),
	}
}

func (IdentifiersEntry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("nodes", Node.Type).Ref("identifiers"),
	}
}

func (IdentifiersEntry) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type", "value").
			Unique().
			StorageKey("idx_identifiers"),
	}
}
