// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package mixin

import (
	"context"
	"crypto/sha256"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	"github.com/protobom/storage/internal/backends/ent/hook"
)

// UUID replaces the default integer `id` field with a generated UUID.
type UUID struct {
	mixin.Schema
}

func (UUID) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Immutable().
			StructTag(`json:"-"`).
			Annotations(schema.Comment("Unique identifier field")).
			Default(func() uuid.UUID { return uuid.Must(uuid.NewV7()) }),
	}
}

func (UUID) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(uuidHook, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
	}
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
		if !fieldSet || !ok {
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
