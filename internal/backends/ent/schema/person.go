// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package schema

import (
	"context"
	"errors"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"

	entint "github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/hook"
)

type Person struct {
	ent.Schema
}

var errInvalidPerson = errors.New("either metadata_id or node_id (exclusive) must be set")

func (Person) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DocumentMixin{},
		ProtoMessageMixin[*sbom.Person]{},
	}
}

func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("metadata_id", uuid.UUID{}).Optional(),
		field.UUID("node_id", uuid.UUID{}).Optional(),
		field.String("name"),
		field.Bool("is_org"),
		field.String("email"),
		field.String("url"),
		field.String("phone"),
	}
}

func (Person) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contacts", Person.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			From("contact_owner").
			Unique(),
		edge.From("metadata", Metadata.Type).
			Ref("authors").
			Unique().
			Field("metadata_id"),
		edge.From("node", Node.Type).
			Ref("suppliers").
			Unique().
			Ref("originators").
			Unique().
			Field("node_id"),
	}
}

func (Person) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(personHook, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
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

func personHook(next ent.Mutator) ent.Mutator {
	return hook.PersonFunc(
		func(ctx context.Context, mutation *entint.PersonMutation) (entint.Value, error) {
			_, nodeExists := mutation.NodeID()
			_, metadataExists := mutation.MetadataID()

			// Fail validation if both metadata_id and node_id are set, or neither are.
			if metadataExists == nodeExists {
				return nil, errInvalidPerson
			}

			return next.Mutate(ctx, mutation)
		},
	)
}
