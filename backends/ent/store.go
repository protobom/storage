// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"context"
	"crypto/sha256"
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"
	"google.golang.org/protobuf/proto"

	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/documenttype"
	"github.com/protobom/storage/internal/backends/ent/edgetype"
	"github.com/protobom/storage/internal/backends/ent/externalreference"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/purpose"
)

type (
	contactOwnerIDKey      struct{}
	documentIDKey          struct{}
	externalReferenceIDKey struct{}
	metadataIDKey          struct{}
	nodeIDKey              struct{}
	nodeListIDKey          struct{}

	txFunc func(*ent.Tx) error
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

	backendOpts, ok := opts.BackendOptions.(*BackendOptions)
	if !ok {
		return fmt.Errorf("%w", errInvalidEntOptions)
	}

	// Append annotations from opts parameter with any previously set on the backend.
	annotations := slices.Concat(backend.Options.Annotations, backendOpts.Annotations)
	clear(backend.Options.Annotations)

	id, err := uuidFromHash(doc)
	if err != nil {
		return err
	}

	backend.ctx = context.WithValue(backend.ctx, documentIDKey{}, id)

	// Set each annotation's document ID if not specified.
	for _, a := range annotations {
		if a.DocumentID == uuid.Nil {
			a.DocumentID = id
		}
	}

	return backend.withTx(
		func(tx *ent.Tx) error {
			return tx.Document.Create().
				SetID(id).
				SetProtoMessage(doc).
				OnConflict().
				Ignore().
				Exec(backend.ctx)
		},
		backend.saveAnnotations(annotations...),
		backend.saveMetadata(doc.Metadata),
		backend.saveNodeList(doc.NodeList),
	)
}

func (backend *Backend) addNodeListIDs(builder interface{ AddNodeListIDs(...uuid.UUID) }) {
	if nodeListID, ok := backend.ctx.Value(nodeListIDKey{}).(uuid.UUID); ok {
		builder.AddNodeListIDs(nodeListID)
	}
}

func (backend *Backend) saveAnnotations(annotations ...*ent.Annotation) txFunc {
	return func(tx *ent.Tx) error {
		builders := []*ent.AnnotationCreate{}

		for idx := range annotations {
			builder := tx.Annotation.Create().
				SetDocumentID(annotations[idx].DocumentID).
				SetName(annotations[idx].Name).
				SetValue(annotations[idx].Value).
				SetIsUnique(annotations[idx].IsUnique)

			builders = append(builders, builder)
		}

		err := tx.Annotation.CreateBulk(builders...).
			OnConflict().
			UpdateNewValues().
			Exec(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("creating annotations: %w", err)
		}

		return nil
	}
}

func (backend *Backend) saveDocumentTypes(docTypes []*sbom.DocumentType) txFunc {
	return func(tx *ent.Tx) error {
		for _, dt := range docTypes {
			id, err := uuidFromHash(dt)
			if err != nil {
				return err
			}

			typeName := documenttype.Type(dt.Type.String())

			newDocType := tx.DocumentType.Create().
				SetID(id).
				SetProtoMessage(dt).
				SetNillableType(&typeName).
				SetNillableName(dt.Name).
				SetNillableDescription(dt.Description)

			backend.setDocumentID(newDocType.Mutation())
			backend.setMetadataID(newDocType.Mutation())

			if err := newDocType.OnConflict().Ignore().Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
				return fmt.Errorf("saving document type: %w", err)
			}
		}

		return nil
	}
}

func (backend *Backend) saveEdges(edges []*sbom.Edge) txFunc {
	return func(tx *ent.Tx) error {
		for _, edge := range edges {
			for _, toID := range edge.To {
				newEdgeType := tx.EdgeType.Create().
					SetType(edgetype.Type(edge.Type.String())).
					SetFromID(edge.From).
					SetToID(toID)

				backend.setDocumentID(newEdgeType.Mutation())

				if err := newEdgeType.
					OnConflict().
					Ignore().
					Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
					return fmt.Errorf("saving edge: %w", err)
				}
			}
		}

		return nil
	}
}

func (backend *Backend) saveExternalReferences(refs []*sbom.ExternalReference) txFunc {
	return func(tx *ent.Tx) error {
		for _, ref := range refs {
			id, err := uuidFromHash(ref)
			if err != nil {
				return err
			}

			newRef := tx.ExternalReference.Create().
				SetID(id).
				SetProtoMessage(ref).
				SetURL(ref.Url).
				SetComment(ref.Comment).
				SetAuthority(ref.Authority).
				SetType(externalreference.Type(ref.Type.String()))

			backend.setDocumentID(newRef.Mutation())
			backend.setNodeID(newRef.Mutation())

			if err := newRef.OnConflict().Ignore().Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
				return fmt.Errorf("saving external reference: %w", err)
			}

			backend.ctx = context.WithValue(backend.ctx, externalReferenceIDKey{}, id)
		}

		return nil
	}
}

func (backend *Backend) saveMetadata(md *sbom.Metadata) txFunc {
	return func(tx *ent.Tx) error {
		newMetadata := tx.Metadata.Create().
			SetID(md.Id).
			SetProtoMessage(md).
			SetVersion(md.Version).
			SetName(md.Name).
			SetComment(md.Comment).
			SetDate(md.Date.AsTime())

		backend.setDocumentID(newMetadata.Mutation())

		if err := newMetadata.OnConflict().Ignore().Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving metadata: %w", err)
		}

		backend.ctx = context.WithValue(backend.ctx, metadataIDKey{}, md.Id)

		for _, fn := range []txFunc{
			backend.savePersons(md.Authors),
			backend.saveDocumentTypes(md.DocumentTypes),
			backend.saveTools(md.Tools),
		} {
			if err := fn(tx); err != nil {
				return err
			}
		}

		return nil
	}
}

func (backend *Backend) saveNodeList(nodeList *sbom.NodeList) txFunc {
	return func(tx *ent.Tx) error {
		id, err := uuidFromHash(nodeList)
		if err != nil {
			return err
		}

		newNodeList := tx.NodeList.Create().
			SetID(id).
			SetProtoMessage(nodeList).
			SetRootElements(nodeList.RootElements)

		backend.setDocumentID(newNodeList.Mutation())

		if err := newNodeList.OnConflict().Ignore().Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving node list: %w", err)
		}

		backend.ctx = context.WithValue(backend.ctx, nodeListIDKey{}, id)

		for _, fn := range []txFunc{
			backend.saveNodes(nodeList.Nodes),
			backend.saveEdges(nodeList.Edges),
		} {
			if err := fn(tx); err != nil {
				return err
			}
		}

		return nil
	}
}

func (backend *Backend) saveNodes(nodes []*sbom.Node) txFunc {
	return func(tx *ent.Tx) error {
		for _, n := range nodes {
			newNode := tx.Node.Create().
				SetID(n.Id).
				SetProtoMessage(n).
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

			backend.addNodeListIDs(newNode.Mutation())
			backend.setDocumentID(newNode.Mutation())

			if err := newNode.OnConflict().Ignore().Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
				return fmt.Errorf("saving node: %w", err)
			}

			backend.ctx = context.WithValue(backend.ctx, nodeIDKey{}, n.Id)

			for _, fn := range []txFunc{
				backend.saveExternalReferences(n.ExternalReferences),
				backend.savePersons(n.Originators),
				backend.savePersons(n.Suppliers),
				backend.savePurposes(n.PrimaryPurpose),
			} {
				if err := fn(tx); err != nil {
					return err
				}
			}
		}

		return nil
	}
}

func (backend *Backend) savePersons(persons []*sbom.Person) txFunc {
	return func(tx *ent.Tx) error {
		for _, p := range persons {
			id, err := uuidFromHash(p)
			if err != nil {
				return err
			}

			newPerson := tx.Person.Create().
				SetID(id).
				SetProtoMessage(p).
				SetName(p.Name).
				SetEmail(p.Email).
				SetIsOrg(p.IsOrg).
				SetPhone(p.Phone).
				SetURL(p.Url)

			backend.setContactOwnerID(newPerson.Mutation())
			backend.setDocumentID(newPerson.Mutation())
			backend.setMetadataID(newPerson.Mutation())

			if err := newPerson.OnConflict().Ignore().Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
				return fmt.Errorf("saving person: %w", err)
			}

			backend.ctx = context.WithValue(backend.ctx, contactOwnerIDKey{}, id)

			if err := backend.savePersons(p.Contacts)(tx); err != nil {
				return err
			}
		}

		return nil
	}
}

func (backend *Backend) savePurposes(purposes []sbom.Purpose) txFunc {
	return func(tx *ent.Tx) error {
		builders := []*ent.PurposeCreate{}

		for idx := range purposes {
			newPurpose := tx.Purpose.Create().
				SetPrimaryPurpose(purpose.PrimaryPurpose(purposes[idx].String()))

			backend.setNodeID(newPurpose.Mutation())
			backend.setDocumentID(newPurpose.Mutation())

			builders = append(builders, newPurpose)
		}

		err := tx.Purpose.CreateBulk(builders...).
			OnConflict().
			Ignore().
			Exec(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving Tool: %w", err)
		}

		return nil
	}
}

func (backend *Backend) saveTools(tools []*sbom.Tool) txFunc {
	return func(tx *ent.Tx) error {
		builders := []*ent.ToolCreate{}

		for _, t := range tools {
			id, err := uuidFromHash(t)
			if err != nil {
				return err
			}

			newTool := tx.Tool.Create().
				SetID(id).
				SetProtoMessage(t).
				SetName(t.Name).
				SetVersion(t.Version).
				SetVendor(t.Vendor)

			backend.setDocumentID(newTool.Mutation())
			backend.setMetadataID(newTool.Mutation())

			builders = append(builders, newTool)
		}

		err := tx.Tool.CreateBulk(builders...).
			OnConflict().
			Ignore().
			Exec(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving Tool: %w", err)
		}

		return nil
	}
}

func (backend *Backend) setContactOwnerID(builder interface{ SetContactOwnerID(uuid.UUID) }) {
	if contactOwnerID, ok := backend.ctx.Value(contactOwnerIDKey{}).(uuid.UUID); ok {
		builder.SetContactOwnerID(contactOwnerID)
	}
}

func (backend *Backend) setDocumentID(builder interface{ SetDocumentID(uuid.UUID) }) {
	if documentID, ok := backend.ctx.Value(documentIDKey{}).(uuid.UUID); ok {
		builder.SetDocumentID(documentID)
	}
}

func (backend *Backend) setMetadataID(builder interface{ SetMetadataID(string) }) {
	if metadataID, ok := backend.ctx.Value(metadataIDKey{}).(string); ok {
		builder.SetMetadataID(metadataID)
	}
}

func (backend *Backend) setNodeID(builder interface{ SetNodeID(string) }) {
	if nodeID, ok := backend.ctx.Value(nodeIDKey{}).(string); ok {
		builder.SetNodeID(nodeID)
	}
}

func rollback(tx *ent.Tx, err error) error {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		return fmt.Errorf("rolling back transaction: %w", rollbackErr)
	}

	return err
}

func uuidFromHash(msg proto.Message) (uuid.UUID, error) {
	data, err := proto.MarshalOptions{Deterministic: true}.Marshal(msg)
	if err != nil {
		return uuid.Nil, fmt.Errorf("marshaling proto: %w", err)
	}

	return uuid.NewHash(sha256.New(), uuid.Max, data, int(uuid.Max.Version())), nil
}
