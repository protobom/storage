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

type HashesEntry struct {
	ent.Schema
}

func (HashesEntry) Fields() []ent.Field {
	return []ent.Field{
		field.Int("external_reference_id").Optional(),
		field.String("node_id").Optional(),
		field.Enum("hash_algorithm_type").Values(
			"UNKNOWN",
			"MD5",
			"SHA1",
			"SHA256",
			"SHA384",
			"SHA512",
			"SHA3_256",
			"SHA3_384",
			"SHA3_512",
			"BLAKE2B_256",
			"BLAKE2B_384",
			"BLAKE2B_512",
			"BLAKE3",
			"MD2",
			"ADLER32",
			"MD4",
			"MD6",
			"SHA224",
		),
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
