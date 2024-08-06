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
	"entgo.io/ent/schema/mixin"
	"google.golang.org/protobuf/proto"
)

type (
	DocumentMixin struct {
		mixin.Schema
	}

	ProtoMessageMixin struct {
		mixin.Schema
		ProtoMessageType proto.Message
	}
)

func (m DocumentMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("document", Document.Type).
			Unique().
			Immutable().
			StorageKey(edge.Column("document_id")),
	}
}

func (m ProtoMessageMixin) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("proto_message", m.ProtoMessageType).Optional(),
	}
}
