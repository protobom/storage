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
	"github.com/protobom/protobom/pkg/sbom"

	entint "github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/hook"
	"github.com/protobom/storage/internal/backends/ent/schema/mixin"
)

type Person struct {
	ent.Schema
}

var errInvalidPerson = errors.New("either metadata_id or node_id (exclusive) must be set")

func (Person) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ProtoMessage[*sbom.Person]{},
	}
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
		edge.To("contacts", Person.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			From("contact_owner"),
		edge.From("documents", Document.Type).
			Ref("persons").
			Required().
			Immutable(),
		edge.From("metadata", Metadata.Type).
			Ref("authors"),
		edge.From("originator_nodes", Node.Type).
			Ref("originators"),
		edge.From("supplier_nodes", Node.Type).
			Ref("suppliers"),
	}
}

func (Person) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(personHook, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
	}
}

func (Person) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "is_org", "email", "url", "phone").
			Unique().
			StorageKey("idx_persons"),
	}
}

func personHook(next ent.Mutator) ent.Mutator {
	return hook.PersonFunc(
		func(ctx context.Context, mutation *entint.PersonMutation) (entint.Value, error) {
			missing := (len(mutation.OriginatorNodesIDs()) +
				len(mutation.SupplierNodesIDs()) +
				len(mutation.MetadataIDs()) +
				len(mutation.ContactOwnerIDs())) == 0

			// Fail validation if no inverse edge reference is set.
			if missing {
				return nil, errInvalidPerson
			}

			return next.Mutate(ctx, mutation)
		},
	)
}
