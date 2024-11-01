// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/protobom/protobom/pkg/reader"

	"github.com/protobom/storage/backends/ent"
)

func Example() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dbFile := filepath.Join(cwd, "example.db")

	// Remove example.db if it already exists.
	if err := os.Remove(dbFile); err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}

	rdr := reader.New()
	backend := ent.NewBackend().WithDatabaseFile(dbFile)

	if err := backend.InitClient(); err != nil {
		panic(err)
	}

	defer backend.CloseClient()

	sbom, err := rdr.ParseFile(filepath.Join(cwd, "testdata", "sbom.cdx.json"))
	if err != nil {
		panic(err)
	}

	if err := backend.Store(sbom, nil); err != nil {
		panic(err)
	}

	retrieved, err := backend.Retrieve(sbom.GetMetadata().GetId(), nil)
	if err != nil {
		panic(err)
	}

	// Remove source data URI to allow comparison.
	retrieved.GetMetadata().GetSourceData().Uri = nil

	output, err := json.MarshalIndent(retrieved, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
	//nolint:lll
	// Output:
	// {
	//   "metadata": {
	//     "id": "urn:uuid:3e671687-395b-41f5-a30f-a58921a69b79",
	//     "version": "1",
	//     "date": {},
	//     "source_data": {
	//       "format": "application/vnd.cyclonedx+json;version=1.5",
	//       "hashes": {
	//         "3": "71a3948e45c0bcd83a617ed94674079778d10a0578932e6e536533339b1bbea5"
	//       },
	//       "size": 5263
	//     }
	//   },
	//   "node_list": {
	//     "nodes": [
	//       {
	//         "id": "protobom-auto--000000001",
	//         "name": "Acme Application",
	//         "version": "9.1.1",
	//         "primary_purpose": [
	//           1
	//         ]
	//       },
	//       {
	//         "id": "pkg:npm/acme/component@1.0.0",
	//         "name": "tomcat-catalina",
	//         "version": "9.0.14",
	//         "licenses": [
	//           "Apache-2.0"
	//         ],
	//         "license_concluded": "Apache-2.0",
	//         "identifiers": {
	//           "1": "pkg:npm/acme/component@1.0.0"
	//         },
	//         "hashes": {
	//           "1": "3942447fac867ae5cdb3229b658f4d48",
	//           "2": "e6b1000b94e835ffd37f4c6dcbdad43f4b48a02a",
	//           "3": "f498a8ff2dd007e29c2074f5e4b01a9a01775c3ff3aeaf6906ea503bc5791b7b",
	//           "5": "e8f33e424f3f4ed6db76a482fde1a5298970e442c531729119e37991884bdffab4f9426b7ee11fccd074eeda0634d71697d6f88a460dce0ac8d627a29f7d1282"
	//         },
	//         "primary_purpose": [
	//           16
	//         ]
	//       },
	//       {
	//         "id": "protobom-auto--000000003",
	//         "name": "mylibrary",
	//         "version": "1.0.0",
	//         "primary_purpose": [
	//           16
	//         ]
	//       }
	//     ],
	//     "edges": [
	//       {
	//         "type": 5,
	//         "from": "protobom-auto--000000001",
	//         "to": [
	//           "pkg:npm/acme/component@1.0.0",
	//           "protobom-auto--000000003"
	//         ]
	//       }
	//     ],
	//     "root_elements": [
	//       "protobom-auto--000000001"
	//     ]
	//   }
	// }
}
