// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"context"
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

type (
	contactOwnerIDKey      struct{}
	externalReferenceIDKey struct{}
	metadataIDKey          struct{}
	nodeIDKey              struct{}
	nodeListIDKey          struct{}
)

// Store implements the storage.Storer interface.
func (backend *Backend) Store(doc *sbom.Document, opts *storage.StoreOptions) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	if opts == nil {
		opts = &storage.StoreOptions{
			BackendOptions: NewBackendOptions(),
		}
	}

	if _, ok := opts.BackendOptions.(*BackendOptions); !ok {
		return fmt.Errorf("%w", errInvalidEntOptions)
	}

	if err := backend.StoreMetadata(doc.Metadata); err != nil {
		return err
	}

	if err := backend.StoreNodeList(doc.NodeList); err != nil {
		return err
	}

	nodeListID, ok := backend.ctx.Value(nodeListIDKey{}).(int)
	if !ok {
		var err error

		nodeListID, err = backend.client.NodeList.Query().
			Where(
				nodelist.Or(
					nodelist.HasDocumentWith(document.HasMetadataWith(metadata.IDEQ(doc.Metadata.Id))),
					nodelist.Not(nodelist.HasDocument()),
				)).
			OnlyID(backend.ctx)
		if err != nil {
			return fmt.Errorf("querying node lists: %w", err)
		}
	}

	err := backend.client.Document.Create().
		SetMetadataID(doc.Metadata.Id).
		SetNodeListID(nodeListID).
		OnConflict().
		Ignore().
		Exec(backend.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.Document: %w", err)
	}

	return nil
}

func (backend *Backend) StoreDocumentTypes(docTypes []*sbom.DocumentType) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	for _, dt := range docTypes {
		typeName := documenttype.Type(dt.String())

		newDocType := backend.client.DocumentType.Create().
			SetNillableType(&typeName).
			SetNillableName(dt.Name).
			SetNillableDescription(dt.Description)

		if metadataID, ok := backend.ctx.Value(metadataIDKey{}).(string); ok {
			newDocType.SetMetadataID(metadataID)
		}

		err := newDocType.OnConflict().Ignore().Exec(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.DocumentType: %w", err)
		}
	}

	return nil
}

func (backend *Backend) StoreEdges(edges []*sbom.Edge) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	for _, edge := range edges {
		for _, toID := range edge.To {
			newEdgeType := backend.client.EdgeType.Create().
				SetType(edgetype.Type(edge.Type.String())).
				SetFromID(edge.From).
				SetToID(toID)

			err := newEdgeType.OnConflict().Ignore().Exec(backend.ctx)
			if err != nil && !ent.IsConstraintError(err) {
				return fmt.Errorf("ent.Node: %w", err)
			}
		}
	}

	return nil
}

func (backend *Backend) StoreExternalReferences(refs []*sbom.ExternalReference) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	for _, ref := range refs {
		newRef := backend.client.ExternalReference.Create().
			SetURL(ref.Url).
			SetComment(ref.Comment).
			SetAuthority(ref.Authority).
			SetType(externalreference.Type(ref.Type.String()))

		if nodeID, ok := backend.ctx.Value(nodeIDKey{}).(string); ok {
			newRef.SetNodeID(nodeID)
		}

		id, err := newRef.OnConflict().Ignore().ID(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.ExternalReference: %w", err)
		}

		backend.ctx = context.WithValue(backend.ctx, externalReferenceIDKey{}, id)

		if err := backend.StoreHashesEntries(ref.Hashes); err != nil {
			return err
		}
	}

	return nil
}

func (backend *Backend) StoreHashesEntries(hashes map[int32]string) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	entries := []*ent.HashesEntryCreate{}

	for alg, content := range hashes {
		algName := sbom.HashAlgorithm_name[alg]

		entry := backend.client.HashesEntry.Create().
			SetHashAlgorithmType(hashesentry.HashAlgorithmType(algName)).
			SetHashData(content)

		if externalReferenceID, ok := backend.ctx.Value(externalReferenceIDKey{}).(int); ok {
			entry.SetExternalReferenceID(externalReferenceID)
		}

		if nodeID, ok := backend.ctx.Value(nodeIDKey{}).(string); ok {
			entry.SetNodeID(nodeID)
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
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	entries := []*ent.IdentifiersEntryCreate{}

	for typ, value := range idents {
		typeName := sbom.SoftwareIdentifierType_name[typ]

		entry := backend.client.IdentifiersEntry.Create().
			SetSoftwareIdentifierType(identifiersentry.SoftwareIdentifierType(typeName)).
			SetSoftwareIdentifierValue(value)

		if nodeID, ok := backend.ctx.Value(nodeIDKey{}).(string); ok {
			entry.SetNodeID(nodeID)
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
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

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

	backend.ctx = context.WithValue(backend.ctx, metadataIDKey{}, md.Id)

	if err := backend.StorePersons(md.Authors); err != nil {
		return err
	}

	if err := backend.StoreDocumentTypes(md.DocumentTypes); err != nil {
		return err
	}

	if err := backend.StoreTools(md.Tools); err != nil {
		return err
	}

	return nil
}

func (backend *Backend) StoreNodeList(nodeList *sbom.NodeList) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	newNodeList := backend.client.NodeList.Create().
		SetRootElements(nodeList.RootElements)

	id, err := newNodeList.OnConflict().Ignore().ID(backend.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.NodeList: %w", err)
	}

	backend.ctx = context.WithValue(backend.ctx, nodeListIDKey{}, id)

	if err := backend.StoreNodes(nodeList.Nodes); err != nil {
		return err
	}

	// Update nodes of this node list with their typed edges.
	if err := backend.StoreEdges(nodeList.Edges); err != nil {
		return err
	}

	return nil
}

func (backend *Backend) StoreNodes(nodes []*sbom.Node) error { //nolint:cyclop
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	for _, n := range nodes {
		newNode := backend.newNodeCreate(n)

		err := newNode.OnConflict().Ignore().Exec(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.Node: %w", err)
		}

		backend.ctx = context.WithValue(backend.ctx, nodeIDKey{}, n.Id)

		if err := backend.StoreExternalReferences(n.ExternalReferences); err != nil {
			return err
		}

		if err := backend.StorePersons(n.Originators); err != nil {
			return err
		}

		if err := backend.StorePersons(n.Suppliers); err != nil {
			return err
		}

		if err := backend.StorePurposes(n.PrimaryPurpose); err != nil {
			return err
		}

		if err := backend.StoreHashesEntries(n.Hashes); err != nil {
			return err
		}

		if err := backend.StoreIdentifiersEntries(n.Identifiers); err != nil {
			return err
		}
	}

	return nil
}

func (backend *Backend) StorePersons(persons []*sbom.Person) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	for _, p := range persons {
		newPerson := backend.client.Person.Create().
			SetName(p.Name).
			SetEmail(p.Email).
			SetIsOrg(p.IsOrg).
			SetPhone(p.Phone).
			SetURL(p.Url)

		if contactOwnerID, ok := backend.ctx.Value(contactOwnerIDKey{}).(int); ok {
			newPerson.SetContactOwnerID(contactOwnerID)
		}

		if metadataID, ok := backend.ctx.Value(metadataIDKey{}).(string); ok {
			newPerson.SetMetadataID(metadataID)
		}

		id, err := newPerson.OnConflict().Ignore().ID(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.ExternalReference: %w", err)
		}

		backend.ctx = context.WithValue(backend.ctx, contactOwnerIDKey{}, id)

		if err := backend.StorePersons(p.Contacts); err != nil {
			return err
		}
	}

	return nil
}

func (backend *Backend) StorePurposes(purposes []sbom.Purpose) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	builders := []*ent.PurposeCreate{}

	for idx := range purposes {
		newPurpose := backend.client.Purpose.Create().
			SetPrimaryPurpose(purpose.PrimaryPurpose(purposes[idx].String()))

		if nodeID, ok := backend.ctx.Value(nodeIDKey{}).(string); ok {
			newPurpose.SetNodeID(nodeID)
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
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	builders := []*ent.ToolCreate{}

	for _, t := range tools {
		newTool := backend.client.Tool.Create().
			SetName(t.Name).
			SetVersion(t.Version).
			SetVendor(t.Vendor)

		if metadataID, ok := backend.ctx.Value(metadataIDKey{}).(string); ok {
			newTool.SetMetadataID(metadataID)
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

	if nodeListID, ok := backend.ctx.Value(nodeListIDKey{}).(int); ok {
		newNode.SetNodeListID(nodeListID)
	}

	return newNode
}
