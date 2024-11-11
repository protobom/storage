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
	"entgo.io/ent/schema/index"
	"github.com/protobom/protobom/pkg/sbom"
)

type Person struct {
	ent.Schema
}

func (Person) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DocumentMixin{},
		OnDeleteCascadeMixin{},
		ProtoMessageMixin[*sbom.Person]{},
		UUIDMixin{},
	}
}

func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.String("metadata_id").Optional(),
		field.String("node_id").Optional(),
		field.String("name"),
		field.Bool("is_org"),
		field.String("email"),
		field.String("url"),
		field.String("phone"),
	}
}

func (Person) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contacts", Person.Type).From("contact_owner").Unique(),
		edge.From("metadata", Metadata.Type).Ref("authors").Unique().Field("metadata_id"),
		edge.From("node", Node.Type).Ref("suppliers").Unique().Ref("originators").Unique().Field("node_id"),
	}
}

func (Person) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("metadata_id", "name", "is_org", "email", "url", "phone").
			Unique().
			Annotations(entsql.IndexWhere("metadata_id IS NOT NULL AND node_id IS NULL")).
			StorageKey("idx_person_metadata_id"),
		index.Fields("node_id", "name", "is_org", "email", "url", "phone").
			Unique().
			Annotations(entsql.IndexWhere("metadata_id IS NULL AND node_id IS NOT NULL")).
			StorageKey("idx_person_node_id"),
	}
}
