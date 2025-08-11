// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"google.golang.org/protobuf/proto"

	"github.com/protobom/storage/internal/backends/ent/hook"
)

const protoMessageField = "proto_message"

// ProtoMessage adds the `proto_message` field containing the wire format bytes.
type ProtoMessage[T proto.Message] struct {
	UUID
}

func (pm ProtoMessage[T]) Fields() []ent.Field {
	var goType T

	return append(
		pm.UUID.Fields(),
		field.Bytes(protoMessageField).
			GoType(goType).
			Nillable().
			Immutable().
			StructTag(`json:"-"`),
	)
}

func (ProtoMessage[T]) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(uuidHook, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
	}
}
