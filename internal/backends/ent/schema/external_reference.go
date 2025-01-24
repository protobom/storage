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

	"github.com/protobom/storage/internal/backends/ent/schema/mixin"
)

type ExternalReference struct {
	ent.Schema
}

func (ExternalReference) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ProtoMessage[*sbom.ExternalReference]{},
	}
}

func (ExternalReference) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
		field.String("comment"),
		field.String("authority").Optional(),
		field.Enum("type").
			Values(enumValues(new(sbom.ExternalReference_ExternalReferenceType))...),
	}
}

func (ExternalReference) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hashes", HashesEntry.Type).
			StorageKey(edge.Table("ext_ref_hashes"), edge.Columns("ext_ref_id", "hash_entry_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("documents", Document.Type).
			Ref("external_references").
			Required().
			Immutable(),
		edge.From("nodes", Node.Type).
			Ref("external_references").
			Required(),
	}
}
