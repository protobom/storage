-- --------------------------------------------------------------
-- SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
-- SPDX-FileType: SOURCE
-- SPDX-License-Identifier: Apache-2.0
-- --------------------------------------------------------------
--
-- ClickHouse schema for storing protobom SBOM data.
--
-- Unlike the relational schemas (schema.ansi.sql / schema.postgres.sql), this
-- variant is deliberately ClickHouse-idiomatic and denormalized. It pairs two
-- concerns:
--
--   * FIDELITY: the full serialized sbom.Document is kept as protobuf bytes in
--     documents.document, so retrieval by id is a single lookup + unmarshal and
--     is always lossless. The wide columns below carry NO round-trip duty.
--
--   * QUERYABILITY: a node and all its sub-collections (hashes, identifiers,
--     licenses, purposes, external references, persons) collapse into ONE wide
--     row using Array / Map / Tuple columns, so analytical queries are
--     single-table and vectorized (e.g. hashes[3], has(licenses, 'MIT')).
--
-- Design notes:
--   * Enums are stored as the raw protobuf enum NUMBER in Int16 columns (same
--     convention as the relational schemas). See README.md for the code tables.
--   * Every table uses ReplacingMergeTree(stored_at): re-storing the same
--     document is idempotent (duplicate rows collapse at merge time). Reads that
--     must be exact use the FINAL modifier.
--   * ClickHouse has no enforced primary/foreign keys; ORDER BY defines both the
--     sort key and the ReplacingMergeTree dedup key.
--   * NodeList is 1:1 with a document, so it is folded into `documents`
--     (root_elements) rather than given its own table.
--   * Person.contacts (recursive) is intentionally not modeled in the query
--     tables; full fidelity lives in the document blob.
--   * DateTime64(3) covers years 1900-2299; unset protobuf timestamps map to the
--     unix epoch (1970), matching the ent backend.

-- ====================================================================
-- documents: retrieval blob + document-level (metadata) query surface
-- ====================================================================
CREATE TABLE IF NOT EXISTS documents
(
    id             String,        -- Metadata.id (retrieval key)
    document       String,        -- serialized sbom.Document (protobuf bytes)
    version        String,        -- Metadata.version
    name           String,        -- Metadata.name
    created_date   DateTime64(3), -- Metadata.date
    comment        String,        -- Metadata.comment
    document_types Array(Tuple(type Int16, name String, description String)),
    tools          Array(Tuple(name String, version String, vendor String)),
    authors        Array(Tuple(name String, is_org UInt8, email String, url String, phone String)),
    source_format  String,        -- Metadata.sourceData.format
    source_size    Int64,         -- Metadata.sourceData.size
    source_uri     String,        -- Metadata.sourceData.uri
    source_hashes  Map(Int16, String),
    root_elements  Array(String), -- NodeList.root_elements
    stored_at      DateTime64(3)
)
ENGINE = ReplacingMergeTree(stored_at)
ORDER BY id;

-- ====================================================================
-- nodes: the wide component table (one row per (document, node))
-- ====================================================================
CREATE TABLE IF NOT EXISTS nodes
(
    document_id       String,        -- owning document (denormalized for scoping)
    id                String,        -- protobom node id
    type              Int16,         -- Node.NodeType (0 PACKAGE, 1 FILE)
    name              String,
    version           String,
    file_name         String,
    url_home          String,
    url_download      String,
    license_concluded String,
    license_comments  String,
    copyright         String,
    source_info       String,
    comment           String,
    summary           String,
    description       String,
    release_date      DateTime64(3),
    build_date        DateTime64(3),
    valid_until_date  DateTime64(3),
    licenses          Array(String), -- Node.licenses
    attribution       Array(String), -- Node.attribution
    file_types        Array(String), -- Node.file_types
    identifiers       Map(Int16, String), -- SoftwareIdentifierType -> value (purl/cpe/gitoid)
    hashes            Map(Int16, String), -- HashAlgorithm -> value
    purposes          Array(Int16),  -- Node.primary_purpose codes
    properties        Array(Tuple(name String, data String)),
    external_references Array(Tuple(url String, comment String, authority String,
                                    type Int16, hashes Map(Int16, String))),
    suppliers         Array(Tuple(name String, is_org UInt8, email String, url String, phone String)),
    originators       Array(Tuple(name String, is_org UInt8, email String, url String, phone String)),
    stored_at         DateTime64(3),
    INDEX idx_name       name                   TYPE bloom_filter GRANULARITY 1,
    INDEX idx_hash_vals  mapValues(hashes)       TYPE bloom_filter GRANULARITY 1,
    INDEX idx_ident_vals mapValues(identifiers)  TYPE bloom_filter GRANULARITY 1
)
ENGINE = ReplacingMergeTree(stored_at)
ORDER BY (document_id, id);

-- ====================================================================
-- edges: relationships (Edge.to collapsed into an array)
-- ====================================================================
CREATE TABLE IF NOT EXISTS edges
(
    document_id  String,
    type         Int16,         -- Edge.Type
    from_node_id String,        -- Edge.from
    to_node_ids  Array(String), -- Edge.to
    stored_at    DateTime64(3)
)
ENGINE = ReplacingMergeTree(stored_at)
ORDER BY (document_id, from_node_id, type);
