// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"github.com/protobom/storage/internal/backends/ent/annotation"
	"github.com/protobom/storage/internal/backends/ent/metadata"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	annotationFields := schema.Annotation{}.Fields()
	_ = annotationFields
	// annotationDescIsUnique is the schema descriptor for is_unique field.
	annotationDescIsUnique := annotationFields[3].Descriptor()
	// annotation.DefaultIsUnique holds the default value on creation for the is_unique field.
	annotation.DefaultIsUnique = annotationDescIsUnique.Default.(bool)
	metadataFields := schema.Metadata{}.Fields()
	_ = metadataFields
	// metadataDescID is the schema descriptor for id field.
	metadataDescID := metadataFields[0].Descriptor()
	// metadata.IDValidator is a validator for the "id" field. It is called by the builders before save.
	metadata.IDValidator = metadataDescID.Validators[0].(func(string) error)
	nodeFields := schema.Node{}.Fields()
	_ = nodeFields
	// nodeDescID is the schema descriptor for id field.
	nodeDescID := nodeFields[0].Descriptor()
	// node.IDValidator is a validator for the "id" field. It is called by the builders before save.
	node.IDValidator = nodeDescID.Validators[0].(func(string) error)
}
