// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package node

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the node type in the database.
	Label = "node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDocumentID holds the string denoting the document_id field in the database.
	FieldDocumentID = "document_id"
	// FieldProtoMessage holds the string denoting the proto_message field in the database.
	FieldProtoMessage = "proto_message"
	// FieldNativeID holds the string denoting the native_id field in the database.
	FieldNativeID = "native_id"
	// FieldNodeListID holds the string denoting the node_list_id field in the database.
	FieldNodeListID = "node_list_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldFileName holds the string denoting the file_name field in the database.
	FieldFileName = "file_name"
	// FieldURLHome holds the string denoting the url_home field in the database.
	FieldURLHome = "url_home"
	// FieldURLDownload holds the string denoting the url_download field in the database.
	FieldURLDownload = "url_download"
	// FieldLicenses holds the string denoting the licenses field in the database.
	FieldLicenses = "licenses"
	// FieldLicenseConcluded holds the string denoting the license_concluded field in the database.
	FieldLicenseConcluded = "license_concluded"
	// FieldLicenseComments holds the string denoting the license_comments field in the database.
	FieldLicenseComments = "license_comments"
	// FieldCopyright holds the string denoting the copyright field in the database.
	FieldCopyright = "copyright"
	// FieldSourceInfo holds the string denoting the source_info field in the database.
	FieldSourceInfo = "source_info"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldSummary holds the string denoting the summary field in the database.
	FieldSummary = "summary"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldReleaseDate holds the string denoting the release_date field in the database.
	FieldReleaseDate = "release_date"
	// FieldBuildDate holds the string denoting the build_date field in the database.
	FieldBuildDate = "build_date"
	// FieldValidUntilDate holds the string denoting the valid_until_date field in the database.
	FieldValidUntilDate = "valid_until_date"
	// FieldAttribution holds the string denoting the attribution field in the database.
	FieldAttribution = "attribution"
	// FieldFileTypes holds the string denoting the file_types field in the database.
	FieldFileTypes = "file_types"
	// EdgeDocument holds the string denoting the document edge name in mutations.
	EdgeDocument = "document"
	// EdgeAnnotations holds the string denoting the annotations edge name in mutations.
	EdgeAnnotations = "annotations"
	// EdgeSuppliers holds the string denoting the suppliers edge name in mutations.
	EdgeSuppliers = "suppliers"
	// EdgeOriginators holds the string denoting the originators edge name in mutations.
	EdgeOriginators = "originators"
	// EdgeExternalReferences holds the string denoting the external_references edge name in mutations.
	EdgeExternalReferences = "external_references"
	// EdgePrimaryPurpose holds the string denoting the primary_purpose edge name in mutations.
	EdgePrimaryPurpose = "primary_purpose"
	// EdgeToNodes holds the string denoting the to_nodes edge name in mutations.
	EdgeToNodes = "to_nodes"
	// EdgeNodes holds the string denoting the nodes edge name in mutations.
	EdgeNodes = "nodes"
	// EdgeHashes holds the string denoting the hashes edge name in mutations.
	EdgeHashes = "hashes"
	// EdgeIdentifiers holds the string denoting the identifiers edge name in mutations.
	EdgeIdentifiers = "identifiers"
	// EdgeProperties holds the string denoting the properties edge name in mutations.
	EdgeProperties = "properties"
	// EdgeNodeLists holds the string denoting the node_lists edge name in mutations.
	EdgeNodeLists = "node_lists"
	// EdgeEdgeTypes holds the string denoting the edge_types edge name in mutations.
	EdgeEdgeTypes = "edge_types"
	// Table holds the table name of the node in the database.
	Table = "nodes"
	// DocumentTable is the table that holds the document relation/edge.
	DocumentTable = "nodes"
	// DocumentInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	DocumentInverseTable = "documents"
	// DocumentColumn is the table column denoting the document relation/edge.
	DocumentColumn = "document_id"
	// AnnotationsTable is the table that holds the annotations relation/edge.
	AnnotationsTable = "annotations"
	// AnnotationsInverseTable is the table name for the Annotation entity.
	// It exists in this package in order to avoid circular dependency with the "annotation" package.
	AnnotationsInverseTable = "annotations"
	// AnnotationsColumn is the table column denoting the annotations relation/edge.
	AnnotationsColumn = "node_id"
	// SuppliersTable is the table that holds the suppliers relation/edge.
	SuppliersTable = "persons"
	// SuppliersInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	SuppliersInverseTable = "persons"
	// SuppliersColumn is the table column denoting the suppliers relation/edge.
	SuppliersColumn = "node_suppliers"
	// OriginatorsTable is the table that holds the originators relation/edge.
	OriginatorsTable = "persons"
	// OriginatorsInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	OriginatorsInverseTable = "persons"
	// OriginatorsColumn is the table column denoting the originators relation/edge.
	OriginatorsColumn = "node_id"
	// ExternalReferencesTable is the table that holds the external_references relation/edge. The primary key declared below.
	ExternalReferencesTable = "node_external_references"
	// ExternalReferencesInverseTable is the table name for the ExternalReference entity.
	// It exists in this package in order to avoid circular dependency with the "externalreference" package.
	ExternalReferencesInverseTable = "external_references"
	// PrimaryPurposeTable is the table that holds the primary_purpose relation/edge.
	PrimaryPurposeTable = "purposes"
	// PrimaryPurposeInverseTable is the table name for the Purpose entity.
	// It exists in this package in order to avoid circular dependency with the "purpose" package.
	PrimaryPurposeInverseTable = "purposes"
	// PrimaryPurposeColumn is the table column denoting the primary_purpose relation/edge.
	PrimaryPurposeColumn = "node_id"
	// ToNodesTable is the table that holds the to_nodes relation/edge. The primary key declared below.
	ToNodesTable = "edge_types"
	// NodesTable is the table that holds the nodes relation/edge. The primary key declared below.
	NodesTable = "edge_types"
	// HashesTable is the table that holds the hashes relation/edge. The primary key declared below.
	HashesTable = "node_hashes"
	// HashesInverseTable is the table name for the HashesEntry entity.
	// It exists in this package in order to avoid circular dependency with the "hashesentry" package.
	HashesInverseTable = "hashes_entries"
	// IdentifiersTable is the table that holds the identifiers relation/edge. The primary key declared below.
	IdentifiersTable = "node_identifiers"
	// IdentifiersInverseTable is the table name for the IdentifiersEntry entity.
	// It exists in this package in order to avoid circular dependency with the "identifiersentry" package.
	IdentifiersInverseTable = "identifiers_entries"
	// PropertiesTable is the table that holds the properties relation/edge.
	PropertiesTable = "properties"
	// PropertiesInverseTable is the table name for the Property entity.
	// It exists in this package in order to avoid circular dependency with the "property" package.
	PropertiesInverseTable = "properties"
	// PropertiesColumn is the table column denoting the properties relation/edge.
	PropertiesColumn = "node_id"
	// NodeListsTable is the table that holds the node_lists relation/edge. The primary key declared below.
	NodeListsTable = "node_list_nodes"
	// NodeListsInverseTable is the table name for the NodeList entity.
	// It exists in this package in order to avoid circular dependency with the "nodelist" package.
	NodeListsInverseTable = "node_lists"
	// EdgeTypesTable is the table that holds the edge_types relation/edge.
	EdgeTypesTable = "edge_types"
	// EdgeTypesInverseTable is the table name for the EdgeType entity.
	// It exists in this package in order to avoid circular dependency with the "edgetype" package.
	EdgeTypesInverseTable = "edge_types"
	// EdgeTypesColumn is the table column denoting the edge_types relation/edge.
	EdgeTypesColumn = "to_node_id"
)

// Columns holds all SQL columns for node fields.
var Columns = []string{
	FieldID,
	FieldDocumentID,
	FieldProtoMessage,
	FieldNativeID,
	FieldNodeListID,
	FieldType,
	FieldName,
	FieldVersion,
	FieldFileName,
	FieldURLHome,
	FieldURLDownload,
	FieldLicenses,
	FieldLicenseConcluded,
	FieldLicenseComments,
	FieldCopyright,
	FieldSourceInfo,
	FieldComment,
	FieldSummary,
	FieldDescription,
	FieldReleaseDate,
	FieldBuildDate,
	FieldValidUntilDate,
	FieldAttribution,
	FieldFileTypes,
}

var (
	// ExternalReferencesPrimaryKey and ExternalReferencesColumn2 are the table columns denoting the
	// primary key for the external_references relation (M2M).
	ExternalReferencesPrimaryKey = []string{"node_id", "external_reference_id"}
	// ToNodesPrimaryKey and ToNodesColumn2 are the table columns denoting the
	// primary key for the to_nodes relation (M2M).
	ToNodesPrimaryKey = []string{"node_id", "to_node_id"}
	// NodesPrimaryKey and NodesColumn2 are the table columns denoting the
	// primary key for the nodes relation (M2M).
	NodesPrimaryKey = []string{"node_id", "to_node_id"}
	// HashesPrimaryKey and HashesColumn2 are the table columns denoting the
	// primary key for the hashes relation (M2M).
	HashesPrimaryKey = []string{"node_id", "hash_entry_id"}
	// IdentifiersPrimaryKey and IdentifiersColumn2 are the table columns denoting the
	// primary key for the identifiers relation (M2M).
	IdentifiersPrimaryKey = []string{"node_id", "identifier_entry_id"}
	// NodeListsPrimaryKey and NodeListsColumn2 are the table columns denoting the
	// primary key for the node_lists relation (M2M).
	NodeListsPrimaryKey = []string{"node_list_id", "node_id"}
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

var (
	// DefaultDocumentID holds the default value on creation for the "document_id" field.
	DefaultDocumentID func() uuid.UUID
	// NativeIDValidator is a validator for the "native_id" field. It is called by the builders before save.
	NativeIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypePACKAGE Type = "PACKAGE"
	TypeFILE    Type = "FILE"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypePACKAGE, TypeFILE:
		return nil
	default:
		return fmt.Errorf("node: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Node queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDocumentID orders the results by the document_id field.
func ByDocumentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDocumentID, opts...).ToFunc()
}

// ByNativeID orders the results by the native_id field.
func ByNativeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNativeID, opts...).ToFunc()
}

// ByNodeListID orders the results by the node_list_id field.
func ByNodeListID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNodeListID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByFileName orders the results by the file_name field.
func ByFileName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileName, opts...).ToFunc()
}

// ByURLHome orders the results by the url_home field.
func ByURLHome(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURLHome, opts...).ToFunc()
}

// ByURLDownload orders the results by the url_download field.
func ByURLDownload(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURLDownload, opts...).ToFunc()
}

// ByLicenseConcluded orders the results by the license_concluded field.
func ByLicenseConcluded(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLicenseConcluded, opts...).ToFunc()
}

// ByLicenseComments orders the results by the license_comments field.
func ByLicenseComments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLicenseComments, opts...).ToFunc()
}

// ByCopyright orders the results by the copyright field.
func ByCopyright(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCopyright, opts...).ToFunc()
}

// BySourceInfo orders the results by the source_info field.
func BySourceInfo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSourceInfo, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
}

// BySummary orders the results by the summary field.
func BySummary(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSummary, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByReleaseDate orders the results by the release_date field.
func ByReleaseDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReleaseDate, opts...).ToFunc()
}

// ByBuildDate orders the results by the build_date field.
func ByBuildDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBuildDate, opts...).ToFunc()
}

// ByValidUntilDate orders the results by the valid_until_date field.
func ByValidUntilDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValidUntilDate, opts...).ToFunc()
}

// ByDocumentField orders the results by document field.
func ByDocumentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDocumentStep(), sql.OrderByField(field, opts...))
	}
}

// ByAnnotationsCount orders the results by annotations count.
func ByAnnotationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAnnotationsStep(), opts...)
	}
}

// ByAnnotations orders the results by annotations terms.
func ByAnnotations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAnnotationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySuppliersCount orders the results by suppliers count.
func BySuppliersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSuppliersStep(), opts...)
	}
}

// BySuppliers orders the results by suppliers terms.
func BySuppliers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSuppliersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOriginatorsCount orders the results by originators count.
func ByOriginatorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOriginatorsStep(), opts...)
	}
}

// ByOriginators orders the results by originators terms.
func ByOriginators(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOriginatorsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByExternalReferencesCount orders the results by external_references count.
func ByExternalReferencesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExternalReferencesStep(), opts...)
	}
}

// ByExternalReferences orders the results by external_references terms.
func ByExternalReferences(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExternalReferencesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPrimaryPurposeCount orders the results by primary_purpose count.
func ByPrimaryPurposeCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPrimaryPurposeStep(), opts...)
	}
}

// ByPrimaryPurpose orders the results by primary_purpose terms.
func ByPrimaryPurpose(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPrimaryPurposeStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByToNodesCount orders the results by to_nodes count.
func ByToNodesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newToNodesStep(), opts...)
	}
}

// ByToNodes orders the results by to_nodes terms.
func ByToNodes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newToNodesStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByHashesCount orders the results by hashes count.
func ByHashesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHashesStep(), opts...)
	}
}

// ByHashes orders the results by hashes terms.
func ByHashes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHashesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByIdentifiersCount orders the results by identifiers count.
func ByIdentifiersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIdentifiersStep(), opts...)
	}
}

// ByIdentifiers orders the results by identifiers terms.
func ByIdentifiers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIdentifiersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPropertiesCount orders the results by properties count.
func ByPropertiesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPropertiesStep(), opts...)
	}
}

// ByProperties orders the results by properties terms.
func ByProperties(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPropertiesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNodeListsCount orders the results by node_lists count.
func ByNodeListsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNodeListsStep(), opts...)
	}
}

// ByNodeLists orders the results by node_lists terms.
func ByNodeLists(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodeListsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
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
func newDocumentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DocumentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DocumentTable, DocumentColumn),
	)
}
func newAnnotationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AnnotationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AnnotationsTable, AnnotationsColumn),
	)
}
func newSuppliersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SuppliersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SuppliersTable, SuppliersColumn),
	)
}
func newOriginatorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OriginatorsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OriginatorsTable, OriginatorsColumn),
	)
}
func newExternalReferencesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExternalReferencesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ExternalReferencesTable, ExternalReferencesPrimaryKey...),
	)
}
func newPrimaryPurposeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PrimaryPurposeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PrimaryPurposeTable, PrimaryPurposeColumn),
	)
}
func newToNodesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ToNodesTable, ToNodesPrimaryKey...),
	)
}
func newNodesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, NodesTable, NodesPrimaryKey...),
	)
}
func newHashesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HashesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, HashesTable, HashesPrimaryKey...),
	)
}
func newIdentifiersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IdentifiersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, IdentifiersTable, IdentifiersPrimaryKey...),
	)
}
func newPropertiesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PropertiesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PropertiesTable, PropertiesColumn),
	)
}
func newNodeListsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodeListsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, NodeListsTable, NodeListsPrimaryKey...),
	)
}
func newEdgeTypesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EdgeTypesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, EdgeTypesTable, EdgeTypesColumn),
	)
}
