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
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/node"
)

// Node is the model entity for the Node schema.
type Node struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"-"`
	// DocumentID holds the value of the "document_id" field.
	DocumentID uuid.UUID `json:"-"`
	// ProtoMessage holds the value of the "proto_message" field.
	ProtoMessage *sbom.Node `json:"-"`
	// NativeID holds the value of the "native_id" field.
	NativeID string `json:"native_id,omitempty"`
	// NodeListID holds the value of the "node_list_id" field.
	NodeListID uuid.UUID `json:"node_list_id,omitempty"`
	// Type holds the value of the "type" field.
	Type node.Type `json:"type,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// FileName holds the value of the "file_name" field.
	FileName string `json:"file_name,omitempty"`
	// URLHome holds the value of the "url_home" field.
	URLHome string `json:"url_home,omitempty"`
	// URLDownload holds the value of the "url_download" field.
	URLDownload string `json:"url_download,omitempty"`
	// Licenses holds the value of the "licenses" field.
	Licenses []string `json:"licenses,omitempty"`
	// LicenseConcluded holds the value of the "license_concluded" field.
	LicenseConcluded string `json:"license_concluded,omitempty"`
	// LicenseComments holds the value of the "license_comments" field.
	LicenseComments string `json:"license_comments,omitempty"`
	// Copyright holds the value of the "copyright" field.
	Copyright string `json:"copyright,omitempty"`
	// SourceInfo holds the value of the "source_info" field.
	SourceInfo string `json:"source_info,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment string `json:"comment,omitempty"`
	// Summary holds the value of the "summary" field.
	Summary string `json:"summary,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ReleaseDate holds the value of the "release_date" field.
	ReleaseDate time.Time `json:"release_date,omitempty"`
	// BuildDate holds the value of the "build_date" field.
	BuildDate time.Time `json:"build_date,omitempty"`
	// ValidUntilDate holds the value of the "valid_until_date" field.
	ValidUntilDate time.Time `json:"valid_until_date,omitempty"`
	// Attribution holds the value of the "attribution" field.
	Attribution []string `json:"attribution,omitempty"`
	// FileTypes holds the value of the "file_types" field.
	FileTypes []string `json:"file_types,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NodeQuery when eager-loading is set.
	Edges        NodeEdges `json:"-"`
	selectValues sql.SelectValues
}

// NodeEdges holds the relations/edges for other nodes in the graph.
type NodeEdges struct {
	// Document holds the value of the document edge.
	Document *Document `json:"document,omitempty"`
	// Annotations holds the value of the annotations edge.
	Annotations []*Annotation `json:"annotations,omitempty"`
	// Suppliers holds the value of the suppliers edge.
	Suppliers []*Person `json:"suppliers,omitempty"`
	// Originators holds the value of the originators edge.
	Originators []*Person `json:"originators,omitempty"`
	// ExternalReferences holds the value of the external_references edge.
	ExternalReferences []*ExternalReference `json:"external_references,omitempty"`
	// PrimaryPurpose holds the value of the primary_purpose edge.
	PrimaryPurpose []*Purpose `json:"primary_purpose,omitempty"`
	// ToNodes holds the value of the to_nodes edge.
	ToNodes []*Node `json:"-"`
	// Nodes holds the value of the nodes edge.
	Nodes []*Node `json:"nodes,omitempty"`
	// Hashes holds the value of the hashes edge.
	Hashes []*HashesEntry `json:"hashes,omitempty"`
	// Identifiers holds the value of the identifiers edge.
	Identifiers []*IdentifiersEntry `json:"identifiers,omitempty"`
	// Properties holds the value of the properties edge.
	Properties []*Property `json:"properties,omitempty"`
	// NodeLists holds the value of the node_lists edge.
	NodeLists []*NodeList `json:"-"`
	// EdgeTypes holds the value of the edge_types edge.
	EdgeTypes []*EdgeType `json:"-"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [13]bool
}

// DocumentOrErr returns the Document value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NodeEdges) DocumentOrErr() (*Document, error) {
	if e.Document != nil {
		return e.Document, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: document.Label}
	}
	return nil, &NotLoadedError{edge: "document"}
}

// AnnotationsOrErr returns the Annotations value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) AnnotationsOrErr() ([]*Annotation, error) {
	if e.loadedTypes[1] {
		return e.Annotations, nil
	}
	return nil, &NotLoadedError{edge: "annotations"}
}

// SuppliersOrErr returns the Suppliers value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) SuppliersOrErr() ([]*Person, error) {
	if e.loadedTypes[2] {
		return e.Suppliers, nil
	}
	return nil, &NotLoadedError{edge: "suppliers"}
}

// OriginatorsOrErr returns the Originators value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) OriginatorsOrErr() ([]*Person, error) {
	if e.loadedTypes[3] {
		return e.Originators, nil
	}
	return nil, &NotLoadedError{edge: "originators"}
}

// ExternalReferencesOrErr returns the ExternalReferences value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) ExternalReferencesOrErr() ([]*ExternalReference, error) {
	if e.loadedTypes[4] {
		return e.ExternalReferences, nil
	}
	return nil, &NotLoadedError{edge: "external_references"}
}

// PrimaryPurposeOrErr returns the PrimaryPurpose value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) PrimaryPurposeOrErr() ([]*Purpose, error) {
	if e.loadedTypes[5] {
		return e.PrimaryPurpose, nil
	}
	return nil, &NotLoadedError{edge: "primary_purpose"}
}

// ToNodesOrErr returns the ToNodes value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) ToNodesOrErr() ([]*Node, error) {
	if e.loadedTypes[6] {
		return e.ToNodes, nil
	}
	return nil, &NotLoadedError{edge: "to_nodes"}
}

// NodesOrErr returns the Nodes value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) NodesOrErr() ([]*Node, error) {
	if e.loadedTypes[7] {
		return e.Nodes, nil
	}
	return nil, &NotLoadedError{edge: "nodes"}
}

// HashesOrErr returns the Hashes value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) HashesOrErr() ([]*HashesEntry, error) {
	if e.loadedTypes[8] {
		return e.Hashes, nil
	}
	return nil, &NotLoadedError{edge: "hashes"}
}

// IdentifiersOrErr returns the Identifiers value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) IdentifiersOrErr() ([]*IdentifiersEntry, error) {
	if e.loadedTypes[9] {
		return e.Identifiers, nil
	}
	return nil, &NotLoadedError{edge: "identifiers"}
}

// PropertiesOrErr returns the Properties value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) PropertiesOrErr() ([]*Property, error) {
	if e.loadedTypes[10] {
		return e.Properties, nil
	}
	return nil, &NotLoadedError{edge: "properties"}
}

// NodeListsOrErr returns the NodeLists value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) NodeListsOrErr() ([]*NodeList, error) {
	if e.loadedTypes[11] {
		return e.NodeLists, nil
	}
	return nil, &NotLoadedError{edge: "node_lists"}
}

// EdgeTypesOrErr returns the EdgeTypes value or an error if the edge
// was not loaded in eager-loading.
func (e NodeEdges) EdgeTypesOrErr() ([]*EdgeType, error) {
	if e.loadedTypes[12] {
		return e.EdgeTypes, nil
	}
	return nil, &NotLoadedError{edge: "edge_types"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Node) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case node.FieldProtoMessage:
			values[i] = &sql.NullScanner{S: new(sbom.Node)}
		case node.FieldLicenses, node.FieldAttribution, node.FieldFileTypes:
			values[i] = new([]byte)
		case node.FieldNativeID, node.FieldType, node.FieldName, node.FieldVersion, node.FieldFileName, node.FieldURLHome, node.FieldURLDownload, node.FieldLicenseConcluded, node.FieldLicenseComments, node.FieldCopyright, node.FieldSourceInfo, node.FieldComment, node.FieldSummary, node.FieldDescription:
			values[i] = new(sql.NullString)
		case node.FieldReleaseDate, node.FieldBuildDate, node.FieldValidUntilDate:
			values[i] = new(sql.NullTime)
		case node.FieldID, node.FieldDocumentID, node.FieldNodeListID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Node fields.
func (n *Node) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case node.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				n.ID = *value
			}
		case node.FieldDocumentID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field document_id", values[i])
			} else if value != nil {
				n.DocumentID = *value
			}
		case node.FieldProtoMessage:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field proto_message", values[i])
			} else if value.Valid {
				n.ProtoMessage = value.S.(*sbom.Node)
			}
		case node.FieldNativeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field native_id", values[i])
			} else if value.Valid {
				n.NativeID = value.String
			}
		case node.FieldNodeListID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field node_list_id", values[i])
			} else if value != nil {
				n.NodeListID = *value
			}
		case node.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				n.Type = node.Type(value.String)
			}
		case node.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				n.Name = value.String
			}
		case node.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				n.Version = value.String
			}
		case node.FieldFileName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_name", values[i])
			} else if value.Valid {
				n.FileName = value.String
			}
		case node.FieldURLHome:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url_home", values[i])
			} else if value.Valid {
				n.URLHome = value.String
			}
		case node.FieldURLDownload:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url_download", values[i])
			} else if value.Valid {
				n.URLDownload = value.String
			}
		case node.FieldLicenses:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field licenses", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &n.Licenses); err != nil {
					return fmt.Errorf("unmarshal field licenses: %w", err)
				}
			}
		case node.FieldLicenseConcluded:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field license_concluded", values[i])
			} else if value.Valid {
				n.LicenseConcluded = value.String
			}
		case node.FieldLicenseComments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field license_comments", values[i])
			} else if value.Valid {
				n.LicenseComments = value.String
			}
		case node.FieldCopyright:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field copyright", values[i])
			} else if value.Valid {
				n.Copyright = value.String
			}
		case node.FieldSourceInfo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_info", values[i])
			} else if value.Valid {
				n.SourceInfo = value.String
			}
		case node.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				n.Comment = value.String
			}
		case node.FieldSummary:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field summary", values[i])
			} else if value.Valid {
				n.Summary = value.String
			}
		case node.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				n.Description = value.String
			}
		case node.FieldReleaseDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field release_date", values[i])
			} else if value.Valid {
				n.ReleaseDate = value.Time
			}
		case node.FieldBuildDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field build_date", values[i])
			} else if value.Valid {
				n.BuildDate = value.Time
			}
		case node.FieldValidUntilDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field valid_until_date", values[i])
			} else if value.Valid {
				n.ValidUntilDate = value.Time
			}
		case node.FieldAttribution:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field attribution", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &n.Attribution); err != nil {
					return fmt.Errorf("unmarshal field attribution: %w", err)
				}
			}
		case node.FieldFileTypes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field file_types", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &n.FileTypes); err != nil {
					return fmt.Errorf("unmarshal field file_types: %w", err)
				}
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Node.
// This includes values selected through modifiers, order, etc.
func (n *Node) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryDocument queries the "document" edge of the Node entity.
func (n *Node) QueryDocument() *DocumentQuery {
	return NewNodeClient(n.config).QueryDocument(n)
}

// QueryAnnotations queries the "annotations" edge of the Node entity.
func (n *Node) QueryAnnotations() *AnnotationQuery {
	return NewNodeClient(n.config).QueryAnnotations(n)
}

// QuerySuppliers queries the "suppliers" edge of the Node entity.
func (n *Node) QuerySuppliers() *PersonQuery {
	return NewNodeClient(n.config).QuerySuppliers(n)
}

// QueryOriginators queries the "originators" edge of the Node entity.
func (n *Node) QueryOriginators() *PersonQuery {
	return NewNodeClient(n.config).QueryOriginators(n)
}

// QueryExternalReferences queries the "external_references" edge of the Node entity.
func (n *Node) QueryExternalReferences() *ExternalReferenceQuery {
	return NewNodeClient(n.config).QueryExternalReferences(n)
}

// QueryPrimaryPurpose queries the "primary_purpose" edge of the Node entity.
func (n *Node) QueryPrimaryPurpose() *PurposeQuery {
	return NewNodeClient(n.config).QueryPrimaryPurpose(n)
}

// QueryToNodes queries the "to_nodes" edge of the Node entity.
func (n *Node) QueryToNodes() *NodeQuery {
	return NewNodeClient(n.config).QueryToNodes(n)
}

// QueryNodes queries the "nodes" edge of the Node entity.
func (n *Node) QueryNodes() *NodeQuery {
	return NewNodeClient(n.config).QueryNodes(n)
}

// QueryHashes queries the "hashes" edge of the Node entity.
func (n *Node) QueryHashes() *HashesEntryQuery {
	return NewNodeClient(n.config).QueryHashes(n)
}

// QueryIdentifiers queries the "identifiers" edge of the Node entity.
func (n *Node) QueryIdentifiers() *IdentifiersEntryQuery {
	return NewNodeClient(n.config).QueryIdentifiers(n)
}

// QueryProperties queries the "properties" edge of the Node entity.
func (n *Node) QueryProperties() *PropertyQuery {
	return NewNodeClient(n.config).QueryProperties(n)
}

// QueryNodeLists queries the "node_lists" edge of the Node entity.
func (n *Node) QueryNodeLists() *NodeListQuery {
	return NewNodeClient(n.config).QueryNodeLists(n)
}

// QueryEdgeTypes queries the "edge_types" edge of the Node entity.
func (n *Node) QueryEdgeTypes() *EdgeTypeQuery {
	return NewNodeClient(n.config).QueryEdgeTypes(n)
}

// Update returns a builder for updating this Node.
// Note that you need to call Node.Unwrap() before calling this method if this Node
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Node) Update() *NodeUpdateOne {
	return NewNodeClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Node entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Node) Unwrap() *Node {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Node is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Node) String() string {
	var builder strings.Builder
	builder.WriteString("Node(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("document_id=")
	builder.WriteString(fmt.Sprintf("%v", n.DocumentID))
	builder.WriteString(", ")
	if v := n.ProtoMessage; v != nil {
		builder.WriteString("proto_message=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("native_id=")
	builder.WriteString(n.NativeID)
	builder.WriteString(", ")
	builder.WriteString("node_list_id=")
	builder.WriteString(fmt.Sprintf("%v", n.NodeListID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", n.Type))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(n.Name)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(n.Version)
	builder.WriteString(", ")
	builder.WriteString("file_name=")
	builder.WriteString(n.FileName)
	builder.WriteString(", ")
	builder.WriteString("url_home=")
	builder.WriteString(n.URLHome)
	builder.WriteString(", ")
	builder.WriteString("url_download=")
	builder.WriteString(n.URLDownload)
	builder.WriteString(", ")
	builder.WriteString("licenses=")
	builder.WriteString(fmt.Sprintf("%v", n.Licenses))
	builder.WriteString(", ")
	builder.WriteString("license_concluded=")
	builder.WriteString(n.LicenseConcluded)
	builder.WriteString(", ")
	builder.WriteString("license_comments=")
	builder.WriteString(n.LicenseComments)
	builder.WriteString(", ")
	builder.WriteString("copyright=")
	builder.WriteString(n.Copyright)
	builder.WriteString(", ")
	builder.WriteString("source_info=")
	builder.WriteString(n.SourceInfo)
	builder.WriteString(", ")
	builder.WriteString("comment=")
	builder.WriteString(n.Comment)
	builder.WriteString(", ")
	builder.WriteString("summary=")
	builder.WriteString(n.Summary)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(n.Description)
	builder.WriteString(", ")
	builder.WriteString("release_date=")
	builder.WriteString(n.ReleaseDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("build_date=")
	builder.WriteString(n.BuildDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("valid_until_date=")
	builder.WriteString(n.ValidUntilDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("attribution=")
	builder.WriteString(fmt.Sprintf("%v", n.Attribution))
	builder.WriteString(", ")
	builder.WriteString("file_types=")
	builder.WriteString(fmt.Sprintf("%v", n.FileTypes))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (n *Node) MarshalJSON() ([]byte, error) {
	type Alias Node
	return json.Marshal(&struct {
		*Alias
		NodeEdges
	}{
		Alias:     (*Alias)(n),
		NodeEdges: n.Edges,
	})
}

// Nodes is a parsable slice of Node.
type Nodes []*Node
