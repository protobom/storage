// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package objectstore

import (
	"context"
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/protobom/protobom/pkg/sbom"
)

// Dimension identifies a queryable attribute the Indexer maintains a secondary
// index for. Its value is used verbatim as a path segment in index keys.
type Dimension string

const (
	// DimensionIdentifier indexes Node.Identifiers values (purl, cpe, gitoid).
	DimensionIdentifier Dimension = "ident"

	// DimensionHash indexes Node.Hashes as "<algorithm-code>:<value>".
	DimensionHash Dimension = "hash"

	// DimensionName indexes the lower-cased Node.Name.
	DimensionName Dimension = "name"
)

// Indexer maintains document-granularity secondary indexes in an ObjectStore.
// Each entry is a marker object keyed
// "<prefix>/index/<dimension>/<value>/<document-id>", so a lookup is a
// list-by-prefix over "<prefix>/index/<dimension>/<value>/". It is reusable by
// any ObjectStore-backed backend.
type Indexer struct {
	store  ObjectStore
	prefix string
}

// NewIndexer returns an Indexer that writes and queries markers through store
// under prefix.
func NewIndexer(store ObjectStore, prefix string) *Indexer {
	return &Indexer{store: store, prefix: prefix}
}

// Index writes the secondary-index markers for a document. It is idempotent:
// re-indexing the same document rewrites the same (empty) marker objects.
func (indexer *Indexer) Index(ctx context.Context, docID string, doc *sbom.Document) error {
	for entry := range collectIndexEntries(doc) {
		key := indexEntryKey(indexer.prefix, entry.dimension, entry.value, docID)
		if err := indexer.store.Put(ctx, key, nil); err != nil {
			return fmt.Errorf("writing index entry %s=%q: %w", entry.dimension, entry.value, err)
		}
	}

	return nil
}

// FindDocumentIDs returns the ids of documents indexed under dimension=value.
func (indexer *Indexer) FindDocumentIDs(ctx context.Context, dimension Dimension, value string) ([]string, error) {
	keys, err := indexer.store.List(ctx, indexValuePrefix(indexer.prefix, dimension, value))
	if err != nil {
		return nil, fmt.Errorf("listing index %s=%q: %w", dimension, value, err)
	}

	ids := make([]string, 0, len(keys))

	for _, key := range keys {
		docID, decodeErr := decodeSegment(path.Base(key))
		if decodeErr != nil {
			return nil, fmt.Errorf("decoding index entry %q: %w", key, decodeErr)
		}

		ids = append(ids, docID)
	}

	return ids, nil
}

// indexEntry is a distinct (dimension, value) pair to index for a document.
type indexEntry struct {
	dimension Dimension
	value     string
}

// collectIndexEntries returns the distinct index entries across every node in doc.
func collectIndexEntries(doc *sbom.Document) map[indexEntry]struct{} {
	entries := map[indexEntry]struct{}{}

	for _, node := range doc.GetNodeList().GetNodes() {
		collectNodeEntries(entries, node)
	}

	return entries
}

func collectNodeEntries(entries map[indexEntry]struct{}, node *sbom.Node) {
	for _, value := range node.GetIdentifiers() {
		addEntry(entries, DimensionIdentifier, value)
	}

	for algorithm, value := range node.GetHashes() {
		addEntry(entries, DimensionHash, hashIndexValue(algorithm, value))
	}

	addEntry(entries, DimensionName, strings.ToLower(node.GetName()))
}

func addEntry(entries map[indexEntry]struct{}, dimension Dimension, value string) {
	if value == "" {
		return
	}

	entries[indexEntry{dimension: dimension, value: value}] = struct{}{}
}

// hashIndexValue is the indexed value for a hash: "<algorithm-code>:<value>".
func hashIndexValue(algorithm int32, value string) string {
	return strconv.Itoa(int(algorithm)) + ":" + value
}
