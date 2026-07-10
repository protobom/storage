// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package clickhouse

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"
	"google.golang.org/protobuf/proto"
)

// Insert statements for the backend's tables. clickhouse-go infers the column
// list from the table and AppendStruct maps struct fields to columns by ch tag.
const (
	insertDocuments = "INSERT INTO documents"
	insertNodes     = "INSERT INTO nodes"
	insertEdges     = "INSERT INTO edges"
)

// documentRow mirrors the columns of the documents table.
type documentRow struct {
	CreatedDate   time.Time        `ch:"created_date"`
	StoredAt      time.Time        `ch:"stored_at"`
	SourceHashes  map[int16]string `ch:"source_hashes"`
	ID            string           `ch:"id"`
	Version       string           `ch:"version"`
	Name          string           `ch:"name"`
	Comment       string           `ch:"comment"`
	SourceFormat  string           `ch:"source_format"`
	SourceURI     string           `ch:"source_uri"`
	Document      []byte           `ch:"document"`
	DocumentTypes []docTypeTuple   `ch:"document_types"`
	Tools         []toolTuple      `ch:"tools"`
	Authors       []personTuple    `ch:"authors"`
	RootElements  []string         `ch:"root_elements"`
	SourceSize    int64            `ch:"source_size"`
}

// nodeRow mirrors the columns of the nodes table.
type nodeRow struct {
	ReleaseDate        time.Time        `ch:"release_date"`
	BuildDate          time.Time        `ch:"build_date"`
	ValidUntilDate     time.Time        `ch:"valid_until_date"`
	StoredAt           time.Time        `ch:"stored_at"`
	Identifiers        map[int16]string `ch:"identifiers"`
	Hashes             map[int16]string `ch:"hashes"`
	DocumentID         string           `ch:"document_id"`
	ID                 string           `ch:"id"`
	Name               string           `ch:"name"`
	Version            string           `ch:"version"`
	FileName           string           `ch:"file_name"`
	URLHome            string           `ch:"url_home"`
	URLDownload        string           `ch:"url_download"`
	LicenseConcluded   string           `ch:"license_concluded"`
	LicenseComments    string           `ch:"license_comments"`
	Copyright          string           `ch:"copyright"`
	SourceInfo         string           `ch:"source_info"`
	Comment            string           `ch:"comment"`
	Summary            string           `ch:"summary"`
	Description        string           `ch:"description"`
	Licenses           []string         `ch:"licenses"`
	Attribution        []string         `ch:"attribution"`
	FileTypes          []string         `ch:"file_types"`
	Purposes           []int16          `ch:"purposes"`
	Properties         []propertyTuple  `ch:"properties"`
	ExternalReferences []extRefTuple    `ch:"external_references"`
	Suppliers          []personTuple    `ch:"suppliers"`
	Originators        []personTuple    `ch:"originators"`
	Type               int16            `ch:"type"`
}

// edgeRow mirrors the columns of the edges table.
type edgeRow struct {
	StoredAt   time.Time `ch:"stored_at"`
	DocumentID string    `ch:"document_id"`
	FromNodeID string    `ch:"from_node_id"`
	ToNodeIDs  []string  `ch:"to_node_ids"`
	Type       int16     `ch:"type"`
}

// docTypeTuple is one element of documents.document_types.
type docTypeTuple struct {
	Name        string `ch:"name"`
	Description string `ch:"description"`
	Type        int16  `ch:"type"`
}

// toolTuple is one element of documents.tools.
type toolTuple struct {
	Name    string `ch:"name"`
	Version string `ch:"version"`
	Vendor  string `ch:"vendor"`
}

// personTuple is one element of the person tuple arrays (authors, suppliers,
// originators). Person.contacts is intentionally not modeled here; full
// fidelity is preserved in the document blob.
type personTuple struct {
	Name  string `ch:"name"`
	Email string `ch:"email"`
	URL   string `ch:"url"`
	Phone string `ch:"phone"`
	IsOrg bool   `ch:"is_org"`
}

// propertyTuple is one element of nodes.properties.
type propertyTuple struct {
	Name string `ch:"name"`
	Data string `ch:"data"`
}

// extRefTuple is one element of nodes.external_references.
type extRefTuple struct {
	Hashes    map[int16]string `ch:"hashes"`
	URL       string           `ch:"url"`
	Comment   string           `ch:"comment"`
	Authority string           `ch:"authority"`
	Type      int16            `ch:"type"`
}

// Store implements the storage.Storer interface. It keeps the full serialized
// document as a blob for lossless retrieval and, in parallel, decomposes the
// node list into the wide nodes and edges tables for querying.
//
// ClickHouse has no cross-table transactions; correctness on re-store comes from
// idempotency: the tables are ReplacingMergeTree engines keyed by content, so
// storing the same document twice converges rather than duplicating rows.
func (backend *Backend) Store(doc *sbom.Document, opts *storage.StoreOptions) error {
	if backend.conn == nil {
		return errUninitializedClient
	}

	key, err := documentKey(doc)
	if err != nil {
		return err
	}

	skip, err := backend.shouldSkip(key, opts)
	if err != nil {
		return err
	}

	if skip {
		return nil
	}

	blob, err := proto.MarshalOptions{Deterministic: true}.Marshal(doc)
	if err != nil {
		return fmt.Errorf("marshaling document: %w", err)
	}

	now := time.Now()
	nodeList := doc.GetNodeList()

	if err := backend.storeDocument(key, blob, doc, now); err != nil {
		return err
	}

	if err := backend.storeNodes(key, nodeList.GetNodes(), now); err != nil {
		return err
	}

	return backend.storeEdges(key, nodeList.GetEdges(), now)
}

// shouldSkip reports whether the store should be skipped because NoClobber is
// set and a document with the same id already exists.
func (backend *Backend) shouldSkip(key string, opts *storage.StoreOptions) (bool, error) {
	if opts == nil || !opts.NoClobber {
		return false, nil
	}

	return backend.documentExists(key)
}

// documentExists reports whether a document with the given id is already stored.
func (backend *Backend) documentExists(id string) (bool, error) {
	const query = "SELECT count() FROM documents WHERE id = ?"

	var count uint64
	if err := backend.conn.QueryRow(backend.ctx, query, id).Scan(&count); err != nil {
		return false, fmt.Errorf("checking document existence: %w", err)
	}

	return count > 0, nil
}

func (backend *Backend) storeDocument(key string, blob []byte, doc *sbom.Document, now time.Time) error {
	rows := []documentRow{newDocumentRow(key, blob, doc, now)}

	return appendBatch(backend, insertDocuments, rows)
}

func (backend *Backend) storeNodes(documentID string, nodes []*sbom.Node, now time.Time) error {
	rows := make([]nodeRow, 0, len(nodes))
	for _, node := range nodes {
		rows = append(rows, newNodeRow(documentID, node, now))
	}

	return appendBatch(backend, insertNodes, rows)
}

func (backend *Backend) storeEdges(documentID string, edges []*sbom.Edge, now time.Time) error {
	rows := make([]edgeRow, 0, len(edges))
	for _, edge := range edges {
		rows = append(rows, newEdgeRow(documentID, edge, now))
	}

	return appendBatch(backend, insertEdges, rows)
}

// appendBatch prepares a batch for the given insert statement, appends every row
// and sends it. An empty slice is a no-op.
func appendBatch[T any](backend *Backend, query string, rows []T) error {
	if len(rows) == 0 {
		return nil
	}

	batch, err := backend.conn.PrepareBatch(backend.ctx, query)
	if err != nil {
		return fmt.Errorf("preparing batch %q: %w", query, err)
	}

	for idx := range rows {
		if err := batch.AppendStruct(&rows[idx]); err != nil {
			return fmt.Errorf("appending row to %q: %w", query, err)
		}
	}

	if err := batch.Send(); err != nil {
		return fmt.Errorf("sending batch %q: %w", query, err)
	}

	return nil
}

func newDocumentRow(key string, blob []byte, doc *sbom.Document, now time.Time) documentRow {
	metadata := doc.GetMetadata()
	source := metadata.GetSourceData()

	return documentRow{
		ID:            key,
		Document:      blob,
		Version:       metadata.GetVersion(),
		Name:          metadata.GetName(),
		CreatedDate:   metadata.GetDate().AsTime(),
		Comment:       metadata.GetComment(),
		DocumentTypes: docTypeTuples(metadata.GetDocumentTypes()),
		Tools:         toolTuples(metadata.GetTools()),
		Authors:       personTuples(metadata.GetAuthors()),
		SourceFormat:  source.GetFormat(),
		SourceSize:    source.GetSize(),
		SourceURI:     source.GetUri(),
		SourceHashes:  int16Map(source.GetHashes()),
		RootElements:  doc.GetNodeList().GetRootElements(),
		StoredAt:      now,
	}
}

func newNodeRow(documentID string, node *sbom.Node, now time.Time) nodeRow {
	return nodeRow{
		DocumentID:         documentID,
		ID:                 node.GetId(),
		Type:               toInt16(int32(node.GetType())),
		Name:               node.GetName(),
		Version:            node.GetVersion(),
		FileName:           node.GetFileName(),
		URLHome:            node.GetUrlHome(),
		URLDownload:        node.GetUrlDownload(),
		LicenseConcluded:   node.GetLicenseConcluded(),
		LicenseComments:    node.GetLicenseComments(),
		Copyright:          node.GetCopyright(),
		SourceInfo:         node.GetSourceInfo(),
		Comment:            node.GetComment(),
		Summary:            node.GetSummary(),
		Description:        node.GetDescription(),
		ReleaseDate:        node.GetReleaseDate().AsTime(),
		BuildDate:          node.GetBuildDate().AsTime(),
		ValidUntilDate:     node.GetValidUntilDate().AsTime(),
		Licenses:           node.GetLicenses(),
		Attribution:        node.GetAttribution(),
		FileTypes:          node.GetFileTypes(),
		Identifiers:        int16Map(node.GetIdentifiers()),
		Hashes:             int16Map(node.GetHashes()),
		Purposes:           purposeCodes(node.GetPrimaryPurpose()),
		Properties:         propertyTuples(node.GetProperties()),
		ExternalReferences: extRefTuples(node.GetExternalReferences()),
		Suppliers:          personTuples(node.GetSuppliers()),
		Originators:        personTuples(node.GetOriginators()),
		StoredAt:           now,
	}
}

func newEdgeRow(documentID string, edge *sbom.Edge, now time.Time) edgeRow {
	return edgeRow{
		DocumentID: documentID,
		Type:       toInt16(int32(edge.GetType())),
		FromNodeID: edge.GetFrom(),
		ToNodeIDs:  edge.GetTo(),
		StoredAt:   now,
	}
}

// documentKey returns the retrieval key for a document: its Metadata.id, or a
// deterministic content hash when the metadata carries no id.
func documentKey(doc *sbom.Document) (string, error) {
	if id := doc.GetMetadata().GetId(); id != "" {
		return id, nil
	}

	data, err := proto.MarshalOptions{Deterministic: true}.Marshal(doc)
	if err != nil {
		return "", fmt.Errorf("marshaling document for id: %w", err)
	}

	return uuid.NewHash(sha256.New(), uuid.Max, data, int(uuid.Max.Version())).String(), nil
}

func docTypeTuples(docTypes []*sbom.DocumentType) []docTypeTuple {
	tuples := make([]docTypeTuple, 0, len(docTypes))
	for _, docType := range docTypes {
		tuples = append(tuples, docTypeTuple{
			Type:        toInt16(int32(docType.GetType())),
			Name:        docType.GetName(),
			Description: docType.GetDescription(),
		})
	}

	return tuples
}

func toolTuples(tools []*sbom.Tool) []toolTuple {
	tuples := make([]toolTuple, 0, len(tools))
	for _, tool := range tools {
		tuples = append(tuples, toolTuple{
			Name:    tool.GetName(),
			Version: tool.GetVersion(),
			Vendor:  tool.GetVendor(),
		})
	}

	return tuples
}

func personTuples(persons []*sbom.Person) []personTuple {
	tuples := make([]personTuple, 0, len(persons))
	for _, person := range persons {
		tuples = append(tuples, personTuple{
			Name:  person.GetName(),
			Email: person.GetEmail(),
			URL:   person.GetUrl(),
			Phone: person.GetPhone(),
			IsOrg: person.GetIsOrg(),
		})
	}

	return tuples
}

func propertyTuples(properties []*sbom.Property) []propertyTuple {
	tuples := make([]propertyTuple, 0, len(properties))
	for _, property := range properties {
		tuples = append(tuples, propertyTuple{
			Name: property.GetName(),
			Data: property.GetData(),
		})
	}

	return tuples
}

func extRefTuples(refs []*sbom.ExternalReference) []extRefTuple {
	tuples := make([]extRefTuple, 0, len(refs))
	for _, ref := range refs {
		tuples = append(tuples, extRefTuple{
			URL:       ref.GetUrl(),
			Comment:   ref.GetComment(),
			Authority: ref.GetAuthority(),
			Type:      toInt16(int32(ref.GetType())),
			Hashes:    int16Map(ref.GetHashes()),
		})
	}

	return tuples
}

func purposeCodes(purposes []sbom.Purpose) []int16 {
	codes := make([]int16, 0, len(purposes))
	for _, purpose := range purposes {
		codes = append(codes, toInt16(int32(purpose)))
	}

	return codes
}

func int16Map(values map[int32]string) map[int16]string {
	result := make(map[int16]string, len(values))
	for key, value := range values {
		result[toInt16(key)] = value
	}

	return result
}

// toInt16 narrows a protobuf enum code or map key to the int16 used by the
// schema. All such values are small and well within range.
func toInt16(value int32) int16 {
	return int16(value) //nolint:gosec // protobuf enum codes and map keys fit in int16
}
