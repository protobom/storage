// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package schema

import (
	"context"
	"errors"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"

	entint "github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/hook"
)

type Annotation struct {
	ent.Schema
}

var errInvalidAnnotation = errors.New("either document_id or node_id (exclusive) must be set")

func (Annotation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("document_id", uuid.UUID{}).
			Optional().
			Nillable(),
		field.UUID("node_id", uuid.UUID{}).
			Optional().
			Nillable(),
		field.String("name"),
		field.String("value"),
		field.Bool("is_unique").Default(false),
	}
}

func (Annotation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("document", Document.Type).
			Ref("annotations").
			Unique().
			Field("document_id"),
		edge.From("node", Node.Type).
			Ref("annotations").
			Unique().
			Field("node_id"),
	}
}

func (Annotation) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(annotationHook, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
	}
}

func (Annotation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("node_id").
			StorageKey("idx_annotations_node_id"),
		index.Fields("document_id").
			StorageKey("idx_annotations_document_id"),
		index.Fields("node_id", "name", "value").
			Unique().
			Annotations(entsql.IndexWhere("node_id IS NOT NULL AND TRIM(node_id) != ''")).
			StorageKey("idx_node_annotations"),
		index.Fields("node_id", "name").
			Unique().
			Annotations(entsql.IndexWhere("node_id IS NOT NULL AND TRIM(node_id) != '' AND is_unique")).
			StorageKey("idx_node_unique_annotations"),
		index.Fields("document_id", "name", "value").
			Unique().
			Annotations(entsql.IndexWhere("document_id IS NOT NULL AND TRIM(document_id) != ''")).
			StorageKey("idx_document_annotations"),
		index.Fields("document_id", "name").
			Unique().
			Annotations(entsql.IndexWhere("document_id IS NOT NULL AND TRIM(document_id) != '' AND is_unique")).
			StorageKey("idx_document_unique_annotations"),
	}
}

func annotationHook(next ent.Mutator) ent.Mutator {
	return hook.AnnotationFunc(
		func(ctx context.Context, mutation *entint.AnnotationMutation) (entint.Value, error) {
			_, docExists := mutation.DocumentID()
			_, nodeExists := mutation.NodeID()

			// Fail validation if both document_id and node_id are set, or neither are.
			if docExists == nodeExists {
				return nil, errInvalidAnnotation
			}

			return next.Mutate(ctx, mutation)
		},
	)
}
