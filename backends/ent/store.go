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

	TxFunc func(*ent.Tx) error
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
				OnConflict().
				Ignore().
				Exec(backend.ctx)
		},
		backend.saveAnnotations(annotations...),
		backend.saveMetadata(doc.GetMetadata()),
		backend.saveNodeList(doc.GetNodeList()),
	)
}

func (backend *Backend) saveAnnotations(annotations ...*ent.Annotation) TxFunc {
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

func (backend *Backend) saveDocumentTypes(docTypes []*sbom.DocumentType) TxFunc {
	return func(tx *ent.Tx) error {
		for _, docType := range docTypes {
			id, err := uuidFromHash(docType)
			if err != nil {
				return err
			}

			typeName := documenttype.Type(docType.GetType().String())

			newDocType := tx.DocumentType.Create().
				SetID(id).
				SetProtoMessage(docType).
				SetNillableType(&typeName).
				SetNillableName(docType.Name).              //nolint:protogetter
				SetNillableDescription(docType.Description) //nolint:protogetter

			setDocumentID(backend.ctx, newDocType)
			setMetadataID(backend.ctx, newDocType)

			if err := newDocType.OnConflict().Ignore().Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
				return fmt.Errorf("saving document type: %w", err)
			}
		}

		return nil
	}
}

func (backend *Backend) saveEdges(edges []*sbom.Edge) TxFunc {
	return func(tx *ent.Tx) error {
		for _, edge := range edges {
			for _, toID := range edge.GetTo() {
				newEdgeType := tx.EdgeType.Create().
					SetType(edgetype.Type(edge.GetType().String())).
					SetFromID(edge.GetFrom()).
					SetToID(toID)

				setDocumentID(backend.ctx, newEdgeType)

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

func (backend *Backend) saveExternalReferences(refs []*sbom.ExternalReference, nodeID string) TxFunc {
	return func(tx *ent.Tx) error {
		for _, ref := range refs {
			id, err := uuidFromHash(ref)
			if err != nil {
				return err
			}

			newRef := tx.ExternalReference.Create().
				SetID(id).
				SetNodeID(nodeID).
				SetProtoMessage(ref).
				SetURL(ref.GetUrl()).
				SetComment(ref.GetComment()).
				SetAuthority(ref.GetAuthority()).
				SetType(externalreference.Type(ref.GetType().String()))

			setDocumentID(backend.ctx, newRef)

			if err := newRef.OnConflict().Ignore().Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
				return fmt.Errorf("saving external reference: %w", err)
			}

			backend.ctx = context.WithValue(backend.ctx, externalReferenceIDKey{}, id)
		}

		return nil
	}
}

func (backend *Backend) saveMetadata(metadata *sbom.Metadata) TxFunc {
	return func(tx *ent.Tx) error {
		newMetadata := tx.Metadata.Create().
			SetID(metadata.GetId()).
			SetProtoMessage(metadata).
			SetVersion(metadata.GetVersion()).
			SetName(metadata.GetName()).
			SetComment(metadata.GetComment()).
			SetDate(metadata.GetDate().AsTime())

		setDocumentID(backend.ctx, newMetadata)

		if err := newMetadata.OnConflict().Ignore().Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving metadata: %w", err)
		}

		backend.ctx = context.WithValue(backend.ctx, metadataIDKey{}, metadata.GetId())

		for _, fn := range []TxFunc{
			backend.savePersons(metadata.GetAuthors()),
			backend.saveDocumentTypes(metadata.GetDocumentTypes()),
			backend.saveTools(metadata.GetTools()),
		} {
			if err := fn(tx); err != nil {
				return err
			}
		}

		return nil
	}
}

func (backend *Backend) saveNodeList(nodeList *sbom.NodeList) TxFunc {
	return func(tx *ent.Tx) error {
		id, err := uuidFromHash(nodeList)
		if err != nil {
			return err
		}

		newNodeList := tx.NodeList.Create().
			SetID(id).
			SetProtoMessage(nodeList).
			SetRootElements(nodeList.GetRootElements())

		setDocumentID(backend.ctx, newNodeList)

		if err := newNodeList.OnConflict().Ignore().Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving node list: %w", err)
		}

		backend.ctx = context.WithValue(backend.ctx, nodeListIDKey{}, id)

		for _, fn := range []TxFunc{
			backend.saveNodes(nodeList.GetNodes()),
			backend.saveEdges(nodeList.GetEdges()),
		} {
			if err := fn(tx); err != nil {
				return err
			}
		}

		return nil
	}
}

func (backend *Backend) saveNodes(nodes []*sbom.Node) TxFunc {
	return func(tx *ent.Tx) error {
		builders := []*ent.NodeCreate{}
		fns := []TxFunc{}

		for _, srcNode := range nodes {
			nodeID := srcNode.GetId()
			newNode := tx.Node.Create().
				SetID(nodeID).
				SetProtoMessage(srcNode).
				SetAttribution(srcNode.GetAttribution()).
				SetBuildDate(srcNode.GetBuildDate().AsTime()).
				SetComment(srcNode.GetComment()).
				SetCopyright(srcNode.GetCopyright()).
				SetDescription(srcNode.GetDescription()).
				SetFileName(srcNode.GetFileName()).
				SetFileTypes(srcNode.GetFileTypes()).
				SetLicenseComments(srcNode.GetLicenseComments()).
				SetLicenseConcluded(srcNode.GetLicenseConcluded()).
				SetLicenses(srcNode.GetLicenses()).
				SetName(srcNode.GetName()).
				SetReleaseDate(srcNode.GetReleaseDate().AsTime()).
				SetSourceInfo(srcNode.GetSourceInfo()).
				SetSummary(srcNode.GetSummary()).
				SetType(node.Type(srcNode.GetType().String())).
				SetURLDownload(srcNode.GetUrlDownload()).
				SetURLHome(srcNode.GetUrlHome()).
				SetValidUntilDate(srcNode.GetValidUntilDate().AsTime()).
				SetVersion(srcNode.GetVersion())

			addNodeListIDs(backend.ctx, newNode)
			setDocumentID(backend.ctx, newNode)

			backend.ctx = context.WithValue(backend.ctx, nodeIDKey{}, nodeID)

			builders = append(builders, newNode)

			fns = append(fns,
				backend.saveExternalReferences(srcNode.GetExternalReferences(), nodeID),
				backend.savePersons(srcNode.GetOriginators()),
				backend.savePersons(srcNode.GetSuppliers()),
				backend.savePurposes(srcNode.GetPrimaryPurpose(), nodeID),
			)
		}

		err := tx.Node.CreateBulk(builders...).
			OnConflict().
			Ignore().
			Exec(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving nodes: %w", err)
		}

		for _, fn := range fns {
			if err := fn(tx); err != nil {
				return err
			}
		}

		return nil
	}
}

func (backend *Backend) savePersons(persons []*sbom.Person) TxFunc { //nolint:gocognit
	return func(tx *ent.Tx) error {
		builders := []*ent.PersonCreate{}

		for _, person := range persons {
			id, err := uuidFromHash(person)
			if err != nil {
				return err
			}

			newPerson := tx.Person.Create().
				SetID(id).
				SetProtoMessage(person).
				SetName(person.GetName()).
				SetEmail(person.GetEmail()).
				SetIsOrg(person.GetIsOrg()).
				SetPhone(person.GetPhone()).
				SetURL(person.GetUrl())

			setContactOwnerID(backend.ctx, newPerson)
			setDocumentID(backend.ctx, newPerson)
			setMetadataID(backend.ctx, newPerson)
			setNodeID(backend.ctx, newPerson)

			builders = append(builders, newPerson)
			backend.ctx = context.WithValue(backend.ctx, contactOwnerIDKey{}, id)

			if err := backend.savePersons(person.GetContacts())(tx); err != nil {
				return err
			}
		}

		if err := tx.Person.CreateBulk(builders...).
			OnConflict().
			Ignore().
			Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving persons: %w", err)
		}

		return nil
	}
}

func (backend *Backend) savePurposes(purposes []sbom.Purpose, nodeID string) TxFunc {
	return func(tx *ent.Tx) error {
		builders := []*ent.PurposeCreate{}

		for idx := range purposes {
			newPurpose := tx.Purpose.Create().
				SetNodeID(nodeID).
				SetPrimaryPurpose(purpose.PrimaryPurpose(purposes[idx].String()))

			setDocumentID(backend.ctx, newPurpose)

			builders = append(builders, newPurpose)
		}

		err := tx.Purpose.CreateBulk(builders...).
			OnConflict().
			Ignore().
			Exec(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving purpose: %w", err)
		}

		return nil
	}
}

func (backend *Backend) saveTools(tools []*sbom.Tool) TxFunc {
	return func(tx *ent.Tx) error {
		builders := []*ent.ToolCreate{}

		for _, tool := range tools {
			id, err := uuidFromHash(tool)
			if err != nil {
				return err
			}

			newTool := tx.Tool.Create().
				SetID(id).
				SetProtoMessage(tool).
				SetName(tool.GetName()).
				SetVersion(tool.GetVersion()).
				SetVendor(tool.GetVendor())

			setDocumentID(backend.ctx, newTool)
			setMetadataID(backend.ctx, newTool)

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

func addNodeListIDs[T interface{ AddNodeListIDs(...uuid.UUID) T }](ctx context.Context, builder T) {
	if nodeListID, ok := ctx.Value(nodeListIDKey{}).(uuid.UUID); ok {
		builder.AddNodeListIDs(nodeListID)
	}
}

func setContactOwnerID[T interface{ SetContactOwnerID(uuid.UUID) T }](ctx context.Context, builder T) {
	if contactOwnerID, ok := ctx.Value(contactOwnerIDKey{}).(uuid.UUID); ok {
		builder.SetContactOwnerID(contactOwnerID)
	}
}

func setDocumentID[T interface{ SetDocumentID(uuid.UUID) T }](ctx context.Context, builder T) {
	if documentID, ok := ctx.Value(documentIDKey{}).(uuid.UUID); ok {
		builder.SetDocumentID(documentID)
	}
}

func setMetadataID[T interface{ SetMetadataID(string) T }](ctx context.Context, builder T) {
	if metadataID, ok := ctx.Value(metadataIDKey{}).(string); ok {
		builder.SetMetadataID(metadataID)
	}
}

func setNodeID[T interface{ SetNodeID(string) T }](ctx context.Context, builder T) {
	if nodeID, ok := ctx.Value(nodeIDKey{}).(string); ok {
		builder.SetNodeID(nodeID)
	}
}

func uuidFromHash(msg proto.Message) (uuid.UUID, error) {
	data, err := proto.MarshalOptions{Deterministic: true}.Marshal(msg)
	if err != nil {
		return uuid.Nil, fmt.Errorf("marshaling proto: %w", err)
	}

	return uuid.NewHash(sha256.New(), uuid.Max, data, int(uuid.Max.Version())), nil
}
