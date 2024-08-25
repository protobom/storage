// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
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

	UUIDMixin struct {
		mixin.Schema
	}
)

func (DocumentMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("document_id", uuid.UUID{}).
			Optional().
			Immutable().
			Default(func() uuid.UUID { return uuid.Must(uuid.NewV7()) }),
	}
}

func (DocumentMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("document", Document.Type).
			Unique().
			Immutable().
			Field("document_id"),
	}
}

func (m ProtoMessageMixin) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("proto_message", m.ProtoMessageType).Optional(),
	}
}

func (UUIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Immutable().
			Annotations(schema.Comment("Unique identifier field")),
	}
}
