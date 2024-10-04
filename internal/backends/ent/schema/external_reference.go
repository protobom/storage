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

type ExternalReference struct {
	ent.Schema
}

func (ExternalReference) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DocumentMixin{},
		ProtoMessageMixin{ProtoMessageType: &sbom.ExternalReference{}},
		UUIDMixin{},
	}
}

func (ExternalReference) Fields() []ent.Field {
	return []ent.Field{
		field.String("node_id").Optional(),
		field.String("url"),
		field.String("comment"),
		field.String("authority").Optional(),
		field.Enum("type").
			Values(enumValues(new(sbom.ExternalReference_ExternalReferenceType))...),
		field.JSON("hashes", map[int32]string{}).Optional(),
	}
}

func (ExternalReference) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("node", Node.Type).Ref("external_references").Unique().Field("node_id"),
	}
}

func (ExternalReference) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("node_id", "url", "type").
			Unique().
			StorageKey("idx_external_references"),
	}
}
