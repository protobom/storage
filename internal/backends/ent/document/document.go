// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package document

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the document type in the database.
	Label = "document"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProtoMessage holds the string denoting the proto_message field in the database.
	FieldProtoMessage = "proto_message"
	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// EdgeNodeList holds the string denoting the node_list edge name in mutations.
	EdgeNodeList = "node_list"
	// Table holds the table name of the document in the database.
	Table = "documents"
	// MetadataTable is the table that holds the metadata relation/edge.
	MetadataTable = "metadata"
	// MetadataInverseTable is the table name for the Metadata entity.
	// It exists in this package in order to avoid circular dependency with the "metadata" package.
	MetadataInverseTable = "metadata"
	// MetadataColumn is the table column denoting the metadata relation/edge.
	MetadataColumn = "id"
	// NodeListTable is the table that holds the node_list relation/edge.
	NodeListTable = "documents"
	// NodeListInverseTable is the table name for the NodeList entity.
	// It exists in this package in order to avoid circular dependency with the "nodelist" package.
	NodeListInverseTable = "node_lists"
	// NodeListColumn is the table column denoting the node_list relation/edge.
	NodeListColumn = "document_id"
)

// Columns holds all SQL columns for document fields.
var Columns = []string{
	FieldID,
	FieldProtoMessage,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "documents"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"document_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Document queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMetadataField orders the results by metadata field.
func ByMetadataField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMetadataStep(), sql.OrderByField(field, opts...))
	}
}

// ByNodeListField orders the results by node_list field.
func ByNodeListField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodeListStep(), sql.OrderByField(field, opts...))
	}
}
func newMetadataStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MetadataInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, MetadataTable, MetadataColumn),
	)
}
func newNodeListStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodeListInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, NodeListTable, NodeListColumn),
	)
}
