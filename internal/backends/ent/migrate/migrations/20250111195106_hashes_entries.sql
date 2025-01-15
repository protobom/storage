-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_hashes_entries" table
CREATE TABLE `new_hashes_entries` (
  `id` uuid NOT NULL,
  `hash_algorithm` text NOT NULL,
  `hash_data` text NOT NULL,
  `document_id` uuid NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `hashes_entries_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE
);
-- Copy rows from old table "hashes_entries" to new temporary table "new_hashes_entries"
INSERT INTO `new_hashes_entries` (`id`, `hash_data`) SELECT `id`, `hash_data` FROM `hashes_entries`;
-- Drop "hashes_entries" table after copying rows
DROP TABLE `hashes_entries`;
-- Rename temporary table "new_hashes_entries" to "hashes_entries"
ALTER TABLE `new_hashes_entries` RENAME TO `hashes_entries`;
-- Create index "idx_hashes" to table: "hashes_entries"
CREATE UNIQUE INDEX `idx_hashes` ON `hashes_entries` (`hash_algorithm`, `hash_data`);
-- Create "new_external_references" table
CREATE TABLE `new_external_references` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `url` text NOT NULL,
  `comment` text NOT NULL,
  `authority` text NULL,
  `type` text NOT NULL,
  `document_id` uuid NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `external_references_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE
);
-- Copy rows from old table "external_references" to new temporary table "new_external_references"
INSERT INTO `new_external_references` (`id`, `proto_message`, `url`, `comment`, `authority`, `type`, `document_id`) SELECT `id`, `proto_message`, `url`, `comment`, `authority`, `type`, `document_id` FROM `external_references`;
-- Drop "external_references" table after copying rows
DROP TABLE `external_references`;
-- Rename temporary table "new_external_references" to "external_references"
ALTER TABLE `new_external_references` RENAME TO `external_references`;
-- Create "ext_ref_hashes" table
CREATE TABLE `ext_ref_hashes` (
  `ext_ref_id` uuid NOT NULL,
  `hash_entry_id` uuid NOT NULL,
  PRIMARY KEY (`ext_ref_id`, `hash_entry_id`),
  CONSTRAINT `ext_ref_hashes_ext_ref_id` FOREIGN KEY (`ext_ref_id`) REFERENCES `external_references` (`id`) ON DELETE CASCADE,
  CONSTRAINT `ext_ref_hashes_hash_entry_id` FOREIGN KEY (`hash_entry_id`) REFERENCES `hashes_entries` (`id`) ON DELETE CASCADE
);
-- Create "node_external_references" table
CREATE TABLE `node_external_references` (
  `node_id` uuid NOT NULL,
  `external_reference_id` uuid NOT NULL,
  PRIMARY KEY (`node_id`, `external_reference_id`),
  CONSTRAINT `node_external_references_node_id` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `node_external_references_external_reference_id` FOREIGN KEY (`external_reference_id`) REFERENCES `external_references` (`id`) ON DELETE CASCADE
);
-- Create "node_hashes" table
CREATE TABLE `node_hashes` (
  `node_id` uuid NOT NULL,
  `hash_entry_id` uuid NOT NULL,
  PRIMARY KEY (`node_id`, `hash_entry_id`),
  CONSTRAINT `node_hashes_node_id` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `node_hashes_hash_entry_id` FOREIGN KEY (`hash_entry_id`) REFERENCES `hashes_entries` (`id`) ON DELETE CASCADE
);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
