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

type HashesEntry struct {
	ent.Schema
}

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

func (HashesEntry) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hash_algorithm", "hash_data").
			Unique().
			StorageKey("idx_hashes"),
	}
}
