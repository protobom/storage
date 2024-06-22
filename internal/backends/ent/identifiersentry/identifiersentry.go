// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package identifiersentry

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the identifiersentry type in the database.
	Label = "identifiers_entry"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNodeID holds the string denoting the node_id field in the database.
	FieldNodeID = "node_id"
	// FieldSoftwareIdentifierType holds the string denoting the software_identifier_type field in the database.
	FieldSoftwareIdentifierType = "software_identifier_type"
	// FieldSoftwareIdentifierValue holds the string denoting the software_identifier_value field in the database.
	FieldSoftwareIdentifierValue = "software_identifier_value"
	// EdgeNode holds the string denoting the node edge name in mutations.
	EdgeNode = "node"
	// Table holds the table name of the identifiersentry in the database.
	Table = "identifiers_entries"
	// NodeTable is the table that holds the node relation/edge.
	NodeTable = "identifiers_entries"
	// NodeInverseTable is the table name for the Node entity.
	// It exists in this package in order to avoid circular dependency with the "node" package.
	NodeInverseTable = "nodes"
	// NodeColumn is the table column denoting the node relation/edge.
	NodeColumn = "node_id"
)

// Columns holds all SQL columns for identifiersentry fields.
var Columns = []string{
	FieldID,
	FieldNodeID,
	FieldSoftwareIdentifierType,
	FieldSoftwareIdentifierValue,
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

// SoftwareIdentifierType defines the type for the "software_identifier_type" enum field.
type SoftwareIdentifierType string

// SoftwareIdentifierType values.
const (
	SoftwareIdentifierTypeUNKNOWN_IDENTIFIER_TYPE SoftwareIdentifierType = "UNKNOWN_IDENTIFIER_TYPE"
	SoftwareIdentifierTypePURL                    SoftwareIdentifierType = "PURL"
	SoftwareIdentifierTypeCPE22                   SoftwareIdentifierType = "CPE22"
	SoftwareIdentifierTypeCPE23                   SoftwareIdentifierType = "CPE23"
	SoftwareIdentifierTypeGITOID                  SoftwareIdentifierType = "GITOID"
)

func (sit SoftwareIdentifierType) String() string {
	return string(sit)
}

// SoftwareIdentifierTypeValidator is a validator for the "software_identifier_type" field enum values. It is called by the builders before save.
func SoftwareIdentifierTypeValidator(sit SoftwareIdentifierType) error {
	switch sit {
	case SoftwareIdentifierTypeUNKNOWN_IDENTIFIER_TYPE, SoftwareIdentifierTypePURL, SoftwareIdentifierTypeCPE22, SoftwareIdentifierTypeCPE23, SoftwareIdentifierTypeGITOID:
		return nil
	default:
		return fmt.Errorf("identifiersentry: invalid enum value for software_identifier_type field: %q", sit)
	}
}

// OrderOption defines the ordering options for the IdentifiersEntry queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNodeID orders the results by the node_id field.
func ByNodeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNodeID, opts...).ToFunc()
}

// BySoftwareIdentifierType orders the results by the software_identifier_type field.
func BySoftwareIdentifierType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSoftwareIdentifierType, opts...).ToFunc()
}

// BySoftwareIdentifierValue orders the results by the software_identifier_value field.
func BySoftwareIdentifierValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSoftwareIdentifierValue, opts...).ToFunc()
}

// ByNodeField orders the results by node field.
func ByNodeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodeStep(), sql.OrderByField(field, opts...))
	}
}
func newNodeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, NodeTable, NodeColumn),
	)
}
