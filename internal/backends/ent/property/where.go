// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package property

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Property {
	return predicate.Property(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Property {
	return predicate.Property(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Property {
	return predicate.Property(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Property {
	return predicate.Property(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Property {
	return predicate.Property(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Property {
	return predicate.Property(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Property {
	return predicate.Property(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Property {
	return predicate.Property(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Property {
	return predicate.Property(sql.FieldLTE(FieldID, id))
}

// ProtoMessage applies equality check predicate on the "proto_message" field. It's identical to ProtoMessageEQ.
func ProtoMessage(v *sbom.Property) predicate.Property {
	return predicate.Property(sql.FieldEQ(FieldProtoMessage, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Property {
	return predicate.Property(sql.FieldEQ(FieldName, v))
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v string) predicate.Property {
	return predicate.Property(sql.FieldEQ(FieldData, v))
}

// ProtoMessageEQ applies the EQ predicate on the "proto_message" field.
func ProtoMessageEQ(v *sbom.Property) predicate.Property {
	return predicate.Property(sql.FieldEQ(FieldProtoMessage, v))
}

// ProtoMessageNEQ applies the NEQ predicate on the "proto_message" field.
func ProtoMessageNEQ(v *sbom.Property) predicate.Property {
	return predicate.Property(sql.FieldNEQ(FieldProtoMessage, v))
}

// ProtoMessageIn applies the In predicate on the "proto_message" field.
func ProtoMessageIn(vs ...*sbom.Property) predicate.Property {
	return predicate.Property(sql.FieldIn(FieldProtoMessage, vs...))
}

// ProtoMessageNotIn applies the NotIn predicate on the "proto_message" field.
func ProtoMessageNotIn(vs ...*sbom.Property) predicate.Property {
	return predicate.Property(sql.FieldNotIn(FieldProtoMessage, vs...))
}

// ProtoMessageGT applies the GT predicate on the "proto_message" field.
func ProtoMessageGT(v *sbom.Property) predicate.Property {
	return predicate.Property(sql.FieldGT(FieldProtoMessage, v))
}

// ProtoMessageGTE applies the GTE predicate on the "proto_message" field.
func ProtoMessageGTE(v *sbom.Property) predicate.Property {
	return predicate.Property(sql.FieldGTE(FieldProtoMessage, v))
}

// ProtoMessageLT applies the LT predicate on the "proto_message" field.
func ProtoMessageLT(v *sbom.Property) predicate.Property {
	return predicate.Property(sql.FieldLT(FieldProtoMessage, v))
}

// ProtoMessageLTE applies the LTE predicate on the "proto_message" field.
func ProtoMessageLTE(v *sbom.Property) predicate.Property {
	return predicate.Property(sql.FieldLTE(FieldProtoMessage, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Property {
	return predicate.Property(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Property {
	return predicate.Property(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Property {
	return predicate.Property(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Property {
	return predicate.Property(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Property {
	return predicate.Property(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Property {
	return predicate.Property(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Property {
	return predicate.Property(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Property {
	return predicate.Property(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Property {
	return predicate.Property(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Property {
	return predicate.Property(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Property {
	return predicate.Property(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Property {
	return predicate.Property(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Property {
	return predicate.Property(sql.FieldContainsFold(FieldName, v))
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v string) predicate.Property {
	return predicate.Property(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v string) predicate.Property {
	return predicate.Property(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...string) predicate.Property {
	return predicate.Property(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...string) predicate.Property {
	return predicate.Property(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v string) predicate.Property {
	return predicate.Property(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v string) predicate.Property {
	return predicate.Property(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v string) predicate.Property {
	return predicate.Property(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v string) predicate.Property {
	return predicate.Property(sql.FieldLTE(FieldData, v))
}

// DataContains applies the Contains predicate on the "data" field.
func DataContains(v string) predicate.Property {
	return predicate.Property(sql.FieldContains(FieldData, v))
}

// DataHasPrefix applies the HasPrefix predicate on the "data" field.
func DataHasPrefix(v string) predicate.Property {
	return predicate.Property(sql.FieldHasPrefix(FieldData, v))
}

// DataHasSuffix applies the HasSuffix predicate on the "data" field.
func DataHasSuffix(v string) predicate.Property {
	return predicate.Property(sql.FieldHasSuffix(FieldData, v))
}

// DataEqualFold applies the EqualFold predicate on the "data" field.
func DataEqualFold(v string) predicate.Property {
	return predicate.Property(sql.FieldEqualFold(FieldData, v))
}

// DataContainsFold applies the ContainsFold predicate on the "data" field.
func DataContainsFold(v string) predicate.Property {
	return predicate.Property(sql.FieldContainsFold(FieldData, v))
}

// HasDocuments applies the HasEdge predicate on the "documents" edge.
func HasDocuments() predicate.Property {
	return predicate.Property(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, DocumentsTable, DocumentsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDocumentsWith applies the HasEdge predicate on the "documents" edge with a given conditions (other predicates).
func HasDocumentsWith(preds ...predicate.Document) predicate.Property {
	return predicate.Property(func(s *sql.Selector) {
		step := newDocumentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNodes applies the HasEdge predicate on the "nodes" edge.
func HasNodes() predicate.Property {
	return predicate.Property(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, NodesTable, NodesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodesWith applies the HasEdge predicate on the "nodes" edge with a given conditions (other predicates).
func HasNodesWith(preds ...predicate.Node) predicate.Property {
	return predicate.Property(func(s *sql.Selector) {
		step := newNodesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Property) predicate.Property {
	return predicate.Property(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Property) predicate.Property {
	return predicate.Property(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Property) predicate.Property {
	return predicate.Property(sql.NotPredicates(p))
}
