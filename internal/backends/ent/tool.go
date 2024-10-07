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
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/metadata"
	"github.com/protobom/storage/internal/backends/ent/tool"
)

// Tool is the model entity for the Tool schema.
type Tool struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// DocumentID holds the value of the "document_id" field.
	DocumentID uuid.UUID `json:"document_id,omitempty"`
	// ProtoMessage holds the value of the "proto_message" field.
	ProtoMessage *sbom.Tool `json:"proto_message,omitempty"`
	// MetadataID holds the value of the "metadata_id" field.
	MetadataID string `json:"metadata_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Vendor holds the value of the "vendor" field.
	Vendor string `json:"vendor,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ToolQuery when eager-loading is set.
	Edges        ToolEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ToolEdges holds the relations/edges for other nodes in the graph.
type ToolEdges struct {
	// Document holds the value of the document edge.
	Document *Document `json:"document,omitempty"`
	// Metadata holds the value of the metadata edge.
	Metadata *Metadata `json:"metadata,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DocumentOrErr returns the Document value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ToolEdges) DocumentOrErr() (*Document, error) {
	if e.Document != nil {
		return e.Document, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: document.Label}
	}
	return nil, &NotLoadedError{edge: "document"}
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ToolEdges) MetadataOrErr() (*Metadata, error) {
	if e.Metadata != nil {
		return e.Metadata, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: metadata.Label}
	}
	return nil, &NotLoadedError{edge: "metadata"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tool) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tool.FieldProtoMessage:
			values[i] = new(sbom.Tool)
		case tool.FieldMetadataID, tool.FieldName, tool.FieldVersion, tool.FieldVendor:
			values[i] = new(sql.NullString)
		case tool.FieldID, tool.FieldDocumentID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tool fields.
func (t *Tool) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tool.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case tool.FieldDocumentID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field document_id", values[i])
			} else if value != nil {
				t.DocumentID = *value
			}
		case tool.FieldProtoMessage:
			if value, ok := values[i].(*sbom.Tool); !ok {
				return fmt.Errorf("unexpected type %T for field proto_message", values[i])
			} else if value != nil {
				t.ProtoMessage = value
			}
		case tool.FieldMetadataID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field metadata_id", values[i])
			} else if value.Valid {
				t.MetadataID = value.String
			}
		case tool.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case tool.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				t.Version = value.String
			}
		case tool.FieldVendor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vendor", values[i])
			} else if value.Valid {
				t.Vendor = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Tool.
// This includes values selected through modifiers, order, etc.
func (t *Tool) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryDocument queries the "document" edge of the Tool entity.
func (t *Tool) QueryDocument() *DocumentQuery {
	return NewToolClient(t.config).QueryDocument(t)
}

// QueryMetadata queries the "metadata" edge of the Tool entity.
func (t *Tool) QueryMetadata() *MetadataQuery {
	return NewToolClient(t.config).QueryMetadata(t)
}

// Update returns a builder for updating this Tool.
// Note that you need to call Tool.Unwrap() before calling this method if this Tool
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tool) Update() *ToolUpdateOne {
	return NewToolClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tool entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tool) Unwrap() *Tool {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tool is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tool) String() string {
	var builder strings.Builder
	builder.WriteString("Tool(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("document_id=")
	builder.WriteString(fmt.Sprintf("%v", t.DocumentID))
	builder.WriteString(", ")
	builder.WriteString("proto_message=")
	builder.WriteString(fmt.Sprintf("%v", t.ProtoMessage))
	builder.WriteString(", ")
	builder.WriteString("metadata_id=")
	builder.WriteString(t.MetadataID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(t.Version)
	builder.WriteString(", ")
	builder.WriteString("vendor=")
	builder.WriteString(t.Vendor)
	builder.WriteByte(')')
	return builder.String()
}

// Tools is a parsable slice of Tool.
type Tools []*Tool
