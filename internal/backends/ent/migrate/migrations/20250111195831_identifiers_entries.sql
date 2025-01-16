-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_identifiers_entries" table
CREATE TABLE `new_identifiers_entries` (
  `id` uuid NOT NULL,
  `type` text NOT NULL,
  `value` text NOT NULL,
  `document_id` uuid NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `identifiers_entries_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE
);
-- Copy rows from old table "identifiers_entries" to new temporary table "new_identifiers_entries"
INSERT INTO `new_identifiers_entries` (`id`) SELECT `id` FROM `identifiers_entries`;
-- Drop "identifiers_entries" table after copying rows
DROP TABLE `identifiers_entries`;
-- Rename temporary table "new_identifiers_entries" to "identifiers_entries"
ALTER TABLE `new_identifiers_entries` RENAME TO `identifiers_entries`;
-- Create index "idx_identifiers" to table: "identifiers_entries"
CREATE UNIQUE INDEX `idx_identifiers` ON `identifiers_entries` (`type`, `value`);
-- Create "node_identifiers" table
CREATE TABLE `node_identifiers` (
  `node_id` uuid NOT NULL,
  `identifier_entry_id` uuid NOT NULL,
  PRIMARY KEY (`node_id`, `identifier_entry_id`),
  CONSTRAINT `node_identifiers_node_id` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `node_identifiers_identifier_entry_id` FOREIGN KEY (`identifier_entry_id`) REFERENCES `identifiers_entries` (`id`) ON DELETE CASCADE
);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
