-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_edge_types" table
CREATE TABLE `new_edge_types` (
  `id` uuid NOT NULL,
  `proto_message` blob NOT NULL,
  `type` text NOT NULL,
  `document_id` uuid NULL,
  `node_id` uuid NOT NULL,
  `to_node_id` uuid NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `edge_types_documents_document` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE,
  CONSTRAINT `edge_types_nodes_from` FOREIGN KEY (`node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `edge_types_nodes_to` FOREIGN KEY (`to_node_id`) REFERENCES `nodes` (`id`) ON DELETE CASCADE
);
-- Copy rows from old table "edge_types" to new temporary table "new_edge_types"
INSERT INTO `new_edge_types` (`id`, `type`, `document_id`, `node_id`, `to_node_id`) SELECT `id`, `type`, `document_id`, `node_id`, `to_node_id` FROM `edge_types`;
-- Drop "edge_types" table after copying rows
DROP TABLE `edge_types`;
-- Rename temporary table "new_edge_types" to "edge_types"
ALTER TABLE `new_edge_types` RENAME TO `edge_types`;
-- Create index "idx_edge_types" to table: "edge_types"
CREATE UNIQUE INDEX `idx_edge_types` ON `edge_types` (`type`, `node_id`, `to_node_id`);
-- Create index "edgetype_node_id_to_node_id" to table: "edge_types"
CREATE UNIQUE INDEX `edgetype_node_id_to_node_id` ON `edge_types` (`node_id`, `to_node_id`);
-- Create "node_list_edges" table
CREATE TABLE `node_list_edges` (
  `node_list_id` uuid NOT NULL,
  `edge_type_id` uuid NOT NULL,
  PRIMARY KEY (`node_list_id`, `edge_type_id`),
  CONSTRAINT `node_list_edges_node_list_id` FOREIGN KEY (`node_list_id`) REFERENCES `node_lists` (`id`) ON DELETE CASCADE,
  CONSTRAINT `node_list_edges_edge_type_id` FOREIGN KEY (`edge_type_id`) REFERENCES `edge_types` (`id`) ON DELETE CASCADE
);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
