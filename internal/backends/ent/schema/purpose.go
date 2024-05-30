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

type Purpose struct {
	ent.Schema
}

func (Purpose) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("primary_purpose").Values(
			"UNKNOWN_PURPOSE",
			"APPLICATION",
			"ARCHIVE",
			"BOM",
			"CONFIGURATION",
			"CONTAINER",
			"DATA",
			"DEVICE",
			"DEVICE_DRIVER",
			"DOCUMENTATION",
			"EVIDENCE",
			"EXECUTABLE",
			"FILE",
			"FIRMWARE",
			"FRAMEWORK",
			"INSTALL",
			"LIBRARY",
			"MACHINE_LEARNING_MODEL",
			"MANIFEST",
			"MODEL",
			"MODULE",
			"OPERATING_SYSTEM",
			"OTHER",
			"PATCH",
			"PLATFORM",
			"REQUIREMENT",
			"SOURCE",
			"SPECIFICATION",
			"TEST",
		),
	}
}

func (Purpose) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("node", Node.Type).Ref("primary_purpose").Unique(),
	}
}

func (Purpose) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("node").Fields("id").Unique(),
	}
}
