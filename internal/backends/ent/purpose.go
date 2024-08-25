// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/purpose"
)

// Purpose is the model entity for the Purpose schema.
type Purpose struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DocumentID holds the value of the "document_id" field.
	DocumentID uuid.UUID `json:"document_id,omitempty"`
	// NodeID holds the value of the "node_id" field.
	NodeID string `json:"node_id,omitempty"`
	// PrimaryPurpose holds the value of the "primary_purpose" field.
	PrimaryPurpose purpose.PrimaryPurpose `json:"primary_purpose,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PurposeQuery when eager-loading is set.
	Edges        PurposeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PurposeEdges holds the relations/edges for other nodes in the graph.
type PurposeEdges struct {
	// Document holds the value of the document edge.
	Document *Document `json:"document,omitempty"`
	// Node holds the value of the node edge.
	Node *Node `json:"node,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DocumentOrErr returns the Document value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PurposeEdges) DocumentOrErr() (*Document, error) {
	if e.Document != nil {
		return e.Document, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: document.Label}
	}
	return nil, &NotLoadedError{edge: "document"}
}

// NodeOrErr returns the Node value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PurposeEdges) NodeOrErr() (*Node, error) {
	if e.Node != nil {
		return e.Node, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: node.Label}
	}
	return nil, &NotLoadedError{edge: "node"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Purpose) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case purpose.FieldID:
			values[i] = new(sql.NullInt64)
		case purpose.FieldNodeID, purpose.FieldPrimaryPurpose:
			values[i] = new(sql.NullString)
		case purpose.FieldDocumentID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Purpose fields.
func (pu *Purpose) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case purpose.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pu.ID = int(value.Int64)
		case purpose.FieldDocumentID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field document_id", values[i])
			} else if value != nil {
				pu.DocumentID = *value
			}
		case purpose.FieldNodeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field node_id", values[i])
			} else if value.Valid {
				pu.NodeID = value.String
			}
		case purpose.FieldPrimaryPurpose:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field primary_purpose", values[i])
			} else if value.Valid {
				pu.PrimaryPurpose = purpose.PrimaryPurpose(value.String)
			}
		default:
			pu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Purpose.
// This includes values selected through modifiers, order, etc.
func (pu *Purpose) Value(name string) (ent.Value, error) {
	return pu.selectValues.Get(name)
}

// QueryDocument queries the "document" edge of the Purpose entity.
func (pu *Purpose) QueryDocument() *DocumentQuery {
	return NewPurposeClient(pu.config).QueryDocument(pu)
}

// QueryNode queries the "node" edge of the Purpose entity.
func (pu *Purpose) QueryNode() *NodeQuery {
	return NewPurposeClient(pu.config).QueryNode(pu)
}

// Update returns a builder for updating this Purpose.
// Note that you need to call Purpose.Unwrap() before calling this method if this Purpose
// was returned from a transaction, and the transaction was committed or rolled back.
func (pu *Purpose) Update() *PurposeUpdateOne {
	return NewPurposeClient(pu.config).UpdateOne(pu)
}

// Unwrap unwraps the Purpose entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pu *Purpose) Unwrap() *Purpose {
	_tx, ok := pu.config.driver.(*txDriver)
	if !ok {
		panic("ent: Purpose is not a transactional entity")
	}
	pu.config.driver = _tx.drv
	return pu
}

// String implements the fmt.Stringer.
func (pu *Purpose) String() string {
	var builder strings.Builder
	builder.WriteString("Purpose(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pu.ID))
	builder.WriteString("document_id=")
	builder.WriteString(fmt.Sprintf("%v", pu.DocumentID))
	builder.WriteString(", ")
	builder.WriteString("node_id=")
	builder.WriteString(pu.NodeID)
	builder.WriteString(", ")
	builder.WriteString("primary_purpose=")
	builder.WriteString(fmt.Sprintf("%v", pu.PrimaryPurpose))
	builder.WriteByte(')')
	return builder.String()
}

// Purposes is a parsable slice of Purpose.
type Purposes []*Purpose
