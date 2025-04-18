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
	"github.com/protobom/storage/internal/backends/ent/hashesentry"
)

// HashesEntry is the model entity for the HashesEntry schema.
type HashesEntry struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"-"`
	// HashAlgorithm holds the value of the "hash_algorithm" field.
	HashAlgorithm hashesentry.HashAlgorithm `json:"hash_algorithm,omitempty"`
	// HashData holds the value of the "hash_data" field.
	HashData string `json:"hash_data,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HashesEntryQuery when eager-loading is set.
	Edges        HashesEntryEdges `json:"-"`
	selectValues sql.SelectValues
}

// HashesEntryEdges holds the relations/edges for other nodes in the graph.
type HashesEntryEdges struct {
	// Documents holds the value of the documents edge.
	Documents []*Document `json:"-"`
	// ExternalReferences holds the value of the external_references edge.
	ExternalReferences []*ExternalReference `json:"-"`
	// Nodes holds the value of the nodes edge.
	Nodes []*Node `json:"-"`
	// SourceData holds the value of the source_data edge.
	SourceData []*SourceData `json:"-"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// DocumentsOrErr returns the Documents value or an error if the edge
// was not loaded in eager-loading.
func (e HashesEntryEdges) DocumentsOrErr() ([]*Document, error) {
	if e.loadedTypes[0] {
		return e.Documents, nil
	}
	return nil, &NotLoadedError{edge: "documents"}
}

// ExternalReferencesOrErr returns the ExternalReferences value or an error if the edge
// was not loaded in eager-loading.
func (e HashesEntryEdges) ExternalReferencesOrErr() ([]*ExternalReference, error) {
	if e.loadedTypes[1] {
		return e.ExternalReferences, nil
	}
	return nil, &NotLoadedError{edge: "external_references"}
}

// NodesOrErr returns the Nodes value or an error if the edge
// was not loaded in eager-loading.
func (e HashesEntryEdges) NodesOrErr() ([]*Node, error) {
	if e.loadedTypes[2] {
		return e.Nodes, nil
	}
	return nil, &NotLoadedError{edge: "nodes"}
}

// SourceDataOrErr returns the SourceData value or an error if the edge
// was not loaded in eager-loading.
func (e HashesEntryEdges) SourceDataOrErr() ([]*SourceData, error) {
	if e.loadedTypes[3] {
		return e.SourceData, nil
	}
	return nil, &NotLoadedError{edge: "source_data"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HashesEntry) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hashesentry.FieldHashAlgorithm, hashesentry.FieldHashData:
			values[i] = new(sql.NullString)
		case hashesentry.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HashesEntry fields.
func (he *HashesEntry) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hashesentry.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				he.ID = *value
			}
		case hashesentry.FieldHashAlgorithm:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash_algorithm", values[i])
			} else if value.Valid {
				he.HashAlgorithm = hashesentry.HashAlgorithm(value.String)
			}
		case hashesentry.FieldHashData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash_data", values[i])
			} else if value.Valid {
				he.HashData = value.String
			}
		default:
			he.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the HashesEntry.
// This includes values selected through modifiers, order, etc.
func (he *HashesEntry) Value(name string) (ent.Value, error) {
	return he.selectValues.Get(name)
}

// QueryDocuments queries the "documents" edge of the HashesEntry entity.
func (he *HashesEntry) QueryDocuments() *DocumentQuery {
	return NewHashesEntryClient(he.config).QueryDocuments(he)
}

// QueryExternalReferences queries the "external_references" edge of the HashesEntry entity.
func (he *HashesEntry) QueryExternalReferences() *ExternalReferenceQuery {
	return NewHashesEntryClient(he.config).QueryExternalReferences(he)
}

// QueryNodes queries the "nodes" edge of the HashesEntry entity.
func (he *HashesEntry) QueryNodes() *NodeQuery {
	return NewHashesEntryClient(he.config).QueryNodes(he)
}

// QuerySourceData queries the "source_data" edge of the HashesEntry entity.
func (he *HashesEntry) QuerySourceData() *SourceDataQuery {
	return NewHashesEntryClient(he.config).QuerySourceData(he)
}

// Update returns a builder for updating this HashesEntry.
// Note that you need to call HashesEntry.Unwrap() before calling this method if this HashesEntry
// was returned from a transaction, and the transaction was committed or rolled back.
func (he *HashesEntry) Update() *HashesEntryUpdateOne {
	return NewHashesEntryClient(he.config).UpdateOne(he)
}

// Unwrap unwraps the HashesEntry entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (he *HashesEntry) Unwrap() *HashesEntry {
	_tx, ok := he.config.driver.(*txDriver)
	if !ok {
		panic("ent: HashesEntry is not a transactional entity")
	}
	he.config.driver = _tx.drv
	return he
}

// String implements the fmt.Stringer.
func (he *HashesEntry) String() string {
	var builder strings.Builder
	builder.WriteString("HashesEntry(")
	builder.WriteString(fmt.Sprintf("id=%v, ", he.ID))
	builder.WriteString("hash_algorithm=")
	builder.WriteString(fmt.Sprintf("%v", he.HashAlgorithm))
	builder.WriteString(", ")
	builder.WriteString("hash_data=")
	builder.WriteString(he.HashData)
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (he *HashesEntry) MarshalJSON() ([]byte, error) {
	type Alias HashesEntry
	return json.Marshal(&struct {
		*Alias
		HashesEntryEdges
	}{
		Alias:            (*Alias)(he),
		HashesEntryEdges: he.Edges,
	})
}

// HashesEntries is a parsable slice of HashesEntry.
type HashesEntries []*HashesEntry
