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

// documentsDir is the key prefix under which whole documents are stored.
const documentsDir = "documents"

// encodeSegment encodes an arbitrary string (document id, index value, ...) into
// a single URL-safe object-key path segment.
func encodeSegment(value string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(value))
}

// documentKey is the object key holding the serialized document with the given id.
func documentKey(prefix, id string) string {
	return path.Join(prefix, documentsDir, encodeSegment(id))
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
