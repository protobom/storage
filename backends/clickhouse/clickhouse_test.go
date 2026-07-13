// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package clickhouse_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/protobom/protobom/pkg/reader"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/protobom/storage/backends/clickhouse"
)

const testDatabase = "protobom_test"

// testAddr returns the ClickHouse address to test against, honoring the
// CLICKHOUSE_ADDR environment variable and defaulting to a local server.
func testAddr() string {
	if addr := os.Getenv("CLICKHOUSE_ADDR"); addr != "" {
		return addr
	}

	return "localhost:9000"
}

// newBackend returns an initialized backend, skipping the test when no
// ClickHouse server is reachable.
func newBackend(t *testing.T) *clickhouse.Backend {
	t.Helper()

	opts := []clickhouse.Option{
		clickhouse.WithAddr(testAddr()),
		clickhouse.WithDatabase(testDatabase),
	}

	if user := os.Getenv("CLICKHOUSE_USER"); user != "" {
		opts = append(opts, clickhouse.WithCredentials(user, os.Getenv("CLICKHOUSE_PASSWORD")))
	}

	backend := clickhouse.NewBackend(opts...)

	if err := backend.InitClient(); err != nil {
		t.Skipf("clickhouse not available at %s: %v", testAddr(), err)
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

// TestBackendQueryable proves the wide nodes table is populated and queryable
// with ClickHouse array/map functions after a store.
func TestBackendQueryable(t *testing.T) {
	t.Parallel()

	backend := newBackend(t)

	doc := loadDocument(t, "sbom.cdx.json")
	require.NoError(t, backend.Store(doc, nil))

	ctx := context.Background()
	id := doc.GetMetadata().GetId()

	const countQuery = "SELECT count() FROM nodes FINAL WHERE document_id = ?"

	var nodeCount uint64

	require.NoError(t, backend.Conn().QueryRow(ctx, countQuery, id).Scan(&nodeCount))
	require.EqualValues(t, len(doc.GetNodeList().GetNodes()), nodeCount)

	const licenseQuery = "SELECT count() FROM nodes FINAL " +
		"WHERE document_id = ? AND has(licenses, 'Apache-2.0')"

	var apacheCount uint64

	require.NoError(t, backend.Conn().QueryRow(ctx, licenseQuery, id).Scan(&apacheCount))
	require.Positive(t, apacheCount, "expected at least one Apache-2.0 licensed node")
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
