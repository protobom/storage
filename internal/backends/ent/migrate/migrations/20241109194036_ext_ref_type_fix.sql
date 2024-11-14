-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_annotations" table
CREATE TABLE `new_annotations` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `name` text NOT NULL,
  `value` text NOT NULL,
  `is_unique` bool NOT NULL DEFAULT (false),
  `document_id` uuid NULL,
  CONSTRAINT `annotations_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE SET NULL
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
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
