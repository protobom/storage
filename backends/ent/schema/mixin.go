// ------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileName: backends/ent/schema/extra_data.go
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// ------------------------------------------------------------------------
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ------------------------------------------------------------------------
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	protobom "github.com/bom-squad/protobom/pkg/sbom"
)

type (
	protobomGenericType interface {
		protobom.Document |
			protobom.DocumentType |
			protobom.Edge |
			protobom.ExternalReference |
			protobom.Metadata |
			protobom.Node |
			protobom.NodeList |
			protobom.Person |
			protobom.Purpose |
			protobom.Tool |
			map[protobom.HashAlgorithm]string |
			map[protobom.SoftwareIdentifierType]string
	}

	SourceDataMixin[T protobomGenericType] struct {
		mixin.Schema
		protobomType *T
	}
)

func (sdm SourceDataMixin[protobomGenericType]) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("original_data", sdm.protobomType).Immutable().Optional(),
		field.Any("CDX_extra").Annotations().Immutable().Optional(),
		field.Any("SPDX_extra").Immutable().Optional(),
	}
}
