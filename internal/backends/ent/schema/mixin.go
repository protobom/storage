// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package schema

import (
	"context"
	"crypto/sha256"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/protobom/storage/internal/backends/ent/hook"
)

const protoMessageField = "proto_message"

type (
	// DocumentMixin adds the `document` edge and corresponding `document_id` edge field.
	DocumentMixin struct {
		mixin.Schema
	}

	// ProtoMessageMixin adds the `proto_message` field containing the wire format bytes.
	ProtoMessageMixin[T proto.Message] struct {
		UUIDMixin
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

func (pmm ProtoMessageMixin[T]) Fields() []ent.Field {
	var goType T

	return append(
		pmm.UUIDMixin.Fields(),
		field.Bytes(protoMessageField).
			GoType(goType).
			Nillable().
			Unique().
			Immutable(),
	)
}

func (ProtoMessageMixin[T]) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(uuidHook, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
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

func uuidHook(next ent.Mutator) ent.Mutator {
	type IDMutation interface {
		ent.Mutation
		SetID(uuid.UUID)
	}

	return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
		// Get the value of the `proto_message` field being set as part of this mutation.
		value, fieldSet := mutation.Field(protoMessageField)

		// Assert that the value of the field has type `proto.Message`.
		protoMessage, ok := value.(proto.Message)
		if !(fieldSet && ok) {
			return next.Mutate(ctx, mutation)
		}

		// Deterministically generate the wire-format encoding of the protobuf message.
		data, err := proto.MarshalOptions{Deterministic: true}.Marshal(protoMessage)
		if err != nil {
			return mutation, fmt.Errorf("marshaling proto: %w", err)
		}

		// Generate a UUID by hashing the wire-format bytes of the protobuf message.
		uuidHash := uuid.NewHash(sha256.New(), uuid.Max, data, int(uuid.Max.Version()))

		// Set the generated UUID as the value to be inserted as part of this mutation.
		if mut, ok := any(mutation).(IDMutation); ok {
			mut.SetID(uuidHash)
			mutation = mut
		}

		return next.Mutate(ctx, mutation)
	})
}
