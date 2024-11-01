// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package sourcedata

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldLTE(FieldID, id))
}

// DocumentID applies equality check predicate on the "document_id" field. It's identical to DocumentIDEQ.
func DocumentID(v uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldDocumentID, v))
}

// ProtoMessage applies equality check predicate on the "proto_message" field. It's identical to ProtoMessageEQ.
func ProtoMessage(v *sbom.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldProtoMessage, v))
}

// MetadataID applies equality check predicate on the "metadata_id" field. It's identical to MetadataIDEQ.
func MetadataID(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldMetadataID, v))
}

// Format applies equality check predicate on the "format" field. It's identical to FormatEQ.
func Format(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldFormat, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldSize, v))
}

// URI applies equality check predicate on the "uri" field. It's identical to URIEQ.
func URI(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldURI, v))
}

// DocumentIDEQ applies the EQ predicate on the "document_id" field.
func DocumentIDEQ(v uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldDocumentID, v))
}

// DocumentIDNEQ applies the NEQ predicate on the "document_id" field.
func DocumentIDNEQ(v uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldDocumentID, v))
}

// DocumentIDIn applies the In predicate on the "document_id" field.
func DocumentIDIn(vs ...uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldDocumentID, vs...))
}

// DocumentIDNotIn applies the NotIn predicate on the "document_id" field.
func DocumentIDNotIn(vs ...uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldDocumentID, vs...))
}

// DocumentIDIsNil applies the IsNil predicate on the "document_id" field.
func DocumentIDIsNil() predicate.SourceData {
	return predicate.SourceData(sql.FieldIsNull(FieldDocumentID))
}

// DocumentIDNotNil applies the NotNil predicate on the "document_id" field.
func DocumentIDNotNil() predicate.SourceData {
	return predicate.SourceData(sql.FieldNotNull(FieldDocumentID))
}

// ProtoMessageEQ applies the EQ predicate on the "proto_message" field.
func ProtoMessageEQ(v *sbom.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldProtoMessage, v))
}

// ProtoMessageNEQ applies the NEQ predicate on the "proto_message" field.
func ProtoMessageNEQ(v *sbom.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldProtoMessage, v))
}

// ProtoMessageIn applies the In predicate on the "proto_message" field.
func ProtoMessageIn(vs ...*sbom.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldProtoMessage, vs...))
}

// ProtoMessageNotIn applies the NotIn predicate on the "proto_message" field.
func ProtoMessageNotIn(vs ...*sbom.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldProtoMessage, vs...))
}

// ProtoMessageGT applies the GT predicate on the "proto_message" field.
func ProtoMessageGT(v *sbom.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.FieldGT(FieldProtoMessage, v))
}

// ProtoMessageGTE applies the GTE predicate on the "proto_message" field.
func ProtoMessageGTE(v *sbom.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.FieldGTE(FieldProtoMessage, v))
}

// ProtoMessageLT applies the LT predicate on the "proto_message" field.
func ProtoMessageLT(v *sbom.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.FieldLT(FieldProtoMessage, v))
}

// ProtoMessageLTE applies the LTE predicate on the "proto_message" field.
func ProtoMessageLTE(v *sbom.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.FieldLTE(FieldProtoMessage, v))
}

// MetadataIDEQ applies the EQ predicate on the "metadata_id" field.
func MetadataIDEQ(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldMetadataID, v))
}

// MetadataIDNEQ applies the NEQ predicate on the "metadata_id" field.
func MetadataIDNEQ(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldMetadataID, v))
}

// MetadataIDIn applies the In predicate on the "metadata_id" field.
func MetadataIDIn(vs ...string) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldMetadataID, vs...))
}

// MetadataIDNotIn applies the NotIn predicate on the "metadata_id" field.
func MetadataIDNotIn(vs ...string) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldMetadataID, vs...))
}

// MetadataIDGT applies the GT predicate on the "metadata_id" field.
func MetadataIDGT(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldGT(FieldMetadataID, v))
}

// MetadataIDGTE applies the GTE predicate on the "metadata_id" field.
func MetadataIDGTE(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldGTE(FieldMetadataID, v))
}

// MetadataIDLT applies the LT predicate on the "metadata_id" field.
func MetadataIDLT(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldLT(FieldMetadataID, v))
}

// MetadataIDLTE applies the LTE predicate on the "metadata_id" field.
func MetadataIDLTE(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldLTE(FieldMetadataID, v))
}

// MetadataIDContains applies the Contains predicate on the "metadata_id" field.
func MetadataIDContains(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldContains(FieldMetadataID, v))
}

// MetadataIDHasPrefix applies the HasPrefix predicate on the "metadata_id" field.
func MetadataIDHasPrefix(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldHasPrefix(FieldMetadataID, v))
}

// MetadataIDHasSuffix applies the HasSuffix predicate on the "metadata_id" field.
func MetadataIDHasSuffix(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldHasSuffix(FieldMetadataID, v))
}

// MetadataIDEqualFold applies the EqualFold predicate on the "metadata_id" field.
func MetadataIDEqualFold(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEqualFold(FieldMetadataID, v))
}

// MetadataIDContainsFold applies the ContainsFold predicate on the "metadata_id" field.
func MetadataIDContainsFold(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldContainsFold(FieldMetadataID, v))
}

// FormatEQ applies the EQ predicate on the "format" field.
func FormatEQ(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldFormat, v))
}

// FormatNEQ applies the NEQ predicate on the "format" field.
func FormatNEQ(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldFormat, v))
}

// FormatIn applies the In predicate on the "format" field.
func FormatIn(vs ...string) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldFormat, vs...))
}

// FormatNotIn applies the NotIn predicate on the "format" field.
func FormatNotIn(vs ...string) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldFormat, vs...))
}

// FormatGT applies the GT predicate on the "format" field.
func FormatGT(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldGT(FieldFormat, v))
}

// FormatGTE applies the GTE predicate on the "format" field.
func FormatGTE(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldGTE(FieldFormat, v))
}

// FormatLT applies the LT predicate on the "format" field.
func FormatLT(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldLT(FieldFormat, v))
}

// FormatLTE applies the LTE predicate on the "format" field.
func FormatLTE(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldLTE(FieldFormat, v))
}

// FormatContains applies the Contains predicate on the "format" field.
func FormatContains(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldContains(FieldFormat, v))
}

// FormatHasPrefix applies the HasPrefix predicate on the "format" field.
func FormatHasPrefix(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldHasPrefix(FieldFormat, v))
}

// FormatHasSuffix applies the HasSuffix predicate on the "format" field.
func FormatHasSuffix(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldHasSuffix(FieldFormat, v))
}

// FormatEqualFold applies the EqualFold predicate on the "format" field.
func FormatEqualFold(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEqualFold(FieldFormat, v))
}

// FormatContainsFold applies the ContainsFold predicate on the "format" field.
func FormatContainsFold(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldContainsFold(FieldFormat, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldLTE(FieldSize, v))
}

// URIEQ applies the EQ predicate on the "uri" field.
func URIEQ(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldURI, v))
}

// URINEQ applies the NEQ predicate on the "uri" field.
func URINEQ(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldURI, v))
}

// URIIn applies the In predicate on the "uri" field.
func URIIn(vs ...string) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldURI, vs...))
}

// URINotIn applies the NotIn predicate on the "uri" field.
func URINotIn(vs ...string) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldURI, vs...))
}

// URIGT applies the GT predicate on the "uri" field.
func URIGT(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldGT(FieldURI, v))
}

// URIGTE applies the GTE predicate on the "uri" field.
func URIGTE(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldGTE(FieldURI, v))
}

// URILT applies the LT predicate on the "uri" field.
func URILT(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldLT(FieldURI, v))
}

// URILTE applies the LTE predicate on the "uri" field.
func URILTE(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldLTE(FieldURI, v))
}

// URIContains applies the Contains predicate on the "uri" field.
func URIContains(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldContains(FieldURI, v))
}

// URIHasPrefix applies the HasPrefix predicate on the "uri" field.
func URIHasPrefix(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldHasPrefix(FieldURI, v))
}

// URIHasSuffix applies the HasSuffix predicate on the "uri" field.
func URIHasSuffix(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldHasSuffix(FieldURI, v))
}

// URIIsNil applies the IsNil predicate on the "uri" field.
func URIIsNil() predicate.SourceData {
	return predicate.SourceData(sql.FieldIsNull(FieldURI))
}

// URINotNil applies the NotNil predicate on the "uri" field.
func URINotNil() predicate.SourceData {
	return predicate.SourceData(sql.FieldNotNull(FieldURI))
}

// URIEqualFold applies the EqualFold predicate on the "uri" field.
func URIEqualFold(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEqualFold(FieldURI, v))
}

// URIContainsFold applies the ContainsFold predicate on the "uri" field.
func URIContainsFold(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldContainsFold(FieldURI, v))
}

// HashesIsNil applies the IsNil predicate on the "hashes" field.
func HashesIsNil() predicate.SourceData {
	return predicate.SourceData(sql.FieldIsNull(FieldHashes))
}

// HashesNotNil applies the NotNil predicate on the "hashes" field.
func HashesNotNil() predicate.SourceData {
	return predicate.SourceData(sql.FieldNotNull(FieldHashes))
}

// HasDocument applies the HasEdge predicate on the "document" edge.
func HasDocument() predicate.SourceData {
	return predicate.SourceData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DocumentTable, DocumentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDocumentWith applies the HasEdge predicate on the "document" edge with a given conditions (other predicates).
func HasDocumentWith(preds ...predicate.Document) predicate.SourceData {
	return predicate.SourceData(func(s *sql.Selector) {
		step := newDocumentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMetadata applies the HasEdge predicate on the "metadata" edge.
func HasMetadata() predicate.SourceData {
	return predicate.SourceData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetadataWith applies the HasEdge predicate on the "metadata" edge with a given conditions (other predicates).
func HasMetadataWith(preds ...predicate.Metadata) predicate.SourceData {
	return predicate.SourceData(func(s *sql.Selector) {
		step := newMetadataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.NotPredicates(p))
}
