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
	"github.com/protobom/storage/internal/backends/ent/sourcedata"
)

// SourceData is the model entity for the SourceData schema.
type SourceData struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"-"`
	// ProtoMessage holds the value of the "proto_message" field.
	ProtoMessage *sbom.SourceData `json:"-"`
	// Format holds the value of the "format" field.
	Format string `json:"format,omitempty"`
	// Size holds the value of the "size" field.
	Size int64 `json:"size,omitempty"`
	// URI holds the value of the "uri" field.
	URI *string `json:"uri,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SourceDataQuery when eager-loading is set.
	Edges        SourceDataEdges `json:"-"`
	selectValues sql.SelectValues
}

// SourceDataEdges holds the relations/edges for other nodes in the graph.
type SourceDataEdges struct {
	// Hashes holds the value of the hashes edge.
	Hashes []*HashesEntry `json:"hashes,omitempty"`
	// Documents holds the value of the documents edge.
	Documents []*Document `json:"-"`
	// Metadata holds the value of the metadata edge.
	Metadata []*Metadata `json:"-"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// HashesOrErr returns the Hashes value or an error if the edge
// was not loaded in eager-loading.
func (e SourceDataEdges) HashesOrErr() ([]*HashesEntry, error) {
	if e.loadedTypes[0] {
		return e.Hashes, nil
	}
	return nil, &NotLoadedError{edge: "hashes"}
}

// DocumentsOrErr returns the Documents value or an error if the edge
// was not loaded in eager-loading.
func (e SourceDataEdges) DocumentsOrErr() ([]*Document, error) {
	if e.loadedTypes[1] {
		return e.Documents, nil
	}
	return nil, &NotLoadedError{edge: "documents"}
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading.
func (e SourceDataEdges) MetadataOrErr() ([]*Metadata, error) {
	if e.loadedTypes[2] {
		return e.Metadata, nil
	}
	return nil, &NotLoadedError{edge: "metadata"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SourceData) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sourcedata.FieldProtoMessage:
			values[i] = &sql.NullScanner{S: new(sbom.SourceData)}
		case sourcedata.FieldSize:
			values[i] = new(sql.NullInt64)
		case sourcedata.FieldFormat, sourcedata.FieldURI:
			values[i] = new(sql.NullString)
		case sourcedata.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SourceData fields.
func (sd *SourceData) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sourcedata.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sd.ID = *value
			}
		case sourcedata.FieldProtoMessage:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field proto_message", values[i])
			} else if value.Valid {
				sd.ProtoMessage = value.S.(*sbom.SourceData)
			}
		case sourcedata.FieldFormat:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field format", values[i])
			} else if value.Valid {
				sd.Format = value.String
			}
		case sourcedata.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				sd.Size = value.Int64
			}
		case sourcedata.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				sd.URI = new(string)
				*sd.URI = value.String
			}
		default:
			sd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SourceData.
// This includes values selected through modifiers, order, etc.
func (sd *SourceData) Value(name string) (ent.Value, error) {
	return sd.selectValues.Get(name)
}

// QueryHashes queries the "hashes" edge of the SourceData entity.
func (sd *SourceData) QueryHashes() *HashesEntryQuery {
	return NewSourceDataClient(sd.config).QueryHashes(sd)
}

// QueryDocuments queries the "documents" edge of the SourceData entity.
func (sd *SourceData) QueryDocuments() *DocumentQuery {
	return NewSourceDataClient(sd.config).QueryDocuments(sd)
}

// QueryMetadata queries the "metadata" edge of the SourceData entity.
func (sd *SourceData) QueryMetadata() *MetadataQuery {
	return NewSourceDataClient(sd.config).QueryMetadata(sd)
}

// Update returns a builder for updating this SourceData.
// Note that you need to call SourceData.Unwrap() before calling this method if this SourceData
// was returned from a transaction, and the transaction was committed or rolled back.
func (sd *SourceData) Update() *SourceDataUpdateOne {
	return NewSourceDataClient(sd.config).UpdateOne(sd)
}

// Unwrap unwraps the SourceData entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sd *SourceData) Unwrap() *SourceData {
	_tx, ok := sd.config.driver.(*txDriver)
	if !ok {
		panic("ent: SourceData is not a transactional entity")
	}
	sd.config.driver = _tx.drv
	return sd
}

// String implements the fmt.Stringer.
func (sd *SourceData) String() string {
	var builder strings.Builder
	builder.WriteString("SourceData(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sd.ID))
	if v := sd.ProtoMessage; v != nil {
		builder.WriteString("proto_message=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("format=")
	builder.WriteString(sd.Format)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", sd.Size))
	builder.WriteString(", ")
	if v := sd.URI; v != nil {
		builder.WriteString("uri=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (sd *SourceData) MarshalJSON() ([]byte, error) {
	type Alias SourceData
	return json.Marshal(&struct {
		*Alias
		SourceDataEdges
	}{
		Alias:           (*Alias)(sd),
		SourceDataEdges: sd.Edges,
	})
}

// SourceDataSlice is a parsable slice of SourceData.
type SourceDataSlice []*SourceData
