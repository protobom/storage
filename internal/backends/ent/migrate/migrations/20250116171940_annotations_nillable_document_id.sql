-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_metadata" table
CREATE TABLE `new_metadata` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `native_id` text NOT NULL,
  `version` text NOT NULL,
  `name` text NOT NULL,
  `date` datetime NOT NULL,
  `comment` text NOT NULL,
  PRIMARY KEY (`id`)
);
-- Copy rows from old table "metadata" to new temporary table "new_metadata"
INSERT INTO `new_metadata` (`id`, `proto_message`, `native_id`, `version`, `name`, `date`, `comment`) SELECT `id`, `proto_message`, `native_id`, `version`, `name`, `date`, `comment` FROM `metadata`;
-- Drop "metadata" table after copying rows
DROP TABLE `metadata`;
-- Rename temporary table "new_metadata" to "metadata"
ALTER TABLE `new_metadata` RENAME TO `metadata`;
-- Create index "idx_metadata" to table: "metadata"
CREATE UNIQUE INDEX `idx_metadata` ON `metadata` (`native_id`, `version`, `name`);
-- Create "new_node_lists" table
CREATE TABLE `new_node_lists` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `root_elements` json NOT NULL,
  PRIMARY KEY (`id`)
);
-- Copy rows from old table "node_lists" to new temporary table "new_node_lists"
INSERT INTO `new_node_lists` (`id`, `proto_message`, `root_elements`) SELECT `id`, `proto_message`, `root_elements` FROM `node_lists`;
-- Drop "node_lists" table after copying rows
DROP TABLE `node_lists`;
-- Rename temporary table "new_node_lists" to "node_lists"
ALTER TABLE `new_node_lists` RENAME TO `node_lists`;
-- Create "new_documents" table
CREATE TABLE `new_documents` (
  `id` uuid NOT NULL,
  `metadata_id` uuid NULL,
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
-- Create index "idx_annotations_node_id" to table: "annotations"
CREATE INDEX `idx_annotations_node_id` ON `annotations` (`node_id`);
-- Create index "idx_annotations_document_id" to table: "annotations"
CREATE INDEX `idx_annotations_document_id` ON `annotations` (`document_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
