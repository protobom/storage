// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package gcs_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/protobom/protobom/pkg/reader"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	gcsbackend "github.com/protobom/storage/backends/gcs"
)

const testBucket = "protobom-test"

// newBackend returns an initialized GCS backend pointed at the emulator, skipping
// the test when STORAGE_EMULATOR_HOST is not set.
func newBackend(t *testing.T) *gcsbackend.Backend {
	t.Helper()

	if os.Getenv("STORAGE_EMULATOR_HOST") == "" {
		t.Skip("STORAGE_EMULATOR_HOST not set; skipping GCS emulator tests")
	}

	backend := gcsbackend.NewBackend(
		gcsbackend.WithBucket(testBucket),
		gcsbackend.WithProjectID("protobom-test"),
		gcsbackend.WithCreateBucket(true),
	)

	if err := backend.InitClient(); err != nil {
		t.Fatalf("initializing gcs backend: %v", err)
	}

	t.Cleanup(func() {
		if err := backend.CloseClient(); err != nil {
			t.Errorf("closing backend: %v", err)
		}
	})

	return backend
}

// loadDocument parses a testdata SBOM into a protobom document.
func loadDocument(t *testing.T, name string) *sbom.Document {
	t.Helper()

	cwd, err := os.Getwd()
	require.NoError(t, err)

	doc, err := reader.New().ParseFile(filepath.Join(cwd, "testdata", name))
	require.NoError(t, err)

	return doc
}

// TestBackendRoundTrip stores each sample document and retrieves it back,
// asserting the retrieved document is a lossless copy of the original.
func TestBackendRoundTrip(t *testing.T) {
	t.Parallel()

	backend := newBackend(t)

	for _, name := range []string{"sbom.cdx.json", "sbom.spdx.json"} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			doc := loadDocument(t, name)
			id := doc.GetMetadata().GetId()
			require.NotEmpty(t, id, "sample document must have a metadata id")

			require.NoError(t, backend.Store(doc, nil))

			got, err := backend.Retrieve(id, nil)
			require.NoError(t, err)
			require.True(t, proto.Equal(doc, got), "retrieved document must equal stored")
		})
	}
}

// TestBackendNoClobber verifies that storing a different document with an
// existing id and NoClobber set leaves the originally stored document intact.
func TestBackendNoClobber(t *testing.T) {
	t.Parallel()

	backend := newBackend(t)

	const id = "urn:protobom:test:noclobber"

	original := loadDocument(t, "sbom.cdx.json")
	original.GetMetadata().Id = id
	original.GetMetadata().Name = "original"
	require.NoError(t, backend.Store(original, nil))

	clobber := loadDocument(t, "sbom.cdx.json")
	clobber.GetMetadata().Id = id
	clobber.GetMetadata().Name = "clobbered"
	require.NoError(t, backend.Store(clobber, &storage.StoreOptions{NoClobber: true}))

	got, err := backend.Retrieve(id, nil)
	require.NoError(t, err)
	require.Equal(t, "original", got.GetMetadata().GetName())
}

// TestBackendIndexQueries stores a document and finds it back through the
// identifier, hash and name secondary indexes.
func TestBackendIndexQueries(t *testing.T) {
	t.Parallel()

	backend := newBackend(t)
	ctx := context.Background()

	doc := loadDocument(t, "sbom.cdx.json")
	id := doc.GetMetadata().GetId()
	require.NoError(t, backend.Store(doc, nil))

	identifier := firstIdentifier(doc)
	require.NotEmpty(t, identifier, "sample document must have an identifier")

	byIdentifier, err := backend.FindDocumentIDsByIdentifier(ctx, identifier)
	require.NoError(t, err)
	require.Contains(t, byIdentifier, id)

	algorithm, hashValue := firstHash(doc)
	require.NotEmpty(t, hashValue, "sample document must have a hash")

	byHash, err := backend.FindDocumentIDsByHash(ctx, algorithm, hashValue)
	require.NoError(t, err)
	require.Contains(t, byHash, id)

	name := firstName(doc)
	require.NotEmpty(t, name, "sample document must have a node name")

	byName, err := backend.FindDocumentIDsByName(ctx, strings.ToUpper(name))
	require.NoError(t, err)
	require.Contains(t, byName, id, "name lookup must be case-insensitive")

	absent, err := backend.FindDocumentIDsByName(ctx, "does-not-exist-"+id)
	require.NoError(t, err)
	require.Empty(t, absent, "an unindexed value must return no documents")
}

func firstIdentifier(doc *sbom.Document) string {
	for _, node := range doc.GetNodeList().GetNodes() {
		for _, value := range node.GetIdentifiers() {
			if value != "" {
				return value
			}
		}
	}

	return ""
}

func firstHash(doc *sbom.Document) (algorithm sbom.HashAlgorithm, value string) {
	for _, node := range doc.GetNodeList().GetNodes() {
		for algo, hashValue := range node.GetHashes() {
			if hashValue != "" {
				return sbom.HashAlgorithm(algo), hashValue
			}
		}
	}

	return sbom.HashAlgorithm_UNKNOWN, ""
}

func firstName(doc *sbom.Document) string {
	for _, node := range doc.GetNodeList().GetNodes() {
		if name := node.GetName(); name != "" {
			return name
		}
	}

	return ""
}
