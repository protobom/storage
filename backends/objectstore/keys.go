// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package objectstore

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"path"

	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
	"google.golang.org/protobuf/proto"
)

const (
	// documentsDir is the key prefix under which whole documents are stored.
	documentsDir = "documents"

	// indexDir is the key prefix under which secondary-index markers are stored.
	indexDir = "index"
)

// encodeSegment encodes an arbitrary string (document id, index value, ...) into
// a single URL-safe object-key path segment.
func encodeSegment(value string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(value))
}

// decodeSegment reverses encodeSegment.
func decodeSegment(segment string) (string, error) {
	data, err := base64.RawURLEncoding.DecodeString(segment)
	if err != nil {
		return "", fmt.Errorf("decoding key segment %q: %w", segment, err)
	}

	return string(data), nil
}

// documentKey is the object key holding the serialized document with the given id.
func documentKey(prefix, id string) string {
	return path.Join(prefix, documentsDir, encodeSegment(id))
}

// indexEntryKey is the marker-object key for "document docID has value under
// dimension".
func indexEntryKey(prefix string, dimension Dimension, value, docID string) string {
	return path.Join(prefix, indexDir, string(dimension), encodeSegment(value), encodeSegment(docID))
}

// indexValuePrefix is the key prefix listing every document indexed under
// dimension=value. The trailing slash prevents matching values that merely share
// an encoded prefix.
func indexValuePrefix(prefix string, dimension Dimension, value string) string {
	return path.Join(prefix, indexDir, string(dimension), encodeSegment(value)) + "/"
}

// documentID returns a document's storage id: its Metadata.id, or a deterministic
// content hash when the metadata carries no id.
func documentID(doc *sbom.Document) (string, error) {
	if id := doc.GetMetadata().GetId(); id != "" {
		return id, nil
	}

	data, err := proto.MarshalOptions{Deterministic: true}.Marshal(doc)
	if err != nil {
		return "", fmt.Errorf("marshaling document for id: %w", err)
	}

	return uuid.NewHash(sha256.New(), uuid.Max, data, int(uuid.Max.Version())).String(), nil
}
