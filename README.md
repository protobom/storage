# protobom storage

[![Go Reference](https://pkg.go.dev/badge/github.com/protobom/storage.svg)](https://pkg.go.dev/github.com/protobom/storage)

Storage backends for [protobom](https://github.com/protobom/protobom), the
format-neutral SBOM data model. Each backend implements the
[`storage.Backend`](https://pkg.go.dev/github.com/protobom/protobom/pkg/storage#Backend)
interface from protobom, so a `sbom.Document` parsed from SPDX or CycloneDX can
be persisted and retrieved from a database or object store with the same
`Store` / `Retrieve` calls.

## Backends

| Package                 | Storage                     | Notes                                                                                       |
| ----------------------- | --------------------------- | ------------------------------------------------------------------------------------------- |
| `backends/ent`          | SQLite via [ent](https://entgo.io) | Fully relational schema, in-memory or file-backed. Supports document and node annotations. |
| `backends/clickhouse`   | [ClickHouse](https://clickhouse.com) | Stores the serialized document plus wide node and edge tables for analytical queries. |
| `backends/gcs`          | Google Cloud Storage        | Object-store backend with a key index for lookups by identifier, hash and name.            |
| `backends/objectstore`  | Any object store            | Provider-agnostic core used by `gcs`. Implement the `ObjectStore` interface to add S3, MinIO, etc. |

## Installation

```console
go get github.com/protobom/storage
```

## Quick start

The example below parses a CycloneDX SBOM, stores it in an in-memory SQLite
database using the `ent` backend, and reads it back by document ID.

```go
package main

import (
	"fmt"

	"github.com/protobom/protobom/pkg/reader"

	"github.com/protobom/storage/backends/ent"
)

func main() {
	// Parse an SBOM (SPDX or CycloneDX) into a protobom document.
	doc, err := reader.New().ParseFile("sbom.cdx.json")
	if err != nil {
		panic(err)
	}

	// Create the backend. The default database is ":memory:";
	// use ent.WithDatabaseFile("sboms.db") to persist to disk.
	backend := ent.NewBackend()
	if err := backend.InitClient(); err != nil {
		panic(err)
	}
	defer backend.CloseClient()

	// Store the document.
	if err := backend.Store(doc, nil); err != nil {
		panic(err)
	}

	// Retrieve it again by its ID.
	stored, err := backend.Retrieve(doc.GetMetadata().GetId(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s has %d nodes\n", stored.GetMetadata().GetId(), len(stored.GetNodeList().GetNodes()))
}
```

Switching backends only changes the constructor. For example, to use ClickHouse:

```go
backend := clickhouse.NewBackend(
	clickhouse.WithAddr("localhost:9000"),
	clickhouse.WithCredentials("default", ""),
)
```

or Google Cloud Storage:

```go
backend := gcs.NewBackend(gcs.WithBucket("my-sboms"))
```

See the `Example` test in [`backends/ent/example_test.go`](backends/ent/example_test.go)
for a complete, runnable round trip and the per-package documentation on
[pkg.go.dev](https://pkg.go.dev/github.com/protobom/storage) for all options.

## Development

```console
make help             # list all targets
make lint             # run golangci-lint
make test-unit        # ent backend unit tests (no external services)
make test-clickhouse  # needs a ClickHouse server; see CLICKHOUSE_ADDR/USER/PASSWORD
make test-gcs         # needs STORAGE_EMULATOR_HOST or GCS credentials
```

## License

Apache-2.0. See [LICENSE](LICENSE).
