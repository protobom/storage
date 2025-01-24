-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_nodes" table
CREATE TABLE `new_nodes` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `native_id` text NOT NULL,
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
-- Copy rows from old table "nodes" to new temporary table "new_nodes"
INSERT INTO `new_nodes` (`id`, `proto_message`, `native_id`, `type`, `name`, `version`, `file_name`, `url_home`, `url_download`, `licenses`, `license_concluded`, `license_comments`, `copyright`, `source_info`, `comment`, `summary`, `description`, `release_date`, `build_date`, `valid_until_date`, `attribution`, `file_types`) SELECT `id`, `proto_message`, `native_id`, `type`, `name`, `version`, `file_name`, `url_home`, `url_download`, `licenses`, `license_concluded`, `license_comments`, `copyright`, `source_info`, `comment`, `summary`, `description`, `release_date`, `build_date`, `valid_until_date`, `attribution`, `file_types` FROM `nodes`;
-- Drop "nodes" table after copying rows
DROP TABLE `nodes`;
-- Rename temporary table "new_nodes" to "nodes"
ALTER TABLE `new_nodes` RENAME TO `nodes`;
-- Create index "nodes_proto_message_key" to table: "nodes"
CREATE UNIQUE INDEX `nodes_proto_message_key` ON `nodes` (`proto_message`);
-- Create "new_persons" table
CREATE TABLE `new_persons` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `name` text NOT NULL,
  `is_org` bool NOT NULL,
  `email` text NOT NULL,
  `url` text NOT NULL,
  `phone` text NOT NULL,
  PRIMARY KEY (`id`)
);
-- Copy rows from old table "persons" to new temporary table "new_persons"
INSERT INTO `new_persons` (`id`, `proto_message`, `name`, `is_org`, `email`, `url`, `phone`) SELECT `id`, `proto_message`, `name`, `is_org`, `email`, `url`, `phone` FROM `persons`;
-- Drop "persons" table after copying rows
DROP TABLE `persons`;
-- Rename temporary table "new_persons" to "persons"
ALTER TABLE `new_persons` RENAME TO `persons`;
-- Create index "persons_proto_message_key" to table: "persons"
CREATE UNIQUE INDEX `persons_proto_message_key` ON `persons` (`proto_message`);
-- Create index "idx_persons" to table: "persons"
CREATE UNIQUE INDEX `idx_persons` ON `persons` (`name`, `is_org`, `email`, `url`, `phone`);
-- Create "new_source_data" table
CREATE TABLE `new_source_data` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `format` text NOT NULL,
  `size` integer NOT NULL,
  `uri` text NULL,
  PRIMARY KEY (`id`)
);
-- Copy rows from old table "source_data" to new temporary table "new_source_data"
INSERT INTO `new_source_data` (`id`, `proto_message`, `format`, `size`, `uri`) SELECT `id`, `proto_message`, `format`, `size`, `uri` FROM `source_data`;
-- Drop "source_data" table after copying rows
DROP TABLE `source_data`;
-- Rename temporary table "new_source_data" to "source_data"
ALTER TABLE `new_source_data` RENAME TO `source_data`;
-- Create index "source_data_proto_message_key" to table: "source_data"
CREATE UNIQUE INDEX `source_data_proto_message_key` ON `source_data` (`proto_message`);
-- Create index "idx_source_data" to table: "source_data"
CREATE UNIQUE INDEX `idx_source_data` ON `source_data` (`format`, `size`, `uri`);
-- Create "new_edge_types" table
CREATE TABLE `new_edge_types` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `type` text NOT NULL,
  `node_id` uuid NOT NULL,
  `to_node_id` uuid NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `edge_types_nodes_from` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `edge_types_nodes_to` FOREIGN KEY (`to_node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE
);
-- Copy rows from old table "edge_types" to new temporary table "new_edge_types"
INSERT INTO `new_edge_types` (`id`, `proto_message`, `type`, `node_id`, `to_node_id`) SELECT `id`, `proto_message`, `type`, `node_id`, `to_node_id` FROM `edge_types`;
-- Drop "edge_types" table after copying rows
DROP TABLE `edge_types`;
-- Rename temporary table "new_edge_types" to "edge_types"
ALTER TABLE `new_edge_types` RENAME TO `edge_types`;
-- Create index "edge_types_proto_message_key" to table: "edge_types"
CREATE UNIQUE INDEX `edge_types_proto_message_key` ON `edge_types` (`proto_message`);
-- Create index "idx_edge_types" to table: "edge_types"
CREATE UNIQUE INDEX `idx_edge_types` ON `edge_types` (`type`, `node_id`, `to_node_id`);
-- Create index "edgetype_node_id_to_node_id" to table: "edge_types"
CREATE UNIQUE INDEX `edgetype_node_id_to_node_id` ON `edge_types` (`node_id`, `to_node_id`);
-- Create "new_hashes_entries" table
CREATE TABLE `new_hashes_entries` (
  `id` uuid NOT NULL,
  `hash_algorithm` text NOT NULL,
  `hash_data` text NOT NULL,
  PRIMARY KEY (`id`)
);
-- Copy rows from old table "hashes_entries" to new temporary table "new_hashes_entries"
INSERT INTO `new_hashes_entries` (`id`, `hash_algorithm`, `hash_data`) SELECT `id`, `hash_algorithm`, `hash_data` FROM `hashes_entries`;
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
  PRIMARY KEY (`id`)
);
-- Copy rows from old table "external_references" to new temporary table "new_external_references"
INSERT INTO `new_external_references` (`id`, `proto_message`, `url`, `comment`, `authority`, `type`) SELECT `id`, `proto_message`, `url`, `comment`, `authority`, `type` FROM `external_references`;
-- Drop "external_references" table after copying rows
DROP TABLE `external_references`;
-- Rename temporary table "new_external_references" to "external_references"
ALTER TABLE `new_external_references` RENAME TO `external_references`;
-- Create index "external_references_proto_message_key" to table: "external_references"
CREATE UNIQUE INDEX `external_references_proto_message_key` ON `external_references` (`proto_message`);
-- Create "new_identifiers_entries" table
CREATE TABLE `new_identifiers_entries` (
  `id` uuid NOT NULL,
  `type` text NOT NULL,
  `value` text NOT NULL,
  PRIMARY KEY (`id`)
);
-- Copy rows from old table "identifiers_entries" to new temporary table "new_identifiers_entries"
INSERT INTO `new_identifiers_entries` (`id`, `type`, `value`) SELECT `id`, `type`, `value` FROM `identifiers_entries`;
-- Drop "identifiers_entries" table after copying rows
DROP TABLE `identifiers_entries`;
-- Rename temporary table "new_identifiers_entries" to "identifiers_entries"
ALTER TABLE `new_identifiers_entries` RENAME TO `identifiers_entries`;
-- Create index "idx_identifiers" to table: "identifiers_entries"
CREATE UNIQUE INDEX `idx_identifiers` ON `identifiers_entries` (`type`, `value`);
-- Create "new_metadata" table
CREATE TABLE `new_metadata` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `native_id` text NOT NULL,
  `version` text NOT NULL,
  `name` text NOT NULL,
  `date` datetime NOT NULL,
  `comment` text NOT NULL,
  `source_data_id` uuid NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `metadata_source_data_source_data` FOREIGN KEY (`source_data_id`) REFERENCES `source_data` (`id`) ON DELETE CASCADE
);
-- Copy rows from old table "metadata" to new temporary table "new_metadata"
INSERT INTO `new_metadata` (`id`, `proto_message`, `native_id`, `version`, `name`, `date`, `comment`) SELECT `id`, `proto_message`, `native_id`, `version`, `name`, `date`, `comment` FROM `metadata`;
-- Drop "metadata" table after copying rows
DROP TABLE `metadata`;
-- Rename temporary table "new_metadata" to "metadata"
ALTER TABLE `new_metadata` RENAME TO `metadata`;
-- Create index "metadata_proto_message_key" to table: "metadata"
CREATE UNIQUE INDEX `metadata_proto_message_key` ON `metadata` (`proto_message`);
-- Create "new_documents" table
CREATE TABLE `new_documents` (
  `id` uuid NOT NULL,
  `metadata_id` uuid NULL,
  `node_list_id` uuid NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `documents_metadata_metadata` FOREIGN KEY (`metadata_id`) REFERENCES `metadata` (`id`) ON DELETE CASCADE,
  CONSTRAINT `documents_node_lists_node_list` FOREIGN KEY (`node_list_id`) REFERENCES `node_lists` (`id`) ON DELETE CASCADE
);
-- Copy rows from old table "documents" to new temporary table "new_documents"
INSERT INTO `new_documents` (`id`, `metadata_id`, `node_list_id`) SELECT `id`, `metadata_id`, `node_list_id` FROM `documents`;
-- Drop "documents" table after copying rows
DROP TABLE `documents`;
-- Rename temporary table "new_documents" to "documents"
ALTER TABLE `new_documents` RENAME TO `documents`;
-- Create index "idx_documents_metadata_id" to table: "documents"
CREATE INDEX `idx_documents_metadata_id` ON `documents` (`metadata_id`);
-- Create index "idx_documents_node_list_id" to table: "documents"
CREATE INDEX `idx_documents_node_list_id` ON `documents` (`node_list_id`);
-- Create "new_document_types" table
CREATE TABLE `new_document_types` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `type` text NULL,
  `name` text NULL,
  `description` text NULL,
  PRIMARY KEY (`id`)
);
-- Copy rows from old table "document_types" to new temporary table "new_document_types"
INSERT INTO `new_document_types` (`id`, `proto_message`, `type`, `name`, `description`) SELECT `id`, `proto_message`, `type`, `name`, `description` FROM `document_types`;
-- Drop "document_types" table after copying rows
DROP TABLE `document_types`;
-- Rename temporary table "new_document_types" to "document_types"
ALTER TABLE `new_document_types` RENAME TO `document_types`;
-- Create index "document_types_proto_message_key" to table: "document_types"
CREATE UNIQUE INDEX `document_types_proto_message_key` ON `document_types` (`proto_message`);
-- Create index "idx_document_types" to table: "document_types"
CREATE UNIQUE INDEX `idx_document_types` ON `document_types` (`type`, `name`, `description`);
-- Create "new_purposes" table
CREATE TABLE `new_purposes` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `primary_purpose` text NOT NULL
);
-- Copy rows from old table "purposes" to new temporary table "new_purposes"
INSERT INTO `new_purposes` (`id`, `primary_purpose`) SELECT `id`, `primary_purpose` FROM `purposes`;
-- Drop "purposes" table after copying rows
DROP TABLE `purposes`;
-- Rename temporary table "new_purposes" to "purposes"
ALTER TABLE `new_purposes` RENAME TO `purposes`;
-- Create "new_tools" table
CREATE TABLE `new_tools` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `name` text NOT NULL,
  `version` text NOT NULL,
  `vendor` text NOT NULL,
  PRIMARY KEY (`id`)
);
-- Copy rows from old table "tools" to new temporary table "new_tools"
INSERT INTO `new_tools` (`id`, `proto_message`, `name`, `version`, `vendor`) SELECT `id`, `proto_message`, `name`, `version`, `vendor` FROM `tools`;
-- Drop "tools" table after copying rows
DROP TABLE `tools`;
-- Rename temporary table "new_tools" to "tools"
ALTER TABLE `new_tools` RENAME TO `tools`;
-- Create index "tools_proto_message_key" to table: "tools"
CREATE UNIQUE INDEX `tools_proto_message_key` ON `tools` (`proto_message`);
-- Create index "idx_tools" to table: "tools"
CREATE UNIQUE INDEX `idx_tools` ON `tools` (`name`, `version`, `vendor`);
-- Create "new_properties" table
CREATE TABLE `new_properties` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `name` text NOT NULL,
  `data` text NOT NULL,
  PRIMARY KEY (`id`)
);
-- Copy rows from old table "properties" to new temporary table "new_properties"
INSERT INTO `new_properties` (`id`, `proto_message`, `name`, `data`) SELECT `id`, `proto_message`, `name`, `data` FROM `properties`;
-- Drop "properties" table after copying rows
DROP TABLE `properties`;
-- Rename temporary table "new_properties" to "properties"
ALTER TABLE `new_properties` RENAME TO `properties`;
-- Create index "properties_proto_message_key" to table: "properties"
CREATE UNIQUE INDEX `properties_proto_message_key` ON `properties` (`proto_message`);
-- Create index "idx_property" to table: "properties"
CREATE UNIQUE INDEX `idx_property` ON `properties` (`name`, `data`);
-- Create "document_document_types" table
CREATE TABLE `document_document_types` (
  `document_id` uuid NOT NULL,
  `document_type_id` uuid NOT NULL,
  PRIMARY KEY (`document_id`, `document_type_id`),
  CONSTRAINT `document_document_types_document_id` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE,
  CONSTRAINT `document_document_types_document_type_id` FOREIGN KEY (`document_type_id`) REFERENCES `document_types` (`id`) ON DELETE CASCADE
);
-- Create "document_edge_types" table
CREATE TABLE `document_edge_types` (
  `document_id` uuid NOT NULL,
  `edge_type_id` uuid NOT NULL,
  PRIMARY KEY (`document_id`, `edge_type_id`),
  CONSTRAINT `document_edge_types_document_id` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE,
  CONSTRAINT `document_edge_types_edge_type_id` FOREIGN KEY (`edge_type_id`) REFERENCES `edge_types` (`id`) ON DELETE CASCADE
);
-- Create "document_external_references" table
CREATE TABLE `document_external_references` (
  `document_id` uuid NOT NULL,
  `external_reference_id` uuid NOT NULL,
  PRIMARY KEY (`document_id`, `external_reference_id`),
  CONSTRAINT `document_external_references_document_id` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE,
  CONSTRAINT `document_external_references_external_reference_id` FOREIGN KEY (`external_reference_id`) REFERENCES `external_references` (`id`) ON DELETE CASCADE
);
-- Create "document_hashes" table
CREATE TABLE `document_hashes` (
  `document_id` uuid NOT NULL,
  `hashes_entry_id` uuid NOT NULL,
  PRIMARY KEY (`document_id`, `hashes_entry_id`),
  CONSTRAINT `document_hashes_document_id` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE,
  CONSTRAINT `document_hashes_hashes_entry_id` FOREIGN KEY (`hashes_entry_id`) REFERENCES `hashes_entries` (`id`) ON DELETE CASCADE
);
-- Create "document_identifiers" table
CREATE TABLE `document_identifiers` (
  `document_id` uuid NOT NULL,
  `identifiers_entry_id` uuid NOT NULL,
  PRIMARY KEY (`document_id`, `identifiers_entry_id`),
  CONSTRAINT `document_identifiers_document_id` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE,
  CONSTRAINT `document_identifiers_identifiers_entry_id` FOREIGN KEY (`identifiers_entry_id`) REFERENCES `identifiers_entries` (`id`) ON DELETE CASCADE
);
-- Create "document_nodes" table
CREATE TABLE `document_nodes` (
  `document_id` uuid NOT NULL,
  `node_id` uuid NOT NULL,
  PRIMARY KEY (`document_id`, `node_id`),
  CONSTRAINT `document_nodes_document_id` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE,
  CONSTRAINT `document_nodes_node_id` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE
);
-- Create "document_persons" table
CREATE TABLE `document_persons` (
  `document_id` uuid NOT NULL,
  `person_id` uuid NOT NULL,
  PRIMARY KEY (`document_id`, `person_id`),
  CONSTRAINT `document_persons_document_id` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE,
  CONSTRAINT `document_persons_person_id` FOREIGN KEY (`person_id`) REFERENCES `persons` (`id`) ON DELETE CASCADE
);
-- Create "document_properties" table
CREATE TABLE `document_properties` (
  `document_id` uuid NOT NULL,
  `property_id` uuid NOT NULL,
  PRIMARY KEY (`document_id`, `property_id`),
  CONSTRAINT `document_properties_document_id` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE,
  CONSTRAINT `document_properties_property_id` FOREIGN KEY (`property_id`) REFERENCES `properties` (`id`) ON DELETE CASCADE
);
-- Create "document_purposes" table
CREATE TABLE `document_purposes` (
  `document_id` uuid NOT NULL,
  `purpose_id` integer NOT NULL,
  PRIMARY KEY (`document_id`, `purpose_id`),
  CONSTRAINT `document_purposes_document_id` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE,
  CONSTRAINT `document_purposes_purpose_id` FOREIGN KEY (`purpose_id`) REFERENCES `purposes` (`id`) ON DELETE CASCADE
);
-- Create "document_source_data" table
CREATE TABLE `document_source_data` (
  `document_id` uuid NOT NULL,
  `source_data_id` uuid NOT NULL,
  PRIMARY KEY (`document_id`, `source_data_id`),
  CONSTRAINT `document_source_data_document_id` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE,
  CONSTRAINT `document_source_data_source_data_id` FOREIGN KEY (`source_data_id`) REFERENCES `source_data` (`id`) ON DELETE CASCADE
);
-- Create "document_tools" table
CREATE TABLE `document_tools` (
  `document_id` uuid NOT NULL,
  `tool_id` uuid NOT NULL,
  PRIMARY KEY (`document_id`, `tool_id`),
  CONSTRAINT `document_tools_document_id` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE,
  CONSTRAINT `document_tools_tool_id` FOREIGN KEY (`tool_id`) REFERENCES `tools` (`id`) ON DELETE CASCADE
);
-- Create "metadata_tools" table
CREATE TABLE `metadata_tools` (
  `metadata_id` uuid NOT NULL,
  `tool_id` uuid NOT NULL,
  PRIMARY KEY (`metadata_id`, `tool_id`),
  CONSTRAINT `metadata_tools_metadata_id` FOREIGN KEY (`metadata_id`) REFERENCES `metadata` (`id`) ON DELETE CASCADE,
  CONSTRAINT `metadata_tools_tool_id` FOREIGN KEY (`tool_id`) REFERENCES `tools` (`id`) ON DELETE CASCADE
);
-- Create "metadata_authors" table
CREATE TABLE `metadata_authors` (
  `metadata_id` uuid NOT NULL,
  `person_id` uuid NOT NULL,
  PRIMARY KEY (`metadata_id`, `person_id`),
  CONSTRAINT `metadata_authors_metadata_id` FOREIGN KEY (`metadata_id`) REFERENCES `metadata` (`id`) ON DELETE CASCADE,
  CONSTRAINT `metadata_authors_person_id` FOREIGN KEY (`person_id`) REFERENCES `persons` (`id`) ON DELETE CASCADE
);
-- Create "metadata_document_types" table
CREATE TABLE `metadata_document_types` (
  `metadata_id` uuid NOT NULL,
  `document_type_id` uuid NOT NULL,
  PRIMARY KEY (`metadata_id`, `document_type_id`),
  CONSTRAINT `metadata_document_types_metadata_id` FOREIGN KEY (`metadata_id`) REFERENCES `metadata` (`id`) ON DELETE CASCADE,
  CONSTRAINT `metadata_document_types_document_type_id` FOREIGN KEY (`document_type_id`) REFERENCES `document_types` (`id`) ON DELETE CASCADE
);
-- Create "node_suppliers" table
CREATE TABLE `node_suppliers` (
  `node_id` uuid NOT NULL,
  `person_id` uuid NOT NULL,
  PRIMARY KEY (`node_id`, `person_id`),
  CONSTRAINT `node_suppliers_node_id` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `node_suppliers_person_id` FOREIGN KEY (`person_id`) REFERENCES `persons` (`id`) ON DELETE CASCADE
);
-- Create "node_originators" table
CREATE TABLE `node_originators` (
  `node_id` uuid NOT NULL,
  `person_id` uuid NOT NULL,
  PRIMARY KEY (`node_id`, `person_id`),
  CONSTRAINT `node_originators_node_id` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `node_originators_person_id` FOREIGN KEY (`person_id`) REFERENCES `persons` (`id`) ON DELETE CASCADE
);
-- Create "node_primary_purposes" table
CREATE TABLE `node_primary_purposes` (
  `node_id` uuid NOT NULL,
  `purpose_id` integer NOT NULL,
  PRIMARY KEY (`node_id`, `purpose_id`),
  CONSTRAINT `node_primary_purposes_node_id` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `node_primary_purposes_purpose_id` FOREIGN KEY (`purpose_id`) REFERENCES `purposes` (`id`) ON DELETE CASCADE
);
-- Create "node_properties" table
CREATE TABLE `node_properties` (
  `node_id` uuid NOT NULL,
  `property_id` uuid NOT NULL,
  PRIMARY KEY (`node_id`, `property_id`),
  CONSTRAINT `node_properties_node_id` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `node_properties_property_id` FOREIGN KEY (`property_id`) REFERENCES `properties` (`id`) ON DELETE CASCADE
);
-- Create "person_contacts" table
CREATE TABLE `person_contacts` (
  `person_id` uuid NOT NULL,
  `contact_owner_id` uuid NOT NULL,
  PRIMARY KEY (`person_id`, `contact_owner_id`),
  CONSTRAINT `person_contacts_person_id` FOREIGN KEY (`person_id`) REFERENCES `persons` (`id`) ON DELETE CASCADE,
  CONSTRAINT `person_contacts_contact_owner_id` FOREIGN KEY (`contact_owner_id`) REFERENCES `persons` (`id`) ON DELETE CASCADE
);
-- Create "source_data_hashes" table
CREATE TABLE `source_data_hashes` (
  `source_data_id` uuid NOT NULL,
  `hash_entry_id` uuid NOT NULL,
  PRIMARY KEY (`source_data_id`, `hash_entry_id`),
  CONSTRAINT `source_data_hashes_source_data_id` FOREIGN KEY (`source_data_id`) REFERENCES `source_data` (`id`) ON DELETE CASCADE,
  CONSTRAINT `source_data_hashes_hash_entry_id` FOREIGN KEY (`hash_entry_id`) REFERENCES `hashes_entries` (`id`) ON DELETE CASCADE
);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
