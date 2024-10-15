-- Create "annotations" table
CREATE TABLE `annotations` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `name` text NOT NULL,
  `value` text NOT NULL,
  `is_unique` bool NOT NULL DEFAULT (false),
  `document_id` text NOT NULL,
  CONSTRAINT `annotations_documents_annotations` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE NO ACTION
);
-- Create index "idx_annotations" to table: "annotations"
CREATE UNIQUE INDEX `idx_annotations` ON `annotations` (`document_id`, `name`, `value`);
-- Create index "idx_document_unique_annotations" to table: "annotations"
CREATE UNIQUE INDEX `idx_document_unique_annotations` ON `annotations` (`document_id`, `name`) WHERE is_unique = true;
-- Create "documents" table
CREATE TABLE `documents` (
  `id` text NOT NULL,
  PRIMARY KEY (`id`)
);
-- Create "document_types" table
CREATE TABLE `document_types` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `type` text NULL,
  `name` text NULL,
  `description` text NULL,
  `metadata_id` text NULL,
  CONSTRAINT `document_types_metadata_document_types` FOREIGN KEY (`metadata_id`) REFERENCES `metadata` (`id`) ON DELETE SET NULL
);
-- Create index "idx_document_types" to table: "document_types"
CREATE UNIQUE INDEX `idx_document_types` ON `document_types` (`metadata_id`, `type`, `name`, `description`);
-- Create "edge_types" table
CREATE TABLE `edge_types` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `type` text NOT NULL,
  `node_id` text NOT NULL,
  `to_node_id` text NOT NULL,
  CONSTRAINT `edge_types_nodes_from` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE NO ACTION,
  CONSTRAINT `edge_types_nodes_to` FOREIGN KEY (`to_node_id`) REFERENCES `nodes` (`id`) ON DELETE NO ACTION
);
-- Create index "idx_edge_types" to table: "edge_types"
CREATE UNIQUE INDEX `idx_edge_types` ON `edge_types` (`type`, `node_id`, `to_node_id`);
-- Create index "edgetype_node_id_to_node_id" to table: "edge_types"
CREATE UNIQUE INDEX `edgetype_node_id_to_node_id` ON `edge_types` (`node_id`, `to_node_id`);
-- Create "external_references" table
CREATE TABLE `external_references` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `url` text NOT NULL,
  `comment` text NOT NULL,
  `authority` text NULL,
  `type` text NOT NULL,
  `node_id` text NULL,
  CONSTRAINT `external_references_nodes_external_references` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE SET NULL
);
-- Create index "idx_external_references" to table: "external_references"
CREATE UNIQUE INDEX `idx_external_references` ON `external_references` (`node_id`, `url`, `type`);
-- Create "hashes_entries" table
CREATE TABLE `hashes_entries` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `hash_algorithm_type` text NOT NULL,
  `hash_data` text NOT NULL,
  `external_reference_id` integer NULL,
  `node_id` text NULL,
  CONSTRAINT `hashes_entries_external_references_hashes` FOREIGN KEY (`external_reference_id`) REFERENCES `external_references` (`id`) ON DELETE SET NULL,
  CONSTRAINT `hashes_entries_nodes_hashes` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE SET NULL
);
-- Create index "idx_hashes_entries" to table: "hashes_entries"
CREATE UNIQUE INDEX `idx_hashes_entries` ON `hashes_entries` (`external_reference_id`, `node_id`, `hash_algorithm_type`, `hash_data`);
-- Create "identifiers_entries" table
CREATE TABLE `identifiers_entries` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `software_identifier_type` text NOT NULL,
  `software_identifier_value` text NOT NULL,
  `node_id` text NULL,
  CONSTRAINT `identifiers_entries_nodes_identifiers` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE SET NULL
);
-- Create index "idx_identifiers_entries" to table: "identifiers_entries"
CREATE UNIQUE INDEX `idx_identifiers_entries` ON `identifiers_entries` (`node_id`, `software_identifier_type`, `software_identifier_value`);
-- Create "metadata" table
CREATE TABLE `metadata` (
  `id` text NOT NULL,
  `version` text NOT NULL,
  `name` text NOT NULL,
  `date` datetime NOT NULL,
  `comment` text NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `metadata_documents_metadata` FOREIGN KEY (`id`) REFERENCES `documents` (`id`) ON DELETE NO ACTION
);
-- Create index "idx_metadata" to table: "metadata"
CREATE UNIQUE INDEX `idx_metadata` ON `metadata` (`id`, `version`, `name`);
-- Create "nodes" table
CREATE TABLE `nodes` (
  `id` text NOT NULL,
  `type` text NOT NULL,
  `name` text NOT NULL,
  `version` text NOT NULL,
  `file_name` text NOT NULL,
  `url_home` text NOT NULL,
  `url_download` text NOT NULL,
  `licenses` json NOT NULL,
  `license_concluded` text NOT NULL,
  `license_comments` text NOT NULL,
  `copyright` text NOT NULL,
  `source_info` text NOT NULL,
  `comment` text NOT NULL,
  `summary` text NOT NULL,
  `description` text NOT NULL,
  `release_date` datetime NOT NULL,
  `build_date` datetime NOT NULL,
  `valid_until_date` datetime NOT NULL,
  `attribution` json NOT NULL,
  `file_types` json NOT NULL,
  PRIMARY KEY (`id`)
);
-- Create "node_lists" table
CREATE TABLE `node_lists` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `root_elements` json NOT NULL,
  `document_id` text NOT NULL,
  CONSTRAINT `node_lists_documents_node_list` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE NO ACTION
);
-- Create index "node_lists_document_id_key" to table: "node_lists"
CREATE UNIQUE INDEX `node_lists_document_id_key` ON `node_lists` (`document_id`);
-- Create "persons" table
CREATE TABLE `persons` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `name` text NOT NULL,
  `is_org` bool NOT NULL,
  `email` text NOT NULL,
  `url` text NOT NULL,
  `phone` text NOT NULL,
  `metadata_id` text NULL,
  `node_suppliers` text NULL,
  `node_id` text NULL,
  `person_contacts` integer NULL,
  CONSTRAINT `persons_metadata_authors` FOREIGN KEY (`metadata_id`) REFERENCES `metadata` (`id`) ON DELETE SET NULL,
  CONSTRAINT `persons_nodes_suppliers` FOREIGN KEY (`node_suppliers`) REFERENCES `nodes` (`id`) ON DELETE SET NULL,
  CONSTRAINT `persons_nodes_originators` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE SET NULL,
  CONSTRAINT `persons_persons_contacts` FOREIGN KEY (`person_contacts`) REFERENCES `persons` (`id`) ON DELETE SET NULL
);
-- Create index "idx_person_metadata_id" to table: "persons"
CREATE UNIQUE INDEX `idx_person_metadata_id` ON `persons` (`metadata_id`, `name`, `is_org`, `email`, `url`, `phone`) WHERE metadata_id IS NOT NULL AND node_id IS NULL;
-- Create index "idx_person_node_id" to table: "persons"
CREATE UNIQUE INDEX `idx_person_node_id` ON `persons` (`node_id`, `name`, `is_org`, `email`, `url`, `phone`) WHERE metadata_id IS NULL AND node_id IS NOT NULL;
-- Create "purposes" table
CREATE TABLE `purposes` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `primary_purpose` text NOT NULL,
  `node_id` text NULL,
  CONSTRAINT `purposes_nodes_primary_purpose` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE SET NULL
);
-- Create index "idx_purposes" to table: "purposes"
CREATE UNIQUE INDEX `idx_purposes` ON `purposes` (`node_id`, `primary_purpose`);
-- Create "tools" table
CREATE TABLE `tools` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `name` text NOT NULL,
  `version` text NOT NULL,
  `vendor` text NOT NULL,
  `metadata_id` text NULL,
  CONSTRAINT `tools_metadata_tools` FOREIGN KEY (`metadata_id`) REFERENCES `metadata` (`id`) ON DELETE SET NULL
);
-- Create index "idx_tools" to table: "tools"
CREATE UNIQUE INDEX `idx_tools` ON `tools` (`metadata_id`, `name`, `version`, `vendor`);
-- Create "node_list_nodes" table
CREATE TABLE `node_list_nodes` (
  `node_list_id` integer NOT NULL,
  `node_id` text NOT NULL,
  PRIMARY KEY (`node_list_id`, `node_id`),
  CONSTRAINT `node_list_nodes_node_list_id` FOREIGN KEY (`node_list_id`) REFERENCES `node_lists` (`id`) ON DELETE CASCADE,
  CONSTRAINT `node_list_nodes_node_id` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE
);
