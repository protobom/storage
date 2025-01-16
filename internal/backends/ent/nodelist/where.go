// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package nodelist

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.NodeList {
	return predicate.NodeList(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.NodeList {
	return predicate.NodeList(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.NodeList {
	return predicate.NodeList(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.NodeList {
	return predicate.NodeList(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.NodeList {
	return predicate.NodeList(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.NodeList {
	return predicate.NodeList(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.NodeList {
	return predicate.NodeList(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.NodeList {
	return predicate.NodeList(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.NodeList {
	return predicate.NodeList(sql.FieldLTE(FieldID, id))
}

// ProtoMessage applies equality check predicate on the "proto_message" field. It's identical to ProtoMessageEQ.
func ProtoMessage(v *sbom.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.FieldEQ(FieldProtoMessage, v))
}

// ProtoMessageEQ applies the EQ predicate on the "proto_message" field.
func ProtoMessageEQ(v *sbom.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.FieldEQ(FieldProtoMessage, v))
}

// ProtoMessageNEQ applies the NEQ predicate on the "proto_message" field.
func ProtoMessageNEQ(v *sbom.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.FieldNEQ(FieldProtoMessage, v))
}

// ProtoMessageIn applies the In predicate on the "proto_message" field.
func ProtoMessageIn(vs ...*sbom.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.FieldIn(FieldProtoMessage, vs...))
}

// ProtoMessageNotIn applies the NotIn predicate on the "proto_message" field.
func ProtoMessageNotIn(vs ...*sbom.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.FieldNotIn(FieldProtoMessage, vs...))
}

// ProtoMessageGT applies the GT predicate on the "proto_message" field.
func ProtoMessageGT(v *sbom.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.FieldGT(FieldProtoMessage, v))
}

// ProtoMessageGTE applies the GTE predicate on the "proto_message" field.
func ProtoMessageGTE(v *sbom.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.FieldGTE(FieldProtoMessage, v))
}

// ProtoMessageLT applies the LT predicate on the "proto_message" field.
func ProtoMessageLT(v *sbom.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.FieldLT(FieldProtoMessage, v))
}

// ProtoMessageLTE applies the LTE predicate on the "proto_message" field.
func ProtoMessageLTE(v *sbom.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.FieldLTE(FieldProtoMessage, v))
}

// HasDocument applies the HasEdge predicate on the "document" edge.
func HasDocument() predicate.NodeList {
	return predicate.NodeList(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, DocumentTable, DocumentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDocumentWith applies the HasEdge predicate on the "document" edge with a given conditions (other predicates).
func HasDocumentWith(preds ...predicate.Document) predicate.NodeList {
	return predicate.NodeList(func(s *sql.Selector) {
		step := newDocumentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEdgeTypes applies the HasEdge predicate on the "edge_types" edge.
func HasEdgeTypes() predicate.NodeList {
	return predicate.NodeList(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, EdgeTypesTable, EdgeTypesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgeTypesWith applies the HasEdge predicate on the "edge_types" edge with a given conditions (other predicates).
func HasEdgeTypesWith(preds ...predicate.EdgeType) predicate.NodeList {
	return predicate.NodeList(func(s *sql.Selector) {
		step := newEdgeTypesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNodes applies the HasEdge predicate on the "nodes" edge.
func HasNodes() predicate.NodeList {
	return predicate.NodeList(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, NodesTable, NodesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodesWith applies the HasEdge predicate on the "nodes" edge with a given conditions (other predicates).
func HasNodesWith(preds ...predicate.Node) predicate.NodeList {
	return predicate.NodeList(func(s *sql.Selector) {
		step := newNodesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.NotPredicates(p))
}
