// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"errors"
	"fmt"

	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"

	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/documenttype"
	"github.com/protobom/storage/internal/backends/ent/edgetype"
	"github.com/protobom/storage/internal/backends/ent/externalreference"
	"github.com/protobom/storage/internal/backends/ent/hashesentry"
	"github.com/protobom/storage/internal/backends/ent/identifiersentry"
	"github.com/protobom/storage/internal/backends/ent/metadata"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/nodelist"
	"github.com/protobom/storage/internal/backends/ent/purpose"
)

var errInvalidEntOptions = errors.New("invalid ent backend options")

// Store implements the storage.Storer interface.
func (backend *Backend) Store(doc *sbom.Document, opts *storage.StoreOptions) error {
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
	}

	if opts == nil {
		opts = &storage.StoreOptions{
			BackendOptions: backend.Options,
		}
	}

	if _, ok := opts.BackendOptions.(*BackendOptions); !ok {
		return fmt.Errorf("%w", errInvalidEntOptions)
	}

	backend.init(backend.Options)

	defer backend.Options.client.Close()

	if err := backend.StoreMetadata(doc.Metadata); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := backend.StoreNodeList(doc.NodeList); err != nil {
		return fmt.Errorf("%w", err)
	}

	entNodeList, err := backend.Options.client.NodeList.Query().
		Where(
			nodelist.Or(
				nodelist.HasDocumentWith(document.HasMetadataWith(metadata.IDEQ(doc.Metadata.Id))),
				nodelist.Not(nodelist.HasDocument()),
			)).
		Only(backend.Options.ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = backend.Options.client.Document.Create().
		SetMetadataID(doc.Metadata.Id).
		SetNodeListID(entNodeList.ID).
		OnConflict().
		Ignore().
		Exec(backend.Options.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.Document: %w", err)
	}

	return nil
}

func (backend *Backend) StoreDocumentTypes(docTypes []*sbom.DocumentType) error {
	// Set up client if this method was called directly.
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
		defer backend.Options.client.Close()
	}

	backend.init(backend.Options)

	for _, dt := range docTypes {
		typeName := documenttype.Type(dt.String())

		newDocType := backend.Options.client.DocumentType.Create().
			SetNillableType(&typeName).
			SetNillableName(dt.Name).
			SetNillableDescription(dt.Description)

		if backend.Options.metadataID != "" {
			newDocType.SetMetadataID(backend.Options.metadataID)
		}

		err := newDocType.OnConflict().Ignore().Exec(backend.Options.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.DocumentType: %w", err)
		}
	}

	return nil
}

func (backend *Backend) StoreEdges(edges []*sbom.Edge) error {
	// Set up client if this method was called directly.
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
		defer backend.Options.client.Close()
	}

	backend.init(backend.Options)

	for _, edge := range edges {
		for _, toID := range edge.To {
			newEdgeType := backend.Options.client.EdgeType.Create().
				SetType(edgetype.Type(edge.Type.String())).
				SetFromID(edge.From).
				SetToID(toID)

			err := newEdgeType.OnConflict().Ignore().Exec(backend.Options.ctx)
			if err != nil && !ent.IsConstraintError(err) {
				return fmt.Errorf("ent.Node: %w", err)
			}
		}
	}

	return nil
}

func (backend *Backend) StoreExternalReferences(refs []*sbom.ExternalReference) error {
	// Set up client if this method was called directly.
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
		defer backend.Options.client.Close()
	}

	backend.init(backend.Options)

	for _, ref := range refs {
		newRef := backend.Options.client.ExternalReference.Create().
			SetURL(ref.Url).
			SetComment(ref.Comment).
			SetAuthority(ref.Authority).
			SetType(externalreference.Type(ref.Type.String()))

		if backend.Options.nodeID != "" {
			newRef.SetNodeID(backend.Options.nodeID)
		}

		id, err := newRef.OnConflict().Ignore().ID(backend.Options.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.ExternalReference: %w", err)
		}

		backend.Options.externalReferenceID = id

		if err := backend.StoreHashesEntries(ref.Hashes); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (backend *Backend) StoreHashesEntries(hashes map[int32]string) error {
	// Set up client if this method was called directly.
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
		defer backend.Options.client.Close()
	}

	backend.init(backend.Options)

	entries := []*ent.HashesEntryCreate{}

	for alg, content := range hashes {
		algName := sbom.HashAlgorithm_name[alg]

		entry := backend.Options.client.HashesEntry.Create().
			SetHashAlgorithmType(hashesentry.HashAlgorithmType(algName)).
			SetHashData(content)

		if backend.Options.externalReferenceID != 0 {
			entry.SetExternalReferenceID(backend.Options.externalReferenceID)
		}

		if backend.Options.nodeID != "" {
			entry.SetNodeID(backend.Options.nodeID)
		}

		entries = append(entries, entry)
	}

	if err := backend.Options.client.HashesEntry.CreateBulk(entries...).
		Exec(backend.Options.ctx); err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.HashesEntry: %w", err)
	}

	return nil
}

func (backend *Backend) StoreIdentifiersEntries(idents map[int32]string) error {
	// Set up client if this method was called directly.
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
		defer backend.Options.client.Close()
	}

	backend.init(backend.Options)

	entries := []*ent.IdentifiersEntryCreate{}

	for typ, value := range idents {
		typeName := sbom.SoftwareIdentifierType_name[typ]

		entry := backend.Options.client.IdentifiersEntry.Create().
			SetSoftwareIdentifierType(identifiersentry.SoftwareIdentifierType(typeName)).
			SetSoftwareIdentifierValue(value)

		if backend.Options.nodeID != "" {
			entry.SetNodeID(backend.Options.nodeID)
		}

		entries = append(entries, entry)
	}

	if err := backend.Options.client.IdentifiersEntry.CreateBulk(entries...).
		Exec(backend.Options.ctx); err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.IdentifiersEntry: %w", err)
	}

	return nil
}

func (backend *Backend) StoreMetadata(md *sbom.Metadata) error {
	// Set up client if this method was called directly.
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
		defer backend.Options.client.Close()
	}

	backend.init(backend.Options)

	newMetadata := backend.Options.client.Metadata.Create().
		SetID(md.Id).
		SetVersion(md.Version).
		SetName(md.Name).
		SetComment(md.Comment).
		SetDate(md.Date.AsTime())

	err := newMetadata.OnConflict().Ignore().Exec(backend.Options.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.Metadata: %w", err)
	}

	backend.Options.metadataID = md.Id

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
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
		defer backend.Options.client.Close()
	}

	backend.init(backend.Options)

	newNodeList := backend.Options.client.NodeList.Create().
		SetRootElements(nodeList.RootElements)

	id, err := newNodeList.OnConflict().Ignore().ID(backend.Options.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.NodeList: %w", err)
	}

	backend.Options.nodeListID = id

	if err := backend.StoreNodes(nodeList.Nodes); err != nil {
		return fmt.Errorf("%w", err)
	}

	// Update nodes of this node list with their typed edges.
	if err := backend.StoreEdges(nodeList.Edges); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (backend *Backend) StoreNodes(nodes []*sbom.Node) error { //nolint:cyclop
	// Set up client if this method was called directly.
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
		defer backend.Options.client.Close()
	}

	backend.init(backend.Options)

	for _, n := range nodes {
		newNode := backend.newNodeCreate(n)

		err := newNode.OnConflict().Ignore().Exec(backend.Options.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.Node: %w", err)
		}

		backend.Options.nodeID = n.Id

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
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
		defer backend.Options.client.Close()
	}

	backend.init(backend.Options)

	for _, p := range persons {
		newPerson := backend.Options.client.Person.Create().
			SetName(p.Name).
			SetEmail(p.Email).
			SetIsOrg(p.IsOrg).
			SetPhone(p.Phone).
			SetURL(p.Url)

		if backend.Options.contactOwnerID != 0 {
			newPerson.SetContactOwnerID(backend.Options.contactOwnerID)
		}

		if backend.Options.metadataID != "" {
			newPerson.SetMetadataID(backend.Options.metadataID)
		}

		if backend.Options.nodeID != "" {
			newPerson.SetNodeID(backend.Options.nodeID)
		}

		id, err := newPerson.OnConflict().Ignore().ID(backend.Options.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.ExternalReference: %w", err)
		}

		backend.Options.contactOwnerID = id

		if err := backend.StorePersons(p.Contacts); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (backend *Backend) StorePurposes(purposes []sbom.Purpose) error {
	// Set up client if this method was called directly.
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
		defer backend.Options.client.Close()
	}

	backend.init(backend.Options)

	builders := []*ent.PurposeCreate{}

	for idx := range purposes {
		newPurpose := backend.Options.client.Purpose.Create().
			SetPrimaryPurpose(purpose.PrimaryPurpose(purposes[idx].String()))

		if backend.Options.nodeID != "" {
			newPurpose.SetNodeID(backend.Options.nodeID)
		}

		builders = append(builders, newPurpose)
	}

	err := backend.Options.client.Purpose.CreateBulk(builders...).
		OnConflict().
		Ignore().
		Exec(backend.Options.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.Tool: %w", err)
	}

	return nil
}

func (backend *Backend) StoreTools(tools []*sbom.Tool) error {
	// Set up client if this method was called directly.
	if backend.Options == nil {
		backend.Options = NewBackendOptions()
		defer backend.Options.client.Close()
	}

	backend.init(backend.Options)

	builders := []*ent.ToolCreate{}

	for _, t := range tools {
		newTool := backend.Options.client.Tool.Create().
			SetName(t.Name).
			SetVersion(t.Version).
			SetVendor(t.Vendor)

		if backend.Options.metadataID != "" {
			newTool.SetMetadataID(backend.Options.metadataID)
		}

		builders = append(builders, newTool)
	}

	err := backend.Options.client.Tool.CreateBulk(builders...).
		OnConflict().
		Ignore().
		Exec(backend.Options.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.Tool: %w", err)
	}

	return nil
}

func (backend *Backend) newNodeCreate(n *sbom.Node) *ent.NodeCreate {
	newNode := backend.Options.client.Node.Create().
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

	if backend.Options.nodeListID != 0 {
		newNode.SetNodeListID(backend.Options.nodeListID)
	}

	return newNode
}
