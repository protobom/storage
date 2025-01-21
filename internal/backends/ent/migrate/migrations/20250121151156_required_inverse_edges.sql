-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_document_types" table
CREATE TABLE `new_document_types` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `type` text NULL,
  `name` text NULL,
  `description` text NULL,
  `document_id` uuid NULL,
  `metadata_id` uuid NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `document_types_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE,
  CONSTRAINT `document_types_metadata_document_types` FOREIGN KEY (`metadata_id`) REFERENCES `metadata` (`id`) ON DELETE CASCADE
);
-- Copy rows from old table "document_types" to new temporary table "new_document_types"
INSERT INTO `new_document_types` (`id`, `proto_message`, `type`, `name`, `description`, `document_id`, `metadata_id`) SELECT `id`, `proto_message`, `type`, `name`, `description`, `document_id`, `metadata_id` FROM `document_types`;
-- Drop "document_types" table after copying rows
DROP TABLE `document_types`;
-- Rename temporary table "new_document_types" to "document_types"
ALTER TABLE `new_document_types` RENAME TO `document_types`;
-- Create index "document_types_proto_message_key" to table: "document_types"
CREATE UNIQUE INDEX `document_types_proto_message_key` ON `document_types` (`proto_message`);
-- Create index "idx_document_types" to table: "document_types"
CREATE UNIQUE INDEX `idx_document_types` ON `document_types` (`metadata_id`, `type`, `name`, `description`);
-- Create "new_purposes" table
CREATE TABLE `new_purposes` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `primary_purpose` text NOT NULL,
  `node_id` uuid NOT NULL,
  `document_id` uuid NULL,
  CONSTRAINT `purposes_nodes_primary_purpose` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `purposes_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE
);
-- Copy rows from old table "purposes" to new temporary table "new_purposes"
INSERT INTO `new_purposes` (`id`, `primary_purpose`, `node_id`, `document_id`) SELECT `id`, `primary_purpose`, `node_id`, `document_id` FROM `purposes`;
-- Drop "purposes" table after copying rows
DROP TABLE `purposes`;
-- Rename temporary table "new_purposes" to "purposes"
ALTER TABLE `new_purposes` RENAME TO `purposes`;
-- Create index "idx_purposes" to table: "purposes"
CREATE UNIQUE INDEX `idx_purposes` ON `purposes` (`node_id`, `primary_purpose`);
-- Create "new_tools" table
CREATE TABLE `new_tools` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `name` text NOT NULL,
  `version` text NOT NULL,
  `vendor` text NOT NULL,
  `metadata_id` uuid NOT NULL,
  `document_id` uuid NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `tools_metadata_tools` FOREIGN KEY (`metadata_id`) REFERENCES `metadata` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tools_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE
);
-- Copy rows from old table "tools" to new temporary table "new_tools"
INSERT INTO `new_tools` (`id`, `proto_message`, `name`, `version`, `vendor`, `metadata_id`, `document_id`) SELECT `id`, `proto_message`, `name`, `version`, `vendor`, `metadata_id`, `document_id` FROM `tools`;
-- Drop "tools" table after copying rows
DROP TABLE `tools`;
-- Rename temporary table "new_tools" to "tools"
ALTER TABLE `new_tools` RENAME TO `tools`;
-- Create index "tools_proto_message_key" to table: "tools"
CREATE UNIQUE INDEX `tools_proto_message_key` ON `tools` (`proto_message`);
-- Create index "idx_tools" to table: "tools"
CREATE UNIQUE INDEX `idx_tools` ON `tools` (`metadata_id`, `name`, `version`, `vendor`);
-- Create "new_properties" table
CREATE TABLE `new_properties` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `name` text NOT NULL,
  `data` text NOT NULL,
  `node_id` uuid NOT NULL,
  `document_id` uuid NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `properties_nodes_properties` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `properties_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE
);
-- Copy rows from old table "properties" to new temporary table "new_properties"
INSERT INTO `new_properties` (`id`, `proto_message`, `name`, `data`, `node_id`, `document_id`) SELECT `id`, `proto_message`, `name`, `data`, `node_id`, `document_id` FROM `properties`;
-- Drop "properties" table after copying rows
DROP TABLE `properties`;
-- Rename temporary table "new_properties" to "properties"
ALTER TABLE `new_properties` RENAME TO `properties`;
-- Create index "properties_proto_message_key" to table: "properties"
CREATE UNIQUE INDEX `properties_proto_message_key` ON `properties` (`proto_message`);
-- Create index "idx_property" to table: "properties"
CREATE UNIQUE INDEX `idx_property` ON `properties` (`name`, `data`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
