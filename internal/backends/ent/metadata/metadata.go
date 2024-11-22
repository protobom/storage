// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package metadata

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the metadata type in the database.
	Label = "metadata"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProtoMessage holds the string denoting the proto_message field in the database.
	FieldProtoMessage = "proto_message"
	// FieldNativeID holds the string denoting the native_id field in the database.
	FieldNativeID = "native_id"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// EdgeTools holds the string denoting the tools edge name in mutations.
	EdgeTools = "tools"
	// EdgeAuthors holds the string denoting the authors edge name in mutations.
	EdgeAuthors = "authors"
	// EdgeDocumentTypes holds the string denoting the document_types edge name in mutations.
	EdgeDocumentTypes = "document_types"
	// EdgeSourceData holds the string denoting the source_data edge name in mutations.
	EdgeSourceData = "source_data"
	// EdgeDocument holds the string denoting the document edge name in mutations.
	EdgeDocument = "document"
	// Table holds the table name of the metadata in the database.
	Table = "metadata"
	// ToolsTable is the table that holds the tools relation/edge.
	ToolsTable = "tools"
	// ToolsInverseTable is the table name for the Tool entity.
	// It exists in this package in order to avoid circular dependency with the "tool" package.
	ToolsInverseTable = "tools"
	// ToolsColumn is the table column denoting the tools relation/edge.
	ToolsColumn = "metadata_id"
	// AuthorsTable is the table that holds the authors relation/edge.
	AuthorsTable = "persons"
	// AuthorsInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	AuthorsInverseTable = "persons"
	// AuthorsColumn is the table column denoting the authors relation/edge.
	AuthorsColumn = "metadata_id"
	// DocumentTypesTable is the table that holds the document_types relation/edge.
	DocumentTypesTable = "document_types"
	// DocumentTypesInverseTable is the table name for the DocumentType entity.
	// It exists in this package in order to avoid circular dependency with the "documenttype" package.
	DocumentTypesInverseTable = "document_types"
	// DocumentTypesColumn is the table column denoting the document_types relation/edge.
	DocumentTypesColumn = "metadata_id"
	// SourceDataTable is the table that holds the source_data relation/edge.
	SourceDataTable = "source_data"
	// SourceDataInverseTable is the table name for the SourceData entity.
	// It exists in this package in order to avoid circular dependency with the "sourcedata" package.
	SourceDataInverseTable = "source_data"
	// SourceDataColumn is the table column denoting the source_data relation/edge.
	SourceDataColumn = "metadata_id"
	// DocumentTable is the table that holds the document relation/edge.
	DocumentTable = "documents"
	// DocumentInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	DocumentInverseTable = "documents"
	// DocumentColumn is the table column denoting the document relation/edge.
	DocumentColumn = "metadata_id"
)

// Columns holds all SQL columns for metadata fields.
var Columns = []string{
	FieldID,
	FieldProtoMessage,
	FieldNativeID,
	FieldVersion,
	FieldName,
	FieldDate,
	FieldComment,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NativeIDValidator is a validator for the "native_id" field. It is called by the builders before save.
	NativeIDValidator func(string) error
)

// OrderOption defines the ordering options for the Metadata queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNativeID orders the results by the native_id field.
func ByNativeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNativeID, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
}

// ByToolsCount orders the results by tools count.
func ByToolsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newToolsStep(), opts...)
	}
}

// ByTools orders the results by tools terms.
func ByTools(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newToolsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAuthorsCount orders the results by authors count.
func ByAuthorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAuthorsStep(), opts...)
	}
}

// ByAuthors orders the results by authors terms.
func ByAuthors(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDocumentTypesCount orders the results by document_types count.
func ByDocumentTypesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDocumentTypesStep(), opts...)
	}
}

// ByDocumentTypes orders the results by document_types terms.
func ByDocumentTypes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDocumentTypesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySourceDataCount orders the results by source_data count.
func BySourceDataCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSourceDataStep(), opts...)
	}
}

// BySourceData orders the results by source_data terms.
func BySourceData(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSourceDataStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDocumentField orders the results by document field.
func ByDocumentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDocumentStep(), sql.OrderByField(field, opts...))
	}
}
func newToolsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ToolsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ToolsTable, ToolsColumn),
	)
}
func newAuthorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AuthorsTable, AuthorsColumn),
	)
}
func newDocumentTypesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DocumentTypesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DocumentTypesTable, DocumentTypesColumn),
	)
}
func newSourceDataStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SourceDataInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SourceDataTable, SourceDataColumn),
	)
}
func newDocumentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DocumentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, DocumentTable, DocumentColumn),
	)
}
