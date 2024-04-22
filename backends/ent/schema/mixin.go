// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type SourceDataMixin struct {
	mixin.Schema
}

func (sdm SourceDataMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("source_format").Values("cdx", "spdx").Immutable().Optional(),
		field.Any("source_data").Immutable().Optional(),
		field.Any("CDX_extra").Immutable().Optional(),
		field.Any("SPDX_extra").Immutable().Optional(),
	}
}
