// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent

import (
	"github.com/google/uuid"
	"github.com/protobom/storage/internal/backends/ent/annotation"
	"github.com/protobom/storage/internal/backends/ent/documenttype"
	"github.com/protobom/storage/internal/backends/ent/edgetype"
	"github.com/protobom/storage/internal/backends/ent/externalreference"
	"github.com/protobom/storage/internal/backends/ent/metadata"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/person"
	"github.com/protobom/storage/internal/backends/ent/property"
	"github.com/protobom/storage/internal/backends/ent/purpose"
	"github.com/protobom/storage/internal/backends/ent/schema"
	"github.com/protobom/storage/internal/backends/ent/tool"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	annotationMixin := schema.Annotation{}.Mixin()
	annotationMixinFields0 := annotationMixin[0].Fields()
	_ = annotationMixinFields0
	annotationFields := schema.Annotation{}.Fields()
	_ = annotationFields
	// annotationDescDocumentID is the schema descriptor for document_id field.
	annotationDescDocumentID := annotationMixinFields0[0].Descriptor()
	// annotation.DefaultDocumentID holds the default value on creation for the document_id field.
	annotation.DefaultDocumentID = annotationDescDocumentID.Default.(func() uuid.UUID)
	// annotationDescIsUnique is the schema descriptor for is_unique field.
	annotationDescIsUnique := annotationFields[2].Descriptor()
	// annotation.DefaultIsUnique holds the default value on creation for the is_unique field.
	annotation.DefaultIsUnique = annotationDescIsUnique.Default.(bool)
	documenttypeMixin := schema.DocumentType{}.Mixin()
	documenttypeMixinFields0 := documenttypeMixin[0].Fields()
	_ = documenttypeMixinFields0
	documenttypeFields := schema.DocumentType{}.Fields()
	_ = documenttypeFields
	// documenttypeDescDocumentID is the schema descriptor for document_id field.
	documenttypeDescDocumentID := documenttypeMixinFields0[0].Descriptor()
	// documenttype.DefaultDocumentID holds the default value on creation for the document_id field.
	documenttype.DefaultDocumentID = documenttypeDescDocumentID.Default.(func() uuid.UUID)
	edgetypeMixin := schema.EdgeType{}.Mixin()
	edgetypeMixinFields0 := edgetypeMixin[0].Fields()
	_ = edgetypeMixinFields0
	edgetypeFields := schema.EdgeType{}.Fields()
	_ = edgetypeFields
	// edgetypeDescDocumentID is the schema descriptor for document_id field.
	edgetypeDescDocumentID := edgetypeMixinFields0[0].Descriptor()
	// edgetype.DefaultDocumentID holds the default value on creation for the document_id field.
	edgetype.DefaultDocumentID = edgetypeDescDocumentID.Default.(func() uuid.UUID)
	externalreferenceMixin := schema.ExternalReference{}.Mixin()
	externalreferenceMixinFields0 := externalreferenceMixin[0].Fields()
	_ = externalreferenceMixinFields0
	externalreferenceFields := schema.ExternalReference{}.Fields()
	_ = externalreferenceFields
	// externalreferenceDescDocumentID is the schema descriptor for document_id field.
	externalreferenceDescDocumentID := externalreferenceMixinFields0[0].Descriptor()
	// externalreference.DefaultDocumentID holds the default value on creation for the document_id field.
	externalreference.DefaultDocumentID = externalreferenceDescDocumentID.Default.(func() uuid.UUID)
	metadataFields := schema.Metadata{}.Fields()
	_ = metadataFields
	// metadataDescID is the schema descriptor for id field.
	metadataDescID := metadataFields[0].Descriptor()
	// metadata.IDValidator is a validator for the "id" field. It is called by the builders before save.
	metadata.IDValidator = metadataDescID.Validators[0].(func(string) error)
	nodeMixin := schema.Node{}.Mixin()
	nodeMixinFields0 := nodeMixin[0].Fields()
	_ = nodeMixinFields0
	nodeFields := schema.Node{}.Fields()
	_ = nodeFields
	// nodeDescDocumentID is the schema descriptor for document_id field.
	nodeDescDocumentID := nodeMixinFields0[0].Descriptor()
	// node.DefaultDocumentID holds the default value on creation for the document_id field.
	node.DefaultDocumentID = nodeDescDocumentID.Default.(func() uuid.UUID)
	// nodeDescID is the schema descriptor for id field.
	nodeDescID := nodeFields[0].Descriptor()
	// node.IDValidator is a validator for the "id" field. It is called by the builders before save.
	node.IDValidator = nodeDescID.Validators[0].(func(string) error)
	personMixin := schema.Person{}.Mixin()
	personMixinFields0 := personMixin[0].Fields()
	_ = personMixinFields0
	personFields := schema.Person{}.Fields()
	_ = personFields
	// personDescDocumentID is the schema descriptor for document_id field.
	personDescDocumentID := personMixinFields0[0].Descriptor()
	// person.DefaultDocumentID holds the default value on creation for the document_id field.
	person.DefaultDocumentID = personDescDocumentID.Default.(func() uuid.UUID)
	propertyMixin := schema.Property{}.Mixin()
	propertyMixinFields0 := propertyMixin[0].Fields()
	_ = propertyMixinFields0
	propertyFields := schema.Property{}.Fields()
	_ = propertyFields
	// propertyDescDocumentID is the schema descriptor for document_id field.
	propertyDescDocumentID := propertyMixinFields0[0].Descriptor()
	// property.DefaultDocumentID holds the default value on creation for the document_id field.
	property.DefaultDocumentID = propertyDescDocumentID.Default.(func() uuid.UUID)
	purposeMixin := schema.Purpose{}.Mixin()
	purposeMixinFields0 := purposeMixin[0].Fields()
	_ = purposeMixinFields0
	purposeFields := schema.Purpose{}.Fields()
	_ = purposeFields
	// purposeDescDocumentID is the schema descriptor for document_id field.
	purposeDescDocumentID := purposeMixinFields0[0].Descriptor()
	// purpose.DefaultDocumentID holds the default value on creation for the document_id field.
	purpose.DefaultDocumentID = purposeDescDocumentID.Default.(func() uuid.UUID)
	toolMixin := schema.Tool{}.Mixin()
	toolMixinFields0 := toolMixin[0].Fields()
	_ = toolMixinFields0
	toolFields := schema.Tool{}.Fields()
	_ = toolFields
	// toolDescDocumentID is the schema descriptor for document_id field.
	toolDescDocumentID := toolMixinFields0[0].Descriptor()
	// tool.DefaultDocumentID holds the default value on creation for the document_id field.
	tool.DefaultDocumentID = toolDescDocumentID.Default.(func() uuid.UUID)
}
