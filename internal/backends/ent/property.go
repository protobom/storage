// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/storage/internal/backends/ent/property"
)

// Property is the model entity for the Property schema.
type Property struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"-"`
	// ProtoMessage holds the value of the "proto_message" field.
	ProtoMessage *sbom.Property `json:"-"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Data holds the value of the "data" field.
	Data string `json:"data,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PropertyQuery when eager-loading is set.
	Edges        PropertyEdges `json:"-"`
	selectValues sql.SelectValues
}

// PropertyEdges holds the relations/edges for other nodes in the graph.
type PropertyEdges struct {
	// Documents holds the value of the documents edge.
	Documents []*Document `json:"-"`
	// Nodes holds the value of the nodes edge.
	Nodes []*Node `json:"-"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DocumentsOrErr returns the Documents value or an error if the edge
// was not loaded in eager-loading.
func (e PropertyEdges) DocumentsOrErr() ([]*Document, error) {
	if e.loadedTypes[0] {
		return e.Documents, nil
	}
	return nil, &NotLoadedError{edge: "documents"}
}

// NodesOrErr returns the Nodes value or an error if the edge
// was not loaded in eager-loading.
func (e PropertyEdges) NodesOrErr() ([]*Node, error) {
	if e.loadedTypes[1] {
		return e.Nodes, nil
	}
	return nil, &NotLoadedError{edge: "nodes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Property) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case property.FieldProtoMessage:
			values[i] = &sql.NullScanner{S: new(sbom.Property)}
		case property.FieldName, property.FieldData:
			values[i] = new(sql.NullString)
		case property.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Property fields.
func (pr *Property) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case property.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case property.FieldProtoMessage:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field proto_message", values[i])
			} else if value.Valid {
				pr.ProtoMessage = value.S.(*sbom.Property)
			}
		case property.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case property.FieldData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value.Valid {
				pr.Data = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Property.
// This includes values selected through modifiers, order, etc.
func (pr *Property) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryDocuments queries the "documents" edge of the Property entity.
func (pr *Property) QueryDocuments() *DocumentQuery {
	return NewPropertyClient(pr.config).QueryDocuments(pr)
}

// QueryNodes queries the "nodes" edge of the Property entity.
func (pr *Property) QueryNodes() *NodeQuery {
	return NewPropertyClient(pr.config).QueryNodes(pr)
}

// Update returns a builder for updating this Property.
// Note that you need to call Property.Unwrap() before calling this method if this Property
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Property) Update() *PropertyUpdateOne {
	return NewPropertyClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Property entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Property) Unwrap() *Property {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Property is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Property) String() string {
	var builder strings.Builder
	builder.WriteString("Property(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	if v := pr.ProtoMessage; v != nil {
		builder.WriteString("proto_message=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(pr.Data)
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (pr *Property) MarshalJSON() ([]byte, error) {
	type Alias Property
	return json.Marshal(&struct {
		*Alias
		PropertyEdges
	}{
		Alias:         (*Alias)(pr),
		PropertyEdges: pr.Edges,
	})
}

// Properties is a parsable slice of Property.
type Properties []*Property
