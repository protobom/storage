-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_annotations" table
CREATE TABLE `new_annotations` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `name` text NOT NULL,
  `value` text NOT NULL,
  `is_unique` bool NOT NULL DEFAULT (false),
  `document_id` uuid NULL,
  `document_annotations` uuid NULL,
  CONSTRAINT `annotations_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE SET NULL,
  CONSTRAINT `annotations_documents_annotations` FOREIGN KEY (`document_annotations`) REFERENCES `documents` (`id`) ON DELETE SET NULL
);
-- Copy rows from old table "annotations" to new temporary table "new_annotations"
INSERT INTO `new_annotations` (`id`, `name`, `value`, `is_unique`, `document_id`) SELECT `id`, `name`, `value`, `is_unique`, `document_id` FROM `annotations`;
-- Drop "annotations" table after copying rows
DROP TABLE `annotations`;
-- Rename temporary table "new_annotations" to "annotations"
ALTER TABLE `new_annotations` RENAME TO `annotations`;
-- Create index "idx_annotations" to table: "annotations"
CREATE UNIQUE INDEX `idx_annotations` ON `annotations` (`document_id`, `name`, `value`);
-- Create index "idx_document_unique_annotations" to table: "annotations"
CREATE UNIQUE INDEX `idx_document_unique_annotations` ON `annotations` (`document_id`, `name`) WHERE is_unique = true;
-- Create "new_documents" table
CREATE TABLE `new_documents` (
  `id` uuid NOT NULL,
  `metadata_id` text NULL,
  `node_list_id` uuid NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `documents_metadata_document` FOREIGN KEY (`metadata_id`) REFERENCES `metadata` (`id`) ON DELETE SET NULL,
  CONSTRAINT `documents_node_lists_document` FOREIGN KEY (`node_list_id`) REFERENCES `node_lists` (`id`) ON DELETE SET NULL
);
-- Copy rows from old table "documents" to new temporary table "new_documents"
INSERT INTO `new_documents` (`id`) SELECT `id` FROM `documents`;
-- Drop "documents" table after copying rows
DROP TABLE `documents`;
-- Rename temporary table "new_documents" to "documents"
ALTER TABLE `new_documents` RENAME TO `documents`;
-- Create index "documents_metadata_id_key" to table: "documents"
CREATE UNIQUE INDEX `documents_metadata_id_key` ON `documents` (`metadata_id`);
-- Create index "documents_node_list_id_key" to table: "documents"
CREATE UNIQUE INDEX `documents_node_list_id_key` ON `documents` (`node_list_id`);
-- Create index "idx_documents" to table: "documents"
CREATE UNIQUE INDEX `idx_documents` ON `documents` (`metadata_id`, `node_list_id`);
-- Create "new_document_types" table
CREATE TABLE `new_document_types` (
  `id` uuid NOT NULL,
  `proto_message` blob NULL,
  `type` text NULL,
  `name` text NULL,
  `description` text NULL,
  `document_id` uuid NULL,
  `metadata_id` text NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `document_types_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE SET NULL,
  CONSTRAINT `document_types_metadata_document_types` FOREIGN KEY (`metadata_id`) REFERENCES `metadata` (`id`) ON DELETE SET NULL
);
-- Copy rows from old table "document_types" to new temporary table "new_document_types"
INSERT INTO `new_document_types` (`id`, `type`, `name`, `description`, `metadata_id`) SELECT `id`, `type`, `name`, `description`, `metadata_id` FROM `document_types`;
-- Drop "document_types" table after copying rows
DROP TABLE `document_types`;
-- Rename temporary table "new_document_types" to "document_types"
ALTER TABLE `new_document_types` RENAME TO `document_types`;
-- Create index "idx_document_types" to table: "document_types"
CREATE UNIQUE INDEX `idx_document_types` ON `document_types` (`metadata_id`, `type`, `name`, `description`);
-- Create "new_edge_types" table
CREATE TABLE `new_edge_types` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `type` text NOT NULL,
  `document_id` uuid NULL,
  `node_id` text NOT NULL,
  `to_node_id` text NOT NULL,
  CONSTRAINT `edge_types_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE SET NULL,
  CONSTRAINT `edge_types_nodes_from` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE NO ACTION,
  CONSTRAINT `edge_types_nodes_to` FOREIGN KEY (`to_node_id`) REFERENCES `nodes` (`id`) ON DELETE NO ACTION
);
-- Copy rows from old table "edge_types" to new temporary table "new_edge_types"
INSERT INTO `new_edge_types` (`id`, `type`, `node_id`, `to_node_id`) SELECT `id`, `type`, `node_id`, `to_node_id` FROM `edge_types`;
-- Drop "edge_types" table after copying rows
DROP TABLE `edge_types`;
-- Rename temporary table "new_edge_types" to "edge_types"
ALTER TABLE `new_edge_types` RENAME TO `edge_types`;
-- Create index "idx_edge_types" to table: "edge_types"
CREATE UNIQUE INDEX `idx_edge_types` ON `edge_types` (`type`, `node_id`, `to_node_id`);
-- Create index "edgetype_node_id_to_node_id" to table: "edge_types"
CREATE UNIQUE INDEX `edgetype_node_id_to_node_id` ON `edge_types` (`node_id`, `to_node_id`);
-- Create "new_external_references" table
CREATE TABLE `new_external_references` (
  `id` uuid NOT NULL,
  `proto_message` blob NULL,
  `url` text NOT NULL,
  `comment` text NOT NULL,
  `authority` text NULL,
  `type` text NOT NULL,
  `hashes` json NULL,
  `document_id` uuid NULL,
  `node_id` text NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `external_references_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE SET NULL,
  CONSTRAINT `external_references_nodes_external_references` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE SET NULL
);
-- Copy rows from old table "external_references" to new temporary table "new_external_references"
INSERT INTO `new_external_references` (`id`, `url`, `comment`, `authority`, `type`, `node_id`) SELECT `id`, `url`, `comment`, `authority`, `type`, `node_id` FROM `external_references`;
-- Drop "external_references" table after copying rows
DROP TABLE `external_references`;
-- Rename temporary table "new_external_references" to "external_references"
ALTER TABLE `new_external_references` RENAME TO `external_references`;
-- Create index "idx_external_references" to table: "external_references"
CREATE UNIQUE INDEX `idx_external_references` ON `external_references` (`node_id`, `url`, `type`);
-- Create "new_metadata" table
CREATE TABLE `new_metadata` (
  `id` text NOT NULL,
  `proto_message` blob NULL,
  `version` text NOT NULL,
  `name` text NOT NULL,
  `date` datetime NOT NULL,
  `comment` text NOT NULL,
  PRIMARY KEY (`id`)
);
-- Copy rows from old table "metadata" to new temporary table "new_metadata"
INSERT INTO `new_metadata` (`id`, `version`, `name`, `date`, `comment`) SELECT `id`, `version`, `name`, `date`, `comment` FROM `metadata`;
-- Drop "metadata" table after copying rows
DROP TABLE `metadata`;
-- Rename temporary table "new_metadata" to "metadata"
ALTER TABLE `new_metadata` RENAME TO `metadata`;
-- Create index "idx_metadata" to table: "metadata"
CREATE UNIQUE INDEX `idx_metadata` ON `metadata` (`id`, `version`, `name`);
-- Create "new_nodes" table
CREATE TABLE `new_nodes` (
  `id` text NOT NULL,
  `proto_message` blob NULL,
  `node_list_id` uuid NULL,
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
  `hashes` json NULL,
  `identifiers` json NULL,
  `document_id` uuid NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `nodes_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE SET NULL
);
-- Copy rows from old table "nodes" to new temporary table "new_nodes"
INSERT INTO `new_nodes` (`id`, `type`, `name`, `version`, `file_name`, `url_home`, `url_download`, `licenses`, `license_concluded`, `license_comments`, `copyright`, `source_info`, `comment`, `summary`, `description`, `release_date`, `build_date`, `valid_until_date`, `attribution`, `file_types`) SELECT `id`, `type`, `name`, `version`, `file_name`, `url_home`, `url_download`, `licenses`, `license_concluded`, `license_comments`, `copyright`, `source_info`, `comment`, `summary`, `description`, `release_date`, `build_date`, `valid_until_date`, `attribution`, `file_types` FROM `nodes`;
-- Drop "nodes" table after copying rows
DROP TABLE `nodes`;
-- Rename temporary table "new_nodes" to "nodes"
ALTER TABLE `new_nodes` RENAME TO `nodes`;
-- Create index "idx_nodes" to table: "nodes"
CREATE UNIQUE INDEX `idx_nodes` ON `nodes` (`id`, `node_list_id`);
-- Create "new_node_lists" table
CREATE TABLE `new_node_lists` (
  `id` uuid NOT NULL,
  `proto_message` blob NULL,
  `root_elements` json NOT NULL,
  PRIMARY KEY (`id`)
);
-- Copy rows from old table "node_lists" to new temporary table "new_node_lists"
INSERT INTO `new_node_lists` (`id`, `root_elements`) SELECT `id`, `root_elements` FROM `node_lists`;
-- Drop "node_lists" table after copying rows
DROP TABLE `node_lists`;
-- Rename temporary table "new_node_lists" to "node_lists"
ALTER TABLE `new_node_lists` RENAME TO `node_lists`;
-- Create "new_persons" table
CREATE TABLE `new_persons` (
  `id` uuid NOT NULL,
  `proto_message` blob NULL,
  `name` text NOT NULL,
  `is_org` bool NOT NULL,
  `email` text NOT NULL,
  `url` text NOT NULL,
  `phone` text NOT NULL,
  `metadata_id` text NULL,
  `node_suppliers` text NULL,
  `node_id` text NULL,
  `document_id` uuid NULL,
  `person_contacts` uuid NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `persons_metadata_authors` FOREIGN KEY (`metadata_id`) REFERENCES `metadata` (`id`) ON DELETE SET NULL,
  CONSTRAINT `persons_nodes_suppliers` FOREIGN KEY (`node_suppliers`) REFERENCES `nodes` (`id`) ON DELETE SET NULL,
  CONSTRAINT `persons_nodes_originators` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE SET NULL,
  CONSTRAINT `persons_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE SET NULL,
  CONSTRAINT `persons_persons_contacts` FOREIGN KEY (`person_contacts`) REFERENCES `persons` (`id`) ON DELETE SET NULL
);
-- Copy rows from old table "persons" to new temporary table "new_persons"
INSERT INTO `new_persons` (`id`, `name`, `is_org`, `email`, `url`, `phone`, `metadata_id`, `node_suppliers`, `node_id`, `person_contacts`) SELECT `id`, `name`, `is_org`, `email`, `url`, `phone`, `metadata_id`, `node_suppliers`, `node_id`, `person_contacts` FROM `persons`;
-- Drop "persons" table after copying rows
DROP TABLE `persons`;
-- Rename temporary table "new_persons" to "persons"
ALTER TABLE `new_persons` RENAME TO `persons`;
-- Create index "idx_person_metadata_id" to table: "persons"
CREATE UNIQUE INDEX `idx_person_metadata_id` ON `persons` (`metadata_id`, `name`, `is_org`, `email`, `url`, `phone`) WHERE metadata_id IS NOT NULL AND node_id IS NULL;
-- Create index "idx_person_node_id" to table: "persons"
CREATE UNIQUE INDEX `idx_person_node_id` ON `persons` (`node_id`, `name`, `is_org`, `email`, `url`, `phone`) WHERE metadata_id IS NULL AND node_id IS NOT NULL;
-- Create "new_purposes" table
CREATE TABLE `new_purposes` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `primary_purpose` text NOT NULL,
  `node_id` text NULL,
  `document_id` uuid NULL,
  CONSTRAINT `purposes_nodes_primary_purpose` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE SET NULL,
  CONSTRAINT `purposes_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE SET NULL
);
-- Copy rows from old table "purposes" to new temporary table "new_purposes"
INSERT INTO `new_purposes` (`id`, `primary_purpose`, `node_id`) SELECT `id`, `primary_purpose`, `node_id` FROM `purposes`;
-- Drop "purposes" table after copying rows
DROP TABLE `purposes`;
-- Rename temporary table "new_purposes" to "purposes"
ALTER TABLE `new_purposes` RENAME TO `purposes`;
-- Create index "idx_purposes" to table: "purposes"
CREATE UNIQUE INDEX `idx_purposes` ON `purposes` (`node_id`, `primary_purpose`);
-- Create "new_tools" table
CREATE TABLE `new_tools` (
  `id` uuid NOT NULL,
  `proto_message` blob NULL,
  `name` text NOT NULL,
  `version` text NOT NULL,
  `vendor` text NOT NULL,
  `metadata_id` text NULL,
  `document_id` uuid NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `tools_metadata_tools` FOREIGN KEY (`metadata_id`) REFERENCES `metadata` (`id`) ON DELETE SET NULL,
  CONSTRAINT `tools_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE SET NULL
);
-- Copy rows from old table "tools" to new temporary table "new_tools"
INSERT INTO `new_tools` (`id`, `name`, `version`, `vendor`, `metadata_id`) SELECT `id`, `name`, `version`, `vendor`, `metadata_id` FROM `tools`;
-- Drop "tools" table after copying rows
DROP TABLE `tools`;
-- Rename temporary table "new_tools" to "tools"
ALTER TABLE `new_tools` RENAME TO `tools`;
-- Create index "idx_tools" to table: "tools"
CREATE UNIQUE INDEX `idx_tools` ON `tools` (`metadata_id`, `name`, `version`, `vendor`);
-- Create "new_node_list_nodes" table
CREATE TABLE `new_node_list_nodes` (
  `node_list_id` uuid NOT NULL,
  `node_id` text NOT NULL,
  PRIMARY KEY (`node_list_id`, `node_id`),
  CONSTRAINT `node_list_nodes_node_list_id` FOREIGN KEY (`node_list_id`) REFERENCES `node_lists` (`id`) ON DELETE CASCADE,
  CONSTRAINT `node_list_nodes_node_id` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE
);
-- Copy rows from old table "node_list_nodes" to new temporary table "new_node_list_nodes"
INSERT INTO `new_node_list_nodes` (`node_list_id`, `node_id`) SELECT `node_list_id`, `node_id` FROM `node_list_nodes`;
-- Drop "node_list_nodes" table after copying rows
DROP TABLE `node_list_nodes`;
-- Rename temporary table "new_node_list_nodes" to "node_list_nodes"
ALTER TABLE `new_node_list_nodes` RENAME TO `node_list_nodes`;
-- Create "properties" table
CREATE TABLE `properties` (
  `id` uuid NOT NULL,
  `proto_message` blob NULL,
  `name` text NOT NULL,
  `data` text NOT NULL,
  `node_id` text NULL,
  `document_id` uuid NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `properties_nodes_properties` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE SET NULL,
  CONSTRAINT `properties_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE SET NULL
);
-- Create index "idx_property" to table: "properties"
CREATE UNIQUE INDEX `idx_property` ON `properties` (`name`, `data`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
