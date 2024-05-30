// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"fmt"

	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"
)

// Store implements the storage.Storer interface.
func (backend *Backend) Store(_document *sbom.Document, _opts *storage.StoreOptions) error {
	if backend.BackendOptions == nil {
		backend.BackendOptions = NewBackendOptions()
	}

	backend.init(backend.BackendOptions)

	defer backend.client.Close()

	if err := backend.client.Document.Create().Exec(backend.ctx); err != nil {
		return fmt.Errorf("failed to store *sbom.Document: %w", err)
	}

	return nil
}

func (backend *Backend) StoreDocumentTypes(_docTypes ...*sbom.DocumentType) error {
	return nil
}

func (backend *Backend) StoreExternalReferences(_refs ...*sbom.ExternalReference) error {
	return nil
}

func (backend *Backend) StoreHashesEntries(_hashes map[int32]string) error {
	return nil
}

func (backend *Backend) StoreIdentifiersEntries(_idents map[int32]string) error {
	return nil
}

func (backend *Backend) StoreMetadata(_metadata *sbom.Metadata) error {
	return nil
}

func (backend *Backend) StoreNodeList(_nodeList *sbom.NodeList) error {
	return nil
}

func (backend *Backend) StoreNodes(_nodes ...*sbom.Node) error {
	return nil
}

func (backend *Backend) StorePersons(_persons ...*sbom.Person) error {
	return nil
}

func (backend *Backend) StorePurposes(_purposes ...sbom.Purpose) error {
	return nil
}

func (backend *Backend) StoreTools(_tools ...*sbom.Tool) error {
	return nil
}
