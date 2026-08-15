<!--
SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
SPDX-License-Identifier: Apache-2.0
-->

# Protobom SQL Schemas

Generic relational schemas for storing [protobom](https://github.com/protobom/protobom)
SBOM data in a SQL database. They are derived directly from the protobom data
model in [`protobom/api/sbom.proto`](https://github.com/protobom/protobom/blob/main/api/sbom.proto).

| File | Purpose |
| --- | --- |
| [`schema.ansi.sql`](./schema.ansi.sql) | Engine-neutral reference schema using portable types. Apply this (with the type mapping noted in its header) on engines such as MySQL/MariaDB, SQLite, SQL Server or Oracle. |
| [`schema.ansi-no-json.sql`](./schema.ansi-no-json.sql) | Same as `schema.ansi.sql` but with **no JSON columns**: the scalar string arrays are fully normalized into ordered child tables. Use on engines without a usable JSON type. |
| [`schema.postgres.sql`](./schema.postgres.sql) | PostgreSQL-tuned variant using native `JSONB`, `TIMESTAMPTZ`, identity columns, etc. |

Both files describe the __same logical model__**__ and the same table/column names,
only the physical types and a few syntactic details differ.

## Schema Notes

- **Enums are stored as integer codes:** Every enum-typed column holds the raw
  protobuf enum _number_ declared in `sbom.proto` (stored as `SMALLINT`), not a
  text label. This keeps values stable across language bindings and compact on
  disk. The complete mapping is in the [Enum reference](#enum-reference) below.

- **Collections use a hybrid model:**
  - Scalar string arrays — `Node.licenses`, `Node.attribution`,
    `Node.file_types`, and `NodeList.root_elements` — are stored inline as a
    JSON array column on the parent row.
  - Structured maps and lists (`hashes`, `identifiers`, `primary_purpose`,
    `properties`, `external_references`, `persons`, and graph `edges`) are
    normalized into their own child tables so they can be joined and indexed.

  `schema.ansi-no-json.sql` drops the hybrid model entirely: the scalar string
  arrays become ordered child tables (`node_licenses`, `node_attributions`,
  `node_file_types`, `node_list_roots`), each with an `ordinal` column that
  preserves the order of the protobuf repeated field.

- **Nodes are content-addressable and shareable:** A node is keyed by its
  protobom `id` and linked to a `node_lists` row through the
  `node_list_nodes` association table, so the same node may belong to more than
  one document without duplication.

- **Cascading deletes:** Foreign keys cascade from a document down through its
  metadata, node list, nodes and all sub-collections, so deleting a `documents`
  row removes the whole graph it owns.

## Entity overview

```text
documents ──1:1── metadata ──┬─< document_types
    │                        ├─< tools
    │                        ├─1:1─ source_data ──< source_data_hashes
    │                        └─< persons (authors)
    │
    └─1:1── node_lists ──┬─< node_list_nodes >── nodes
                         └─< edges ──< edge_targets

nodes ──┬─< node_identifiers
        ├─< node_hashes
        ├─< node_purposes
        ├─< node_properties
        ├─< external_references ──< external_reference_hashes
        └─< persons (suppliers / originators)

persons ──< persons (contacts, self-reference)
```

## Enum Reference

All values below are the protobuf enum numbers from `sbom.proto`. Store these
integers in corresponding `SMALLINT` columns.

### `Node.NodeType` > `nodes.type`

| Code | Name |
| --- | --- |
| 0 | PACKAGE |
| 1 | FILE |

### `Edge.Type` > `edges.type`

| Code | Name | Code | Name | Code | Name |
| --- | --- | --- | --- | --- | --- |
| 0 | UNKNOWN | 15 | devDependency | 30 | other |
| 1 | amends | 16 | devTool | 31 | packages |
| 2 | ancestor | 17 | distributionArtifact | 32 | patch |
| 3 | buildDependency | 18 | documentation | 33 | prerequisite |
| 4 | buildTool | 19 | dynamicLink | 34 | prerequisiteFor |
| 5 | contains | 20 | example | 35 | providedDependency |
| 6 | contained_by | 21 | expandedFromArchive | 36 | requirementFor |
| 7 | copy | 22 | fileAdded | 37 | runtimeDependency |
| 8 | dataFile | 23 | fileDeleted | 38 | specificationFor |
| 9 | dependencyManifest | 24 | fileModified | 39 | staticLink |
| 10 | dependsOn | 25 | generates | 40 | test |
| 11 | dependencyOf | 26 | generatedFrom | 41 | testCase |
| 12 | descendant | 27 | metafile | 42 | testDependency |
| 13 | describes | 28 | optionalComponent | 43 | testTool |
| 14 | describedBy | 29 | optionalDependency | 44 | variant |

### `DocumentType.SBOMType` > `document_types.type`

| Code | Name | Code | Name |
| --- | --- | --- | --- |
| 0 | OTHER | 5 | DEPLOYED |
| 1 | DESIGN | 6 | RUNTIME |
| 2 | SOURCE | 7 | DISCOVERY |
| 3 | BUILD | 8 | DECOMISSION |
| 4 | ANALYZED | | |

### `SoftwareIdentifierType` > `node_identifiers.type`

| Code | Name |
| --- | --- |
| 0 | UNKNOWN_IDENTIFIER_TYPE |
| 1 | PURL |
| 2 | CPE22 |
| 3 | CPE23 |
| 4 | GITOID |

### `HashAlgorithm` > `*_hashes.algorithm`

| Code | Name | Code | Name |
| --- | --- | --- | --- |
| 0 | UNKNOWN | 9 | BLAKE2B_256 |
| 1 | MD5 | 10 | BLAKE2B_384 |
| 2 | SHA1 | 11 | BLAKE2B_512 |
| 3 | SHA256 | 12 | BLAKE3 |
| 4 | SHA384 | 13 | MD2 |
| 5 | SHA512 | 14 | ADLER32 |
| 6 | SHA3_256 | 15 | MD4 |
| 7 | SHA3_384 | 16 | MD6 |
| 8 | SHA3_512 | 17 | SHA224 |

### `Purpose` > `node_purposes.purpose`

| Code | Name | Code | Name | Code | Name |
| --- | --- | --- | --- | --- | --- |
| 0 | UNKNOWN_PURPOSE | 10 | EVIDENCE | 20 | MODULE |
| 1 | APPLICATION | 11 | EXECUTABLE | 21 | OPERATING_SYSTEM |
| 2 | ARCHIVE | 12 | FILE | 22 | OTHER |
| 3 | BOM | 13 | FIRMWARE | 23 | PATCH |
| 4 | CONFIGURATION | 14 | FRAMEWORK | 24 | PLATFORM |
| 5 | CONTAINER | 15 | INSTALL | 25 | REQUIREMENT |
| 6 | DATA | 16 | LIBRARY | 26 | SOURCE |
| 7 | DEVICE | 17 | MACHINE_LEARNING_MODEL | 27 | SPECIFICATION |
| 8 | DEVICE_DRIVER | 18 | MANIFEST | 28 | TEST |
| 9 | DOCUMENTATION | 19 | MODEL | | |

### `ExternalReference.ExternalReferenceType` > `external_references.type`

| Code | Name | Code | Name | Code | Name |
| --- | --- | --- | --- | --- | --- |
| 0 | UNKNOWN | 21 | ISSUE_TRACKER | 42 | SECURE_SOFTWARE_ATTESTATION |
| 1 | ATTESTATION | 22 | LICENSE | 43 | SECURITY_ADVERSARY_MODEL |
| 2 | BINARY | 23 | LOG | 44 | SECURITY_ADVISORY |
| 3 | BOM | 24 | MAILING_LIST | 45 | SECURITY_CONTACT |
| 4 | BOWER | 25 | MATURITY_REPORT | 46 | SECURITY_FIX |
| 5 | BUILD_META | 26 | MAVEN_CENTRAL | 47 | SECURITY_OTHER |
| 6 | BUILD_SYSTEM | 27 | METRICS | 48 | SECURITY_PENTEST_REPORT |
| 7 | CERTIFICATION_REPORT | 28 | MODEL_CARD | 49 | SECURITY_POLICY |
| 8 | CHAT | 29 | NPM | 50 | SECURITY_SWID |
| 9 | CODIFIED_INFRASTRUCTURE | 30 | NUGET | 51 | SECURITY_THREAT_MODEL |
| 10 | COMPONENT_ANALYSIS_REPORT | 31 | OTHER | 52 | SOCIAL |
| 11 | CONFIGURATION | 32 | POAM | 53 | SOURCE_ARTIFACT |
| 12 | DISTRIBUTION_INTAKE | 33 | PRIVACY_ASSESSMENT | 54 | STATIC_ANALYSIS_REPORT |
| 13 | DOCUMENTATION | 34 | PRODUCT_METADATA | 55 | SUPPORT |
| 14 | DOWNLOAD | 35 | PURCHASE_ORDER | 56 | VCS |
| 15 | DYNAMIC_ANALYSIS_REPORT | 36 | QUALITY_ASSESSMENT_REPORT | 57 | VULNERABILITY_ASSERTION |
| 16 | EOL_NOTICE | 37 | QUALITY_METRICS | 58 | VULNERABILITY_DISCLOSURE_REPORT |
| 17 | EVIDENCE | 38 | RELEASE_HISTORY | 59 | VULNERABILITY_EXPLOITABILITY_ASSESSMENT |
| 18 | EXPORT_CONTROL_ASSESSMENT | 39 | RELEASE_NOTES | 60 | WEBSITE |
| 19 | FORMULATION | 40 | RISK_ASSESSMENT | | |
| 20 | FUNDING | 41 | RUNTIME_ANALYSIS_REPORT | | |
