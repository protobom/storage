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

type EdgeType struct {
	ent.Schema
}

func (EdgeType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values(
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
		),
		field.String("node_id"),
		field.String("to_node_id"),
	}
}

func (EdgeType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("from", Node.Type).Required().Unique().Field("node_id"),
		edge.To("to", Node.Type).Required().Unique().Field("to_node_id"),
	}
}

func (EdgeType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "type", "node_id", "to_node_id").Unique(),
	}
}
