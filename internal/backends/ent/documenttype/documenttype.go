// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package documenttype

import (
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the documenttype type in the database.
	Label = "document_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProtoMessage holds the string denoting the proto_message field in the database.
	FieldProtoMessage = "proto_message"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeDocuments holds the string denoting the documents edge name in mutations.
	EdgeDocuments = "documents"
	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// Table holds the table name of the documenttype in the database.
	Table = "document_types"
	// DocumentsTable is the table that holds the documents relation/edge. The primary key declared below.
	DocumentsTable = "document_document_types"
	// DocumentsInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	DocumentsInverseTable = "documents"
	// MetadataTable is the table that holds the metadata relation/edge. The primary key declared below.
	MetadataTable = "metadata_document_types"
	// MetadataInverseTable is the table name for the Metadata entity.
	// It exists in this package in order to avoid circular dependency with the "metadata" package.
	MetadataInverseTable = "metadata"
)

// Columns holds all SQL columns for documenttype fields.
var Columns = []string{
	FieldID,
	FieldProtoMessage,
	FieldType,
	FieldName,
	FieldDescription,
}

var (
	// DocumentsPrimaryKey and DocumentsColumn2 are the table columns denoting the
	// primary key for the documents relation (M2M).
	DocumentsPrimaryKey = []string{"document_id", "document_type_id"}
	// MetadataPrimaryKey and MetadataColumn2 are the table columns denoting the
	// primary key for the metadata relation (M2M).
	MetadataPrimaryKey = []string{"metadata_id", "document_type_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/protobom/storage/internal/backends/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeOTHER       Type = "OTHER"
	TypeDESIGN      Type = "DESIGN"
	TypeSOURCE      Type = "SOURCE"
	TypeBUILD       Type = "BUILD"
	TypeANALYZED    Type = "ANALYZED"
	TypeDEPLOYED    Type = "DEPLOYED"
	TypeRUNTIME     Type = "RUNTIME"
	TypeDISCOVERY   Type = "DISCOVERY"
	TypeDECOMISSION Type = "DECOMISSION"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeOTHER, TypeDESIGN, TypeSOURCE, TypeBUILD, TypeANALYZED, TypeDEPLOYED, TypeRUNTIME, TypeDISCOVERY, TypeDECOMISSION:
		return nil
	default:
		return fmt.Errorf("documenttype: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the DocumentType queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByDocumentsCount orders the results by documents count.
func ByDocumentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDocumentsStep(), opts...)
	}
}

// ByDocuments orders the results by documents terms.
func ByDocuments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDocumentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMetadataCount orders the results by metadata count.
func ByMetadataCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMetadataStep(), opts...)
	}
}

// ByMetadata orders the results by metadata terms.
func ByMetadata(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMetadataStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDocumentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DocumentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, DocumentsTable, DocumentsPrimaryKey...),
	)
}
func newMetadataStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MetadataInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, MetadataTable, MetadataPrimaryKey...),
	)
}
