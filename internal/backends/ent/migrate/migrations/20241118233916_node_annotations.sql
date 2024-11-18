-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_annotations" table
CREATE TABLE `new_annotations` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `name` text NOT NULL,
  `value` text NOT NULL,
  `is_unique` bool NOT NULL DEFAULT (false),
  `document_id` uuid NULL,
  `node_id` text NULL,
  CONSTRAINT `annotations_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE,
  CONSTRAINT `annotations_nodes_annotations` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE
);
-- Copy rows from old table "annotations" to new temporary table "new_annotations"
INSERT INTO `new_annotations` (`id`, `name`, `value`, `is_unique`, `document_id`) SELECT `id`, `name`, `value`, `is_unique`, `document_id` FROM `annotations`;
-- Drop "annotations" table after copying rows
DROP TABLE `annotations`;
-- Rename temporary table "new_annotations" to "annotations"
ALTER TABLE `new_annotations` RENAME TO `annotations`;
-- Create index "idx_node_annotations" to table: "annotations"
CREATE UNIQUE INDEX `idx_node_annotations` ON `annotations` (`node_id`, `name`, `value`) WHERE node_id IS NOT NULL AND TRIM(node_id) != '';
-- Create index "idx_node_unique_annotations" to table: "annotations"
CREATE UNIQUE INDEX `idx_node_unique_annotations` ON `annotations` (`node_id`, `name`) WHERE node_id IS NOT NULL AND TRIM(node_id) != '' AND is_unique;
-- Create index "idx_document_annotations" to table: "annotations"
CREATE UNIQUE INDEX `idx_document_annotations` ON `annotations` (`document_id`, `name`, `value`) WHERE document_id IS NOT NULL AND TRIM(document_id) != '';
-- Create index "idx_document_unique_annotations" to table: "annotations"
CREATE UNIQUE INDEX `idx_document_unique_annotations` ON `annotations` (`document_id`, `name`) WHERE document_id IS NOT NULL AND TRIM(document_id) != '' AND is_unique;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
