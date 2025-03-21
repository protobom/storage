// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package tool

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Tool {
	return predicate.Tool(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Tool {
	return predicate.Tool(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Tool {
	return predicate.Tool(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Tool {
	return predicate.Tool(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Tool {
	return predicate.Tool(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Tool {
	return predicate.Tool(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Tool {
	return predicate.Tool(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Tool {
	return predicate.Tool(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Tool {
	return predicate.Tool(sql.FieldLTE(FieldID, id))
}

// ProtoMessage applies equality check predicate on the "proto_message" field. It's identical to ProtoMessageEQ.
func ProtoMessage(v *sbom.Tool) predicate.Tool {
	return predicate.Tool(sql.FieldEQ(FieldProtoMessage, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Tool {
	return predicate.Tool(sql.FieldEQ(FieldName, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.Tool {
	return predicate.Tool(sql.FieldEQ(FieldVersion, v))
}

// Vendor applies equality check predicate on the "vendor" field. It's identical to VendorEQ.
func Vendor(v string) predicate.Tool {
	return predicate.Tool(sql.FieldEQ(FieldVendor, v))
}

// ProtoMessageEQ applies the EQ predicate on the "proto_message" field.
func ProtoMessageEQ(v *sbom.Tool) predicate.Tool {
	return predicate.Tool(sql.FieldEQ(FieldProtoMessage, v))
}

// ProtoMessageNEQ applies the NEQ predicate on the "proto_message" field.
func ProtoMessageNEQ(v *sbom.Tool) predicate.Tool {
	return predicate.Tool(sql.FieldNEQ(FieldProtoMessage, v))
}

// ProtoMessageIn applies the In predicate on the "proto_message" field.
func ProtoMessageIn(vs ...*sbom.Tool) predicate.Tool {
	return predicate.Tool(sql.FieldIn(FieldProtoMessage, vs...))
}

// ProtoMessageNotIn applies the NotIn predicate on the "proto_message" field.
func ProtoMessageNotIn(vs ...*sbom.Tool) predicate.Tool {
	return predicate.Tool(sql.FieldNotIn(FieldProtoMessage, vs...))
}

// ProtoMessageGT applies the GT predicate on the "proto_message" field.
func ProtoMessageGT(v *sbom.Tool) predicate.Tool {
	return predicate.Tool(sql.FieldGT(FieldProtoMessage, v))
}

// ProtoMessageGTE applies the GTE predicate on the "proto_message" field.
func ProtoMessageGTE(v *sbom.Tool) predicate.Tool {
	return predicate.Tool(sql.FieldGTE(FieldProtoMessage, v))
}

// ProtoMessageLT applies the LT predicate on the "proto_message" field.
func ProtoMessageLT(v *sbom.Tool) predicate.Tool {
	return predicate.Tool(sql.FieldLT(FieldProtoMessage, v))
}

// ProtoMessageLTE applies the LTE predicate on the "proto_message" field.
func ProtoMessageLTE(v *sbom.Tool) predicate.Tool {
	return predicate.Tool(sql.FieldLTE(FieldProtoMessage, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Tool {
	return predicate.Tool(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Tool {
	return predicate.Tool(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Tool {
	return predicate.Tool(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Tool {
	return predicate.Tool(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Tool {
	return predicate.Tool(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Tool {
	return predicate.Tool(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Tool {
	return predicate.Tool(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Tool {
	return predicate.Tool(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Tool {
	return predicate.Tool(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Tool {
	return predicate.Tool(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Tool {
	return predicate.Tool(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Tool {
	return predicate.Tool(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Tool {
	return predicate.Tool(sql.FieldContainsFold(FieldName, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.Tool {
	return predicate.Tool(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.Tool {
	return predicate.Tool(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.Tool {
	return predicate.Tool(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.Tool {
	return predicate.Tool(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.Tool {
	return predicate.Tool(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.Tool {
	return predicate.Tool(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.Tool {
	return predicate.Tool(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.Tool {
	return predicate.Tool(sql.FieldLTE(FieldVersion, v))
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.Tool {
	return predicate.Tool(sql.FieldContains(FieldVersion, v))
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.Tool {
	return predicate.Tool(sql.FieldHasPrefix(FieldVersion, v))
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.Tool {
	return predicate.Tool(sql.FieldHasSuffix(FieldVersion, v))
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.Tool {
	return predicate.Tool(sql.FieldEqualFold(FieldVersion, v))
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.Tool {
	return predicate.Tool(sql.FieldContainsFold(FieldVersion, v))
}

// VendorEQ applies the EQ predicate on the "vendor" field.
func VendorEQ(v string) predicate.Tool {
	return predicate.Tool(sql.FieldEQ(FieldVendor, v))
}

// VendorNEQ applies the NEQ predicate on the "vendor" field.
func VendorNEQ(v string) predicate.Tool {
	return predicate.Tool(sql.FieldNEQ(FieldVendor, v))
}

// VendorIn applies the In predicate on the "vendor" field.
func VendorIn(vs ...string) predicate.Tool {
	return predicate.Tool(sql.FieldIn(FieldVendor, vs...))
}

// VendorNotIn applies the NotIn predicate on the "vendor" field.
func VendorNotIn(vs ...string) predicate.Tool {
	return predicate.Tool(sql.FieldNotIn(FieldVendor, vs...))
}

// VendorGT applies the GT predicate on the "vendor" field.
func VendorGT(v string) predicate.Tool {
	return predicate.Tool(sql.FieldGT(FieldVendor, v))
}

// VendorGTE applies the GTE predicate on the "vendor" field.
func VendorGTE(v string) predicate.Tool {
	return predicate.Tool(sql.FieldGTE(FieldVendor, v))
}

// VendorLT applies the LT predicate on the "vendor" field.
func VendorLT(v string) predicate.Tool {
	return predicate.Tool(sql.FieldLT(FieldVendor, v))
}

// VendorLTE applies the LTE predicate on the "vendor" field.
func VendorLTE(v string) predicate.Tool {
	return predicate.Tool(sql.FieldLTE(FieldVendor, v))
}

// VendorContains applies the Contains predicate on the "vendor" field.
func VendorContains(v string) predicate.Tool {
	return predicate.Tool(sql.FieldContains(FieldVendor, v))
}

// VendorHasPrefix applies the HasPrefix predicate on the "vendor" field.
func VendorHasPrefix(v string) predicate.Tool {
	return predicate.Tool(sql.FieldHasPrefix(FieldVendor, v))
}

// VendorHasSuffix applies the HasSuffix predicate on the "vendor" field.
func VendorHasSuffix(v string) predicate.Tool {
	return predicate.Tool(sql.FieldHasSuffix(FieldVendor, v))
}

// VendorEqualFold applies the EqualFold predicate on the "vendor" field.
func VendorEqualFold(v string) predicate.Tool {
	return predicate.Tool(sql.FieldEqualFold(FieldVendor, v))
}

// VendorContainsFold applies the ContainsFold predicate on the "vendor" field.
func VendorContainsFold(v string) predicate.Tool {
	return predicate.Tool(sql.FieldContainsFold(FieldVendor, v))
}

// HasDocuments applies the HasEdge predicate on the "documents" edge.
func HasDocuments() predicate.Tool {
	return predicate.Tool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, DocumentsTable, DocumentsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDocumentsWith applies the HasEdge predicate on the "documents" edge with a given conditions (other predicates).
func HasDocumentsWith(preds ...predicate.Document) predicate.Tool {
	return predicate.Tool(func(s *sql.Selector) {
		step := newDocumentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMetadata applies the HasEdge predicate on the "metadata" edge.
func HasMetadata() predicate.Tool {
	return predicate.Tool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MetadataTable, MetadataPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetadataWith applies the HasEdge predicate on the "metadata" edge with a given conditions (other predicates).
func HasMetadataWith(preds ...predicate.Metadata) predicate.Tool {
	return predicate.Tool(func(s *sql.Selector) {
		step := newMetadataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tool) predicate.Tool {
	return predicate.Tool(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tool) predicate.Tool {
	return predicate.Tool(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tool) predicate.Tool {
	return predicate.Tool(sql.NotPredicates(p))
}
