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
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/protobom/protobom/pkg/sbom"

	entint "github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/hook"
)

type HashesEntry struct {
	ent.Schema
}

var errInvalidHashesEntry = errors.New("at least one of external_reference_id or node_id must be set")

func (HashesEntry) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DocumentMixin{},
		UUIDMixin{},
	}
}

func (HashesEntry) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("hash_algorithm").Values(enumValues(new(sbom.HashAlgorithm))...),
		field.String("hash_data"),
	}
}

func (HashesEntry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("external_references", ExternalReference.Type).Ref("hashes"),
		edge.From("nodes", Node.Type).Ref("hashes"),
	}
}

func (HashesEntry) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(hashesEntryHook, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
	}
}

func (HashesEntry) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hash_algorithm", "hash_data").
			Unique().
			StorageKey("idx_hashes"),
	}
}

func hashesEntryHook(next ent.Mutator) ent.Mutator {
	return hook.HashesEntryFunc(
		func(ctx context.Context, mutation *entint.HashesEntryMutation) (entint.Value, error) {
			extRefs := mutation.ExternalReferencesIDs()
			nodes := mutation.NodesIDs()

			// Fail validation if neither external_reference_id nor node_id are set.
			if len(extRefs) == 0 && len(nodes) == 0 {
				return nil, errInvalidHashesEntry
			}

			return next.Mutate(ctx, mutation)
		},
	)
}
