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

	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/documenttype"
	"github.com/protobom/storage/internal/backends/ent/externalreference"
	"github.com/protobom/storage/internal/backends/ent/hashesentry"
	"github.com/protobom/storage/internal/backends/ent/identifiersentry"
	"github.com/protobom/storage/internal/backends/ent/metadata"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/nodelist"
	"github.com/protobom/storage/internal/backends/ent/purpose"
)

// Store implements the storage.Storer interface.
func (backend *Backend) Store(doc *sbom.Document, opts *storage.StoreOptions) error {
	if backend.BackendOptions == nil {
		backend.BackendOptions = NewBackendOptions()
	}

	backend.init(backend.BackendOptions)

	defer backend.client.Close()

	if opts == nil {
		opts = &storage.StoreOptions{
			BackendOptions: backend.BackendOptions,
		}
	}

	if _, ok := opts.BackendOptions.(*BackendOptions); !ok {
		return fmt.Errorf("%w", errInvalidEntOptions)
	}

	if err := backend.StoreMetadata(doc.Metadata); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := backend.StoreNodeList(doc.NodeList); err != nil {
		return fmt.Errorf("%w", err)
	}

	entNodeList, err := backend.client.NodeList.Query().
		Where(
			nodelist.Or(
				nodelist.HasDocumentWith(document.HasMetadataWith(metadata.IDEQ(doc.Metadata.Id))),
				nodelist.Not(nodelist.HasDocument()),
			)).
		Only(backend.ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = backend.client.Document.Create().
		SetMetadataID(doc.Metadata.Id).
		SetNodeListID(entNodeList.ID).
		OnConflict().
		Ignore().
		Exec(backend.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.Document: %w", err)
	}

	return nil
}

func (backend *Backend) StoreDocumentTypes(docTypes []*sbom.DocumentType) error {
	// Set up client if this method was called directly.
	if backend.BackendOptions == nil {
		backend.BackendOptions = NewBackendOptions()
		defer backend.client.Close()
	}

	backend.init(backend.BackendOptions)

	for _, dt := range docTypes {
		typeName := documenttype.Type(dt.String())

		newDocType := backend.client.DocumentType.Create().
			SetNillableType(&typeName).
			SetNillableName(dt.Name).
			SetNillableDescription(dt.Description)

		if backend.metadataID != "" {
			newDocType.SetMetadataID(backend.metadataID)
		}

		err := newDocType.OnConflict().Ignore().Exec(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.DocumentType: %w", err)
		}
	}

	return nil
}

func (backend *Backend) StoreExternalReferences(refs []*sbom.ExternalReference) error {
	// Set up client if this method was called directly.
	if backend.BackendOptions == nil {
		backend.BackendOptions = NewBackendOptions()
		defer backend.client.Close()
	}

	backend.init(backend.BackendOptions)

	for _, ref := range refs {
		newRef := backend.client.ExternalReference.Create().
			SetURL(ref.Url).
			SetComment(ref.Comment).
			SetAuthority(ref.Authority).
			SetType(externalreference.Type(ref.Type.String()))

		if backend.nodeID != "" {
			newRef.SetNodeID(backend.nodeID)
		}

		id, err := newRef.OnConflict().Ignore().ID(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.ExternalReference: %w", err)
		}

		backend.externalReferenceID = id

		if err := backend.StoreHashesEntries(ref.Hashes); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (backend *Backend) StoreHashesEntries(hashes map[int32]string) error {
	// Set up client if this method was called directly.
	if backend.BackendOptions == nil {
		backend.BackendOptions = NewBackendOptions()
		defer backend.client.Close()
	}

	backend.init(backend.BackendOptions)

	entries := []*ent.HashesEntryCreate{}

	for alg, content := range hashes {
		algName := sbom.HashAlgorithm_name[alg]

		entry := backend.client.HashesEntry.Create().
			SetHashAlgorithmType(hashesentry.HashAlgorithmType(algName)).
			SetHashData(content)

		if backend.externalReferenceID != 0 {
			entry.SetExternalReferenceID(backend.externalReferenceID)
		}

		if backend.nodeID != "" {
			entry.SetNodeID(backend.nodeID)
		}

		entries = append(entries, entry)
	}

	if err := backend.client.HashesEntry.CreateBulk(entries...).
		Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.HashesEntry: %w", err)
	}

	return nil
}

func (backend *Backend) StoreIdentifiersEntries(idents map[int32]string) error {
	// Set up client if this method was called directly.
	if backend.BackendOptions == nil {
		backend.BackendOptions = NewBackendOptions()
		defer backend.client.Close()
	}

	backend.init(backend.BackendOptions)

	entries := []*ent.IdentifiersEntryCreate{}

	for typ, value := range idents {
		typeName := sbom.SoftwareIdentifierType_name[typ]

		entry := backend.client.IdentifiersEntry.Create().
			SetSoftwareIdentifierType(identifiersentry.SoftwareIdentifierType(typeName)).
			SetSoftwareIdentifierValue(value)

		if backend.nodeID != "" {
			entry.SetNodeID(backend.nodeID)
		}

		entries = append(entries, entry)
	}

	if err := backend.client.IdentifiersEntry.CreateBulk(entries...).
		Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.IdentifiersEntry: %w", err)
	}

	return nil
}

func (backend *Backend) StoreMetadata(md *sbom.Metadata) error {
	// Set up client if this method was called directly.
	if backend.BackendOptions == nil {
		backend.BackendOptions = NewBackendOptions()
		defer backend.client.Close()
	}

	backend.init(backend.BackendOptions)

	newMetadata := backend.client.Metadata.Create().
		SetID(md.Id).
		SetVersion(md.Version).
		SetName(md.Name).
		SetComment(md.Comment).
		SetDate(md.Date.AsTime())

	err := newMetadata.OnConflict().Ignore().Exec(backend.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.Metadata: %w", err)
	}

	backend.metadataID = md.Id

	if err := backend.StorePersons(md.Authors); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := backend.StoreDocumentTypes(md.DocumentTypes); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := backend.StoreTools(md.Tools); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (backend *Backend) StoreNodeList(nodeList *sbom.NodeList) error {
	// Set up client if this method was called directly.
	if backend.BackendOptions == nil {
		backend.BackendOptions = NewBackendOptions()
		defer backend.client.Close()
	}

	backend.init(backend.BackendOptions)

	newNodeList := backend.client.NodeList.Create().
		SetRootElements(nodeList.RootElements)

	id, err := newNodeList.OnConflict().Ignore().ID(backend.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.NodeList: %w", err)
	}

	backend.nodeListID = id

	if err := backend.StoreNodes(nodeList.Nodes); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (backend *Backend) StoreNodes(nodes []*sbom.Node) error { //nolint:cyclop
	// Set up client if this method was called directly.
	if backend.BackendOptions == nil {
		backend.BackendOptions = NewBackendOptions()
		defer backend.client.Close()
	}

	backend.init(backend.BackendOptions)

	for _, n := range nodes {
		newNode := backend.newNodeCreate(n)

		err := newNode.OnConflict().Ignore().Exec(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.Node: %w", err)
		}

		backend.nodeID = n.Id

		if err := backend.StoreExternalReferences(n.ExternalReferences); err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := backend.StorePersons(n.Originators); err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := backend.StorePersons(n.Suppliers); err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := backend.StorePurposes(n.PrimaryPurpose); err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := backend.StoreHashesEntries(n.Hashes); err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := backend.StoreIdentifiersEntries(n.Identifiers); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (backend *Backend) StorePersons(persons []*sbom.Person) error {
	// Set up client if this method was called directly.
	if backend.BackendOptions == nil {
		backend.BackendOptions = NewBackendOptions()
		defer backend.client.Close()
	}

	backend.init(backend.BackendOptions)

	for _, p := range persons {
		newPerson := backend.client.Person.Create().
			SetName(p.Name).
			SetEmail(p.Email).
			SetIsOrg(p.IsOrg).
			SetPhone(p.Phone).
			SetURL(p.Url)

		if backend.contactOwnerID != 0 {
			newPerson.SetContactOwnerID(backend.contactOwnerID)
		}

		if backend.metadataID != "" {
			newPerson.SetMetadataID(backend.metadataID)
		}

		if backend.nodeID != "" {
			newPerson.SetNodeID(backend.nodeID)
		}

		id, err := newPerson.OnConflict().Ignore().ID(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.ExternalReference: %w", err)
		}

		backend.contactOwnerID = id

		if err := backend.StorePersons(p.Contacts); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (backend *Backend) StorePurposes(purposes []sbom.Purpose) error {
	// Set up client if this method was called directly.
	if backend.BackendOptions == nil {
		backend.BackendOptions = NewBackendOptions()
		defer backend.client.Close()
	}

	backend.init(backend.BackendOptions)

	builders := []*ent.PurposeCreate{}

	for idx := range purposes {
		newPurpose := backend.client.Purpose.Create().
			SetPrimaryPurpose(purpose.PrimaryPurpose(purposes[idx].String()))

		if backend.nodeID != "" {
			newPurpose.SetNodeID(backend.nodeID)
		}

		builders = append(builders, newPurpose)
	}

	err := backend.client.Purpose.CreateBulk(builders...).
		OnConflict().
		Ignore().
		Exec(backend.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.Tool: %w", err)
	}

	return nil
}

func (backend *Backend) StoreTools(tools []*sbom.Tool) error {
	// Set up client if this method was called directly.
	if backend.BackendOptions == nil {
		backend.BackendOptions = NewBackendOptions()
		defer backend.client.Close()
	}

	backend.init(backend.BackendOptions)

	builders := []*ent.ToolCreate{}

	for _, t := range tools {
		newTool := backend.client.Tool.Create().
			SetName(t.Name).
			SetVersion(t.Version).
			SetVendor(t.Vendor)

		if backend.metadataID != "" {
			newTool.SetMetadataID(backend.metadataID)
		}

		builders = append(builders, newTool)
	}

	err := backend.client.Tool.CreateBulk(builders...).
		OnConflict().
		Ignore().
		Exec(backend.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.Tool: %w", err)
	}

	return nil
}

func (backend *Backend) newNodeCreate(n *sbom.Node) *ent.NodeCreate {
	newNode := backend.client.Node.Create().
		SetID(n.Id).
		SetAttribution(n.Attribution).
		SetBuildDate(n.BuildDate.AsTime()).
		SetComment(n.Comment).
		SetCopyright(n.Copyright).
		SetDescription(n.Description).
		SetFileName(n.FileName).
		SetFileTypes(n.FileTypes).
		SetLicenseComments(n.LicenseComments).
		SetLicenseConcluded(n.LicenseConcluded).
		SetLicenses(n.Licenses).
		SetName(n.Name).
		SetReleaseDate(n.ReleaseDate.AsTime()).
		SetSourceInfo(n.SourceInfo).
		SetSummary(n.Summary).
		SetType(node.Type(n.Type.String())).
		SetURLDownload(n.UrlDownload).
		SetURLHome(n.UrlHome).
		SetValidUntilDate(n.ValidUntilDate.AsTime()).
		SetVersion(n.Version)

	if backend.nodeListID != 0 {
		newNode.SetNodeListID(backend.nodeListID)
	}

	return newNode
}
