// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package purpose

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Purpose {
	return predicate.Purpose(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Purpose {
	return predicate.Purpose(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Purpose {
	return predicate.Purpose(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Purpose {
	return predicate.Purpose(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Purpose {
	return predicate.Purpose(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Purpose {
	return predicate.Purpose(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Purpose {
	return predicate.Purpose(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Purpose {
	return predicate.Purpose(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Purpose {
	return predicate.Purpose(sql.FieldLTE(FieldID, id))
}

// DocumentID applies equality check predicate on the "document_id" field. It's identical to DocumentIDEQ.
func DocumentID(v uuid.UUID) predicate.Purpose {
	return predicate.Purpose(sql.FieldEQ(FieldDocumentID, v))
}

// NodeID applies equality check predicate on the "node_id" field. It's identical to NodeIDEQ.
func NodeID(v uuid.UUID) predicate.Purpose {
	return predicate.Purpose(sql.FieldEQ(FieldNodeID, v))
}

// DocumentIDEQ applies the EQ predicate on the "document_id" field.
func DocumentIDEQ(v uuid.UUID) predicate.Purpose {
	return predicate.Purpose(sql.FieldEQ(FieldDocumentID, v))
}

// DocumentIDNEQ applies the NEQ predicate on the "document_id" field.
func DocumentIDNEQ(v uuid.UUID) predicate.Purpose {
	return predicate.Purpose(sql.FieldNEQ(FieldDocumentID, v))
}

// DocumentIDIn applies the In predicate on the "document_id" field.
func DocumentIDIn(vs ...uuid.UUID) predicate.Purpose {
	return predicate.Purpose(sql.FieldIn(FieldDocumentID, vs...))
}

// DocumentIDNotIn applies the NotIn predicate on the "document_id" field.
func DocumentIDNotIn(vs ...uuid.UUID) predicate.Purpose {
	return predicate.Purpose(sql.FieldNotIn(FieldDocumentID, vs...))
}

// DocumentIDIsNil applies the IsNil predicate on the "document_id" field.
func DocumentIDIsNil() predicate.Purpose {
	return predicate.Purpose(sql.FieldIsNull(FieldDocumentID))
}

// DocumentIDNotNil applies the NotNil predicate on the "document_id" field.
func DocumentIDNotNil() predicate.Purpose {
	return predicate.Purpose(sql.FieldNotNull(FieldDocumentID))
}

// NodeIDEQ applies the EQ predicate on the "node_id" field.
func NodeIDEQ(v uuid.UUID) predicate.Purpose {
	return predicate.Purpose(sql.FieldEQ(FieldNodeID, v))
}

// NodeIDNEQ applies the NEQ predicate on the "node_id" field.
func NodeIDNEQ(v uuid.UUID) predicate.Purpose {
	return predicate.Purpose(sql.FieldNEQ(FieldNodeID, v))
}

// NodeIDIn applies the In predicate on the "node_id" field.
func NodeIDIn(vs ...uuid.UUID) predicate.Purpose {
	return predicate.Purpose(sql.FieldIn(FieldNodeID, vs...))
}

// NodeIDNotIn applies the NotIn predicate on the "node_id" field.
func NodeIDNotIn(vs ...uuid.UUID) predicate.Purpose {
	return predicate.Purpose(sql.FieldNotIn(FieldNodeID, vs...))
}

// NodeIDIsNil applies the IsNil predicate on the "node_id" field.
func NodeIDIsNil() predicate.Purpose {
	return predicate.Purpose(sql.FieldIsNull(FieldNodeID))
}

// NodeIDNotNil applies the NotNil predicate on the "node_id" field.
func NodeIDNotNil() predicate.Purpose {
	return predicate.Purpose(sql.FieldNotNull(FieldNodeID))
}

// PrimaryPurposeEQ applies the EQ predicate on the "primary_purpose" field.
func PrimaryPurposeEQ(v PrimaryPurpose) predicate.Purpose {
	return predicate.Purpose(sql.FieldEQ(FieldPrimaryPurpose, v))
}

// PrimaryPurposeNEQ applies the NEQ predicate on the "primary_purpose" field.
func PrimaryPurposeNEQ(v PrimaryPurpose) predicate.Purpose {
	return predicate.Purpose(sql.FieldNEQ(FieldPrimaryPurpose, v))
}

// PrimaryPurposeIn applies the In predicate on the "primary_purpose" field.
func PrimaryPurposeIn(vs ...PrimaryPurpose) predicate.Purpose {
	return predicate.Purpose(sql.FieldIn(FieldPrimaryPurpose, vs...))
}

// PrimaryPurposeNotIn applies the NotIn predicate on the "primary_purpose" field.
func PrimaryPurposeNotIn(vs ...PrimaryPurpose) predicate.Purpose {
	return predicate.Purpose(sql.FieldNotIn(FieldPrimaryPurpose, vs...))
}

// HasDocument applies the HasEdge predicate on the "document" edge.
func HasDocument() predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DocumentTable, DocumentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDocumentWith applies the HasEdge predicate on the "document" edge with a given conditions (other predicates).
func HasDocumentWith(preds ...predicate.Document) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		step := newDocumentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNode applies the HasEdge predicate on the "node" edge.
func HasNode() predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NodeTable, NodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodeWith applies the HasEdge predicate on the "node" edge with a given conditions (other predicates).
func HasNodeWith(preds ...predicate.Node) predicate.Purpose {
	return predicate.Purpose(func(s *sql.Selector) {
		step := newNodeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Purpose) predicate.Purpose {
	return predicate.Purpose(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Purpose) predicate.Purpose {
	return predicate.Purpose(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Purpose) predicate.Purpose {
	return predicate.Purpose(sql.NotPredicates(p))
}
