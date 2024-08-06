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

func (HashesEntry) Fields() []ent.Field {
	values := []string{}
	for idx := range len(sbom.HashAlgorithm_name) {
		values = append(values, sbom.HashAlgorithm_name[int32(idx)])
	}

	return []ent.Field{
		field.Int("external_reference_id").Optional(),
		field.String("node_id").Optional(),
		field.Enum("hash_algorithm_type").Values(values...),
		field.String("hash_data"),
	}
}

func (HashesEntry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("external_reference", ExternalReference.Type).
			Ref("hashes").
			Unique().
			Field("external_reference_id"),
		edge.From("node", Node.Type).
			Ref("hashes").
			Unique().
			Field("node_id"),
	}
}

func (HashesEntry) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("external_reference_id", "node_id", "hash_algorithm_type", "hash_data").
			Unique().
			StorageKey("idx_hashes_entries"),
	}
}
