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

type Node struct {
	ent.Schema
}

func (Node) Fields() []ent.Field { //nolint:funlen
	return []ent.Field{
		field.String("id").NotEmpty().Unique().Immutable(),
		field.String("from_node_id").Optional(),
		field.Int("node_list_id").Optional(),
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
		field.Enum("edge_type").Values(
			"UNKNOWN",
			"amends",
			"ancestor",
			"buildDependency",
			"buildTool",
			"contains",
			"contained_by",
			"copy",
			"dataFile",
			"dependencyManifest",
			"dependsOn",
			"dependencyOf",
			"descendant",
			"describes",
			"describedBy",
			"devDependency",
			"devTool",
			"distributionArtifact",
			"documentation",
			"dynamicLink",
			"example",
			"expandedFromArchive",
			"fileAdded",
			"fileDeleted",
			"fileModified",
			"generates",
			"generatedFrom",
			"metafile",
			"optionalComponent",
			"optionalDependency",
			"other",
			"packages",
			"patch",
			"prerequisite",
			"prerequisiteFor",
			"providedDependency",
			"requirementFor",
			"runtimeDependency",
			"specificationFor",
			"staticLink",
			"test",
			"testCase",
			"testDependency",
			"testTool",
			"variant",
		).Optional(),
	}
}

func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("suppliers", Person.Type),
		edge.To("originators", Person.Type),
		edge.To("external_references", ExternalReference.Type),
		edge.To("identifiers", IdentifiersEntry.Type),
		edge.To("hashes", HashesEntry.Type),
		edge.To("primary_purpose", Purpose.Type),
		edge.To("nodes", Node.Type).From("from_node").Unique().Field("from_node_id"),
		edge.From("node_list", NodeList.Type).Ref("nodes").Unique().Field("node_list_id"),
	}
}

func (Node) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "node_list_id").Unique(),
	}
}
