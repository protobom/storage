// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type (
	// DocumentMixin adds the `document` edge and corresponding `document_id` edge field.
	DocumentMixin struct {
		mixin.Schema
	}

	// ProtoMessageMixin adds the `proto_message` field containing the wire format bytes.
	ProtoMessageMixin[T proto.Message] struct {
		mixin.Schema
	}

	// UUIDMixin replaces the default integer `id` field with a generated UUID.
	UUIDMixin struct {
		mixin.Schema
	}
)

func (DocumentMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.OnDelete(entsql.Cascade),
	}
}

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
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Field("document_id"),
	}
}

func (ProtoMessageMixin[T]) Fields() []ent.Field {
	var goType T

	return []ent.Field{
		field.Bytes("proto_message").GoType(goType).Nillable().Immutable(),
	}
}

func (UUIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Immutable().
			Default(func() uuid.UUID { return uuid.Must(uuid.NewV7()) }).
			Annotations(schema.Comment("Unique identifier field")),
	}
}

// enumValues returns the values of a protobuf enum type deterministically, preserving their order.
func enumValues(enum protoreflect.Enum) []string {
	values := []string{}

	enumValues := enum.Descriptor().Values()
	for idx := range enumValues.Len() {
		values = append(values, string(enumValues.Get(idx).Name()))
	}

	return values
}
