-- --------------------------------------------------------------
-- SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
-- SPDX-FileType: SOURCE
-- SPDX-License-Identifier: Apache-2.0
-- --------------------------------------------------------------
--
-- PostgreSQL schema for storing protobom SBOM data.
--
-- This is the PostgreSQL-tuned counterpart of schema.ansi.sql. It is identical
-- in structure but uses native Postgres types and conventions:
--   * TEXT for all variable-length strings (no length limits needed)
--   * JSONB for inline scalar arrays (queryable, indexable)
--   * TIMESTAMPTZ for UTC instants
--   * BOOLEAN for flags
--   * SMALLINT for protobuf enum codes (stored as the raw enum NUMBER)
--   * GENERATED ALWAYS AS IDENTITY for surrogate keys
--
-- See README.md for design notes and the full protobuf enum code reference.
-- Tables are declared in dependency order so the file applies top to bottom.

-- ====================================================================
-- Document graph roots
-- ====================================================================

-- A single SBOM document. id is the document identifier (CycloneDX serial
-- number or SPDX SPDXID), mirroring Metadata.id in the proto.
CREATE TABLE documents (
    id TEXT PRIMARY KEY
);

-- Document-level metadata (1:1 with documents).
CREATE TABLE metadata (
    document_id TEXT PRIMARY KEY REFERENCES documents (id) ON DELETE CASCADE,
    version     TEXT        NOT NULL DEFAULT '',
    name        TEXT        NOT NULL DEFAULT '',
    date        TIMESTAMPTZ,                 -- Metadata.date (creation time)
    comment     TEXT        NOT NULL DEFAULT ''
);

-- Metadata.documentTypes (repeated DocumentType).
-- type stores DocumentType.SBOMType (0 OTHER .. 8 DECOMISSION, NULL = unset).
CREATE TABLE document_types (
    id          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    document_id TEXT NOT NULL REFERENCES metadata (document_id) ON DELETE CASCADE,
    type        SMALLINT,
    name        TEXT,
    description TEXT
);
CREATE INDEX idx_document_types_document ON document_types (document_id);

-- Metadata.tools (repeated Tool).
CREATE TABLE tools (
    id          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    document_id TEXT NOT NULL REFERENCES metadata (document_id) ON DELETE CASCADE,
    name        TEXT NOT NULL DEFAULT '',
    version     TEXT NOT NULL DEFAULT '',
    vendor      TEXT NOT NULL DEFAULT ''
);
CREATE INDEX idx_tools_document ON tools (document_id);

-- Metadata.source_data: provenance of the original SBOM document (1:1 metadata).
CREATE TABLE source_data (
    id          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    document_id TEXT   NOT NULL UNIQUE REFERENCES metadata (document_id) ON DELETE CASCADE,
    format      TEXT   NOT NULL DEFAULT '',  -- e.g. text/spdx+json;version=2.3
    size        BIGINT,                      -- original size in bytes
    uri         TEXT
);

-- SourceData.hashes: map<HashAlgorithm, value> of the original document.
-- algorithm stores HashAlgorithm (0 UNKNOWN, 1 MD5, 2 SHA1, 3 SHA256, ...).
CREATE TABLE source_data_hashes (
    source_data_id BIGINT   NOT NULL REFERENCES source_data (id) ON DELETE CASCADE,
    algorithm      SMALLINT NOT NULL,
    hash_value     TEXT     NOT NULL,
    PRIMARY KEY (source_data_id, algorithm)
);

-- ====================================================================
-- Nodes
-- ====================================================================

-- NodeList: graph container holding the document's nodes and edges (1:1 doc).
CREATE TABLE node_lists (
    id            BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    document_id   TEXT  NOT NULL UNIQUE REFERENCES documents (id) ON DELETE CASCADE,
    root_elements JSONB                  -- NodeList.root_elements (array of ids)
);

-- Node: a software component (the central vertex of the SBOM graph).
-- Nodes are keyed by their protobom id and linked to node lists through
-- node_list_nodes, allowing a node to be shared across documents.
-- type stores Node.NodeType (0 PACKAGE, 1 FILE).
CREATE TABLE nodes (
    id                TEXT     PRIMARY KEY,
    type              SMALLINT NOT NULL DEFAULT 0,
    name              TEXT     NOT NULL DEFAULT '',
    version           TEXT     NOT NULL DEFAULT '',
    file_name         TEXT     NOT NULL DEFAULT '',
    url_home          TEXT     NOT NULL DEFAULT '',
    url_download      TEXT     NOT NULL DEFAULT '',
    license_concluded TEXT     NOT NULL DEFAULT '',
    license_comments  TEXT     NOT NULL DEFAULT '',
    copyright         TEXT     NOT NULL DEFAULT '',
    source_info       TEXT     NOT NULL DEFAULT '',
    comment           TEXT     NOT NULL DEFAULT '',
    summary           TEXT     NOT NULL DEFAULT '',
    description       TEXT     NOT NULL DEFAULT '',
    release_date      TIMESTAMPTZ,
    build_date        TIMESTAMPTZ,
    valid_until_date  TIMESTAMPTZ,
    -- Scalar string arrays stored inline (hybrid model).
    licenses          JSONB,             -- Node.licenses (repeated string)
    attribution       JSONB,             -- Node.attribution (repeated string)
    file_types        JSONB              -- Node.file_types (repeated string)
);
CREATE INDEX idx_nodes_name ON nodes (name);
CREATE INDEX idx_nodes_type ON nodes (type);

-- Membership of nodes in node lists (NodeList.nodes), many-to-many.
CREATE TABLE node_list_nodes (
    node_list_id BIGINT NOT NULL REFERENCES node_lists (id) ON DELETE CASCADE,
    node_id      TEXT   NOT NULL REFERENCES nodes (id) ON DELETE CASCADE,
    PRIMARY KEY (node_list_id, node_id)
);
CREATE INDEX idx_node_list_nodes_node ON node_list_nodes (node_id);

-- ====================================================================
-- Node sub-collections
-- ====================================================================

-- Node.identifiers: map<SoftwareIdentifierType, value>.
-- type stores SoftwareIdentifierType (0 UNKNOWN, 1 PURL, 2 CPE22, 3 CPE23, 4 GITOID).
CREATE TABLE node_identifiers (
    node_id TEXT     NOT NULL REFERENCES nodes (id) ON DELETE CASCADE,
    type    SMALLINT NOT NULL,
    value   TEXT     NOT NULL,
    PRIMARY KEY (node_id, type)
);
CREATE INDEX idx_node_identifiers_value ON node_identifiers (value);

-- Node.hashes: map<HashAlgorithm, value>.
CREATE TABLE node_hashes (
    node_id    TEXT     NOT NULL REFERENCES nodes (id) ON DELETE CASCADE,
    algorithm  SMALLINT NOT NULL,
    hash_value TEXT     NOT NULL,
    PRIMARY KEY (node_id, algorithm)
);
CREATE INDEX idx_node_hashes_value ON node_hashes (hash_value);

-- Node.primary_purpose: repeated Purpose.
-- purpose stores Purpose (0 UNKNOWN_PURPOSE, 1 APPLICATION, 2 ARCHIVE, ...).
CREATE TABLE node_purposes (
    node_id TEXT     NOT NULL REFERENCES nodes (id) ON DELETE CASCADE,
    purpose SMALLINT NOT NULL,
    PRIMARY KEY (node_id, purpose)
);

-- Node.properties: repeated Property (key/value pairs).
CREATE TABLE node_properties (
    id      BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    node_id TEXT NOT NULL REFERENCES nodes (id) ON DELETE CASCADE,
    name    TEXT NOT NULL DEFAULT '',
    data    TEXT NOT NULL DEFAULT ''
);
CREATE INDEX idx_node_properties_node ON node_properties (node_id);

-- Node.external_references: repeated ExternalReference.
-- type stores ExternalReference.ExternalReferenceType (0 UNKNOWN .. 60 WEBSITE).
CREATE TABLE external_references (
    id        BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    node_id   TEXT     NOT NULL REFERENCES nodes (id) ON DELETE CASCADE,
    url       TEXT     NOT NULL DEFAULT '',
    comment   TEXT     NOT NULL DEFAULT '',
    authority TEXT,
    type      SMALLINT NOT NULL DEFAULT 0
);
CREATE INDEX idx_external_references_node ON external_references (node_id);

-- ExternalReference.hashes: map<HashAlgorithm, value>.
CREATE TABLE external_reference_hashes (
    external_reference_id BIGINT   NOT NULL REFERENCES external_references (id) ON DELETE CASCADE,
    algorithm             SMALLINT NOT NULL,
    hash_value            TEXT     NOT NULL,
    PRIMARY KEY (external_reference_id, algorithm)
);

-- ====================================================================
-- Edges (relationships between nodes)
-- ====================================================================

-- Edge: a typed relationship rooted at one node (Edge.from). The target nodes
-- (Edge.to, repeated) live in edge_targets.
-- type stores Edge.Type (0 UNKNOWN, 1 amends, 2 ancestor, ... 44 variant).
CREATE TABLE edges (
    id           BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    node_list_id BIGINT   NOT NULL REFERENCES node_lists (id) ON DELETE CASCADE,
    type         SMALLINT NOT NULL DEFAULT 0,
    from_node_id TEXT     NOT NULL REFERENCES nodes (id) ON DELETE CASCADE
);
CREATE INDEX idx_edges_node_list ON edges (node_list_id);
CREATE INDEX idx_edges_from_node ON edges (from_node_id);

-- Edge.to: the target nodes of an edge.
CREATE TABLE edge_targets (
    edge_id    BIGINT NOT NULL REFERENCES edges (id) ON DELETE CASCADE,
    to_node_id TEXT   NOT NULL REFERENCES nodes (id) ON DELETE CASCADE,
    PRIMARY KEY (edge_id, to_node_id)
);
CREATE INDEX idx_edge_targets_node ON edge_targets (to_node_id);

-- ====================================================================
-- Persons (authors, suppliers, originators, contacts)
-- ====================================================================

-- Person: an individual or organization, attached to exactly one owner via one
-- of the nullable foreign keys, matching the four Person references in the proto:
--   metadata_id            -> Metadata.authors
--   supplier_of_node_id    -> Node.suppliers
--   originator_of_node_id  -> Node.originators
--   contact_of_person_id   -> Person.contacts (self reference, recursive)
CREATE TABLE persons (
    id                    BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name                  TEXT    NOT NULL DEFAULT '',
    is_org                BOOLEAN NOT NULL DEFAULT FALSE,
    email                 TEXT    NOT NULL DEFAULT '',
    url                   TEXT    NOT NULL DEFAULT '',
    phone                 TEXT    NOT NULL DEFAULT '',
    metadata_id           TEXT    REFERENCES metadata (document_id) ON DELETE CASCADE,
    supplier_of_node_id   TEXT    REFERENCES nodes (id) ON DELETE CASCADE,
    originator_of_node_id TEXT    REFERENCES nodes (id) ON DELETE CASCADE,
    contact_of_person_id  BIGINT  REFERENCES persons (id) ON DELETE CASCADE
);
CREATE INDEX idx_persons_metadata ON persons (metadata_id);
CREATE INDEX idx_persons_supplier_node ON persons (supplier_of_node_id);
CREATE INDEX idx_persons_originator_node ON persons (originator_of_node_id);
CREATE INDEX idx_persons_contact_of ON persons (contact_of_person_id);
