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
	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
)

type Node struct {
	ent.Schema
}

func (Node) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DocumentMixin{},
		ProtoMessageMixin[*sbom.Node]{},
	}
}

func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.String("native_id").NotEmpty().Immutable(),
		field.UUID("node_list_id", uuid.UUID{}).Optional(),
		field.Enum("type").Values("PACKAGE", "FILE"),
		field.String("name"),
		field.String("version"),
		field.String("file_name"),
		field.String("url_home"),
		field.String("url_download"),
		field.Strings("licenses"),
		field.String("license_concluded"),
		field.String("license_comments"),
		field.String("copyright"),
		field.String("source_info"),
		field.String("comment"),
		field.String("summary"),
		field.String("description"),
		field.Time("release_date"),
		field.Time("build_date"),
		field.Time("valid_until_date"),
		field.Strings("attribution"),
		field.Strings("file_types"),
	}
}

func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("annotations", Annotation.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("suppliers", Person.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("originators", Person.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("external_references", ExternalReference.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("primary_purpose", Purpose.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("nodes", Node.Type).
			From("to_nodes").
			Through("edge_types", EdgeType.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("hashes", HashesEntry.Type).
			StorageKey(edge.Table("node_hashes"), edge.Columns("node_id", "hash_entry_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("identifiers", IdentifiersEntry.Type).
			StorageKey(edge.Table("node_identifiers"), edge.Columns("node_id", "identifier_entry_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("properties", Property.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("node_lists", NodeList.Type).Ref("nodes"),
	}
}

func (Node) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("native_id", "node_list_id").
			Unique().
			StorageKey("idx_nodes"),
	}
}
