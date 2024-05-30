// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"

	"github.com/protobom/storage/internal/backends/ent"
)

// Retrieve implements the storage.Retriever interface.
func (backend *Backend) Retrieve(_id string, _opts *storage.RetrieveOptions) (*sbom.Document, error) {
	if backend.BackendOptions == nil {
		backend.BackendOptions = NewBackendOptions()
	}

	backend.init(backend.BackendOptions)

	defer backend.client.Close()

	return nil, nil
}

func (backend *Backend) RetrieveDocumentTypes(_docTypes ...*ent.DocumentType) ([]*sbom.DocumentType, error) {
	return nil, nil
}

func (backend *Backend) RetrieveExternalReferences(_refs ...*ent.ExternalReference) ([]*sbom.ExternalReference, error) {
	return nil, nil
}

func (backend *Backend) RetrieveHashesEntries(_hashes *ent.HashesEntries) (map[int32]string, error) {
	return nil, nil
}

func (backend *Backend) RetrieveIdentifiersEntries(_idents *ent.IdentifiersEntries) (map[int32]string, error) {
	return nil, nil
}

func (backend *Backend) RetrieveMetadata(_metadata *ent.Metadata) (*sbom.Metadata, error) {
	return nil, nil
}

func (backend *Backend) RetrieveNodeList(_nodeList *ent.NodeList) (*sbom.NodeList, error) {
	return nil, nil
}

func (backend *Backend) RetrieveNodes(_nodes ...*ent.Node) ([]*sbom.Node, error) {
	return nil, nil
}

func (backend *Backend) RetrievePersons(_persons ...*ent.Person) ([]*sbom.Person, error) {
	return nil, nil
}

func (backend *Backend) RetrievePurposes(_purposes ...*ent.Purpose) ([]sbom.Purpose, error) {
	return nil, nil
}

func (backend *Backend) RetrieveTools(_tools ...*ent.Tool) ([]*sbom.Tool, error) {
	return nil, nil
}
