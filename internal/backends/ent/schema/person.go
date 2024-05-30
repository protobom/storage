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
)

type Person struct {
	ent.Schema
}

func (Person) Fields() []ent.Field {
	return []ent.Field{
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
		edge.From("metadata", Metadata.Type).Ref("authors").Unique(),
		edge.From("node", Node.Type).Ref("suppliers").Unique().Ref("originators").Unique(),
	}
}

func (Person) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "is_org", "email", "url", "phone").Unique(),
		index.Edges("metadata").Fields("id").Unique(),
		index.Edges("node").Fields("id").Unique(),
	}
}
