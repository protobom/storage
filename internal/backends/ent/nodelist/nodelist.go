// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package nodelist

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the nodelist type in the database.
	Label = "node_list"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProtoMessage holds the string denoting the proto_message field in the database.
	FieldProtoMessage = "proto_message"
	// FieldRootElements holds the string denoting the root_elements field in the database.
	FieldRootElements = "root_elements"
	// EdgeEdgeTypes holds the string denoting the edge_types edge name in mutations.
	EdgeEdgeTypes = "edge_types"
	// EdgeNodes holds the string denoting the nodes edge name in mutations.
	EdgeNodes = "nodes"
	// EdgeDocuments holds the string denoting the documents edge name in mutations.
	EdgeDocuments = "documents"
	// Table holds the table name of the nodelist in the database.
	Table = "node_lists"
	// EdgeTypesTable is the table that holds the edge_types relation/edge. The primary key declared below.
	EdgeTypesTable = "node_list_edges"
	// EdgeTypesInverseTable is the table name for the EdgeType entity.
	// It exists in this package in order to avoid circular dependency with the "edgetype" package.
	EdgeTypesInverseTable = "edge_types"
	// NodesTable is the table that holds the nodes relation/edge. The primary key declared below.
	NodesTable = "node_list_nodes"
	// NodesInverseTable is the table name for the Node entity.
	// It exists in this package in order to avoid circular dependency with the "node" package.
	NodesInverseTable = "nodes"
	// DocumentsTable is the table that holds the documents relation/edge.
	DocumentsTable = "documents"
	// DocumentsInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	DocumentsInverseTable = "documents"
	// DocumentsColumn is the table column denoting the documents relation/edge.
	DocumentsColumn = "node_list_id"
)

// Columns holds all SQL columns for nodelist fields.
var Columns = []string{
	FieldID,
	FieldProtoMessage,
	FieldRootElements,
}

var (
	// EdgeTypesPrimaryKey and EdgeTypesColumn2 are the table columns denoting the
	// primary key for the edge_types relation (M2M).
	EdgeTypesPrimaryKey = []string{"node_list_id", "edge_type_id"}
	// NodesPrimaryKey and NodesColumn2 are the table columns denoting the
	// primary key for the nodes relation (M2M).
	NodesPrimaryKey = []string{"node_list_id", "node_id"}
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

// OrderOption defines the ordering options for the NodeList queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEdgeTypesCount orders the results by edge_types count.
func ByEdgeTypesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEdgeTypesStep(), opts...)
	}
}

// ByEdgeTypes orders the results by edge_types terms.
func ByEdgeTypes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEdgeTypesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNodesCount orders the results by nodes count.
func ByNodesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNodesStep(), opts...)
	}
}

// ByNodes orders the results by nodes terms.
func ByNodes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
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
func newEdgeTypesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EdgeTypesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, EdgeTypesTable, EdgeTypesPrimaryKey...),
	)
}
func newNodesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, NodesTable, NodesPrimaryKey...),
	)
}
func newDocumentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DocumentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, DocumentsTable, DocumentsColumn),
	)
}
