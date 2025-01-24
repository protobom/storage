// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent

import (
	"context"
	"crypto/sha256"
	"errors"
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
	"github.com/protobom/storage/internal/backends/ent/hashesentry"
	"github.com/protobom/storage/internal/backends/ent/identifiersentry"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/purpose"
)

type (
	documentIDKey          struct{}
	nodeNativeIDMappingKey struct{}

	TxFunc func(*ent.Tx) error
)

var errNativeIDMap = errors.New("retrieving node map from context")

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

	id, err := GenerateUUID(doc)
	if err != nil {
		return err
	}

	backend.ctx = context.WithValue(backend.ctx, documentIDKey{}, id)

	// Set each annotation's document ID if not specified.
	for _, a := range annotations {
		if a.DocumentID == nil || a.DocumentID == &uuid.Nil {
			a.DocumentID = &id
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
				SetNillableDocumentID(annotations[idx].DocumentID).
				SetNillableNodeID(annotations[idx].NodeID).
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

func (backend *Backend) saveDocumentTypes(docTypes []*sbom.DocumentType, opts ...func(*ent.DocumentTypeCreate)) TxFunc {
	return func(tx *ent.Tx) error {
		for _, docType := range docTypes {
			typeName := documenttype.Type(docType.GetType().String())

			newDocType := tx.DocumentType.Create().
				SetProtoMessage(docType).
				SetNillableType(&typeName).
				SetNillableName(docType.Name).              //nolint:protogetter
				SetNillableDescription(docType.Description) //nolint:protogetter

			for _, fn := range opts {
				fn(newDocType)
			}

			if err := newDocType.OnConflict().Ignore().Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
				return fmt.Errorf("saving document type: %w", err)
			}
		}

		return nil
	}
}

func (backend *Backend) saveEdges(edges []*sbom.Edge, opts ...func(*ent.EdgeTypeCreate)) TxFunc { //nolint:gocognit
	return func(tx *ent.Tx) error {
		nativeIDMap, ok := backend.ctx.Value(nodeNativeIDMappingKey{}).(map[string]uuid.UUID)
		if !ok {
			return errNativeIDMap
		}

		for _, edge := range edges {
			for _, toID := range edge.GetTo() {
				newEdgeType := tx.EdgeType.Create().
					SetProtoMessage(edge).
					SetType(edgetype.Type(edge.GetType().String())).
					SetFromID(nativeIDMap[edge.GetFrom()]).
					SetToID(nativeIDMap[toID])

				for _, fn := range opts {
					fn(newEdgeType)
				}

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

func (backend *Backend) saveExternalReferences(refs []*sbom.ExternalReference, opts ...func(*ent.ExternalReferenceCreate)) TxFunc { //nolint:gocognit,lll
	return func(tx *ent.Tx) error {
		builders := []*ent.ExternalReferenceCreate{}
		fns := []TxFunc{}

		for _, ref := range refs {
			extRefID, err := GenerateUUID(ref)
			if err != nil {
				return err
			}

			newRef := tx.ExternalReference.Create().
				SetProtoMessage(ref).
				SetURL(ref.GetUrl()).
				SetComment(ref.GetComment()).
				SetAuthority(ref.GetAuthority()).
				SetType(externalreference.Type(ref.GetType().String()))

			for _, fn := range opts {
				fn(newRef)
			}

			builders = append(builders, newRef)

			fns = append(fns, backend.saveHashes(ref.GetHashes(), func(hec *ent.HashesEntryCreate) {
				hec.AddExternalReferenceIDs(extRefID)
				addDocumentIDs(backend.ctx, hec)
			}))
		}

		err := tx.ExternalReference.CreateBulk(builders...).
			OnConflict().
			Ignore().
			Exec(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving external references: %w", err)
		}

		for _, fn := range fns {
			if err := fn(tx); err != nil {
				return err
			}
		}

		return nil
	}
}

func (backend *Backend) saveHashes(hashes map[int32]string, opts ...func(*ent.HashesEntryCreate)) TxFunc {
	return func(tx *ent.Tx) error {
		builders := []*ent.HashesEntryCreate{}

		for key, value := range hashes {
			alg := sbom.HashAlgorithm(key)

			hashesEntry := tx.HashesEntry.Create().
				SetHashAlgorithm(hashesentry.HashAlgorithm(alg.String())).
				SetHashData(value)

			for _, fn := range opts {
				fn(hashesEntry)
			}

			builders = append(builders, hashesEntry)
		}

		if err := tx.HashesEntry.CreateBulk(builders...).
			OnConflict().
			Ignore().
			Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving hashes: %w", err)
		}

		return nil
	}
}

func (backend *Backend) saveIdentifiers(idents map[int32]string, opts ...func(*ent.IdentifiersEntryCreate)) TxFunc {
	return func(tx *ent.Tx) error {
		builders := []*ent.IdentifiersEntryCreate{}

		for key, value := range idents {
			identType := sbom.SoftwareIdentifierType(key)

			identEntry := tx.IdentifiersEntry.Create().
				SetType(identifiersentry.Type(identType.String())).
				SetValue(value)

			for _, fn := range opts {
				fn(identEntry)
			}

			builders = append(builders, identEntry)
		}

		if err := tx.IdentifiersEntry.CreateBulk(builders...).
			OnConflict().
			Ignore().
			Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving identifiers: %w", err)
		}

		return nil
	}
}

func (backend *Backend) saveMetadata(metadata *sbom.Metadata) TxFunc {
	id, err := GenerateUUID(metadata)
	if err != nil {
		return nil
	}

	return func(tx *ent.Tx) error {
		newMetadata := tx.Metadata.Create().
			SetNativeID(metadata.GetId()).
			SetProtoMessage(metadata).
			SetVersion(metadata.GetVersion()).
			SetName(metadata.GetName()).
			SetComment(metadata.GetComment()).
			SetDate(metadata.GetDate().AsTime())

		addDocumentIDs(backend.ctx, newMetadata)

		if err := newMetadata.OnConflict().Ignore().Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving metadata: %w", err)
		}

		for _, fn := range []TxFunc{
			backend.savePersons(metadata.GetAuthors(), func(pc *ent.PersonCreate) {
				pc.AddMetadatumIDs(id)
				addDocumentIDs(backend.ctx, pc)
			}),
			backend.saveDocumentTypes(metadata.GetDocumentTypes(), func(dtc *ent.DocumentTypeCreate) {
				dtc.AddMetadatumIDs(id)
				addDocumentIDs(backend.ctx, dtc)
			}),
			backend.saveSourceData(metadata.GetSourceData(), func(sdc *ent.SourceDataCreate) {
				sdc.AddMetadatumIDs(id)
				addDocumentIDs(backend.ctx, sdc)
			}),
			backend.saveTools(metadata.GetTools(), func(tc *ent.ToolCreate) {
				tc.AddMetadatumIDs(id)
				addDocumentIDs(backend.ctx, tc)
			}),
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
		id, err := GenerateUUID(nodeList)
		if err != nil {
			return err
		}

		newNodeList := tx.NodeList.Create().
			SetProtoMessage(nodeList).
			SetRootElements(nodeList.GetRootElements())

		addDocumentIDs(backend.ctx, newNodeList)

		if err := newNodeList.OnConflict().Ignore().Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving node list: %w", err)
		}

		for _, fn := range []TxFunc{
			backend.saveNodes(nodeList.GetNodes(), func(nc *ent.NodeCreate) {
				nc.AddNodeListIDs(id)
				addDocumentIDs(backend.ctx, nc)
			}),
			backend.saveEdges(nodeList.GetEdges(), func(etc *ent.EdgeTypeCreate) {
				etc.AddNodeListIDs(id)
				addDocumentIDs(backend.ctx, etc)
			}),
		} {
			if err := fn(tx); err != nil {
				return err
			}
		}

		return nil
	}
}

func (backend *Backend) saveNodes(nodes []*sbom.Node, opts ...func(*ent.NodeCreate)) TxFunc { //nolint:funlen,gocognit
	return func(tx *ent.Tx) error {
		builders := []*ent.NodeCreate{}
		fns := []TxFunc{}
		nativeIDMap := make(map[string]uuid.UUID)

		for _, srcNode := range nodes {
			nodeID, err := GenerateUUID(srcNode)
			if err != nil {
				return fmt.Errorf("generating UUID: %w", err)
			}

			nativeIDMap[srcNode.GetId()] = nodeID

			newNode := tx.Node.Create().
				SetProtoMessage(srcNode).
				SetNativeID(srcNode.GetId()).
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

			for _, fn := range opts {
				fn(newNode)
			}

			builders = append(builders, newNode)

			fns = append(fns,
				backend.saveExternalReferences(srcNode.GetExternalReferences(), func(erc *ent.ExternalReferenceCreate) {
					erc.AddNodeIDs(nodeID)
					addDocumentIDs(backend.ctx, erc)
				}),
				backend.saveHashes(srcNode.GetHashes(), func(hec *ent.HashesEntryCreate) {
					hec.AddNodeIDs(nodeID)
					addDocumentIDs(backend.ctx, hec)
				}),
				backend.saveIdentifiers(srcNode.GetIdentifiers(), func(iec *ent.IdentifiersEntryCreate) {
					iec.AddNodeIDs(nodeID)
					addDocumentIDs(backend.ctx, iec)
				}),
				backend.savePersons(srcNode.GetOriginators(), func(pc *ent.PersonCreate) {
					pc.AddOriginatorNodeIDs(nodeID)
					addDocumentIDs(backend.ctx, pc)
				}),
				backend.savePersons(srcNode.GetSuppliers(), func(pc *ent.PersonCreate) {
					pc.AddSupplierNodeIDs(nodeID)
					addDocumentIDs(backend.ctx, pc)
				}),
				backend.saveProperties(srcNode.GetProperties(), func(pc *ent.PropertyCreate) {
					pc.AddNodeIDs(nodeID)
					addDocumentIDs(backend.ctx, pc)
				}),
				backend.savePurposes(srcNode.GetPrimaryPurpose(), func(pc *ent.PurposeCreate) {
					pc.AddNodeIDs(nodeID)
					addDocumentIDs(backend.ctx, pc)
				}),
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

		backend.ctx = context.WithValue(backend.ctx, nodeNativeIDMappingKey{}, nativeIDMap)

		return nil
	}
}

func (backend *Backend) savePersons(persons []*sbom.Person, opts ...func(*ent.PersonCreate)) TxFunc { //nolint:gocognit
	return func(tx *ent.Tx) error {
		builders := []*ent.PersonCreate{}

		for _, person := range persons {
			id, err := GenerateUUID(person)
			if err != nil {
				return err
			}

			newPerson := tx.Person.Create().
				SetProtoMessage(person).
				SetName(person.GetName()).
				SetEmail(person.GetEmail()).
				SetIsOrg(person.GetIsOrg()).
				SetPhone(person.GetPhone()).
				SetURL(person.GetUrl())

			for _, fn := range opts {
				fn(newPerson)
			}

			builders = append(builders, newPerson)

			if err := backend.savePersons(person.GetContacts(), func(pc *ent.PersonCreate) {
				pc.AddContactOwnerIDs(id)
				addDocumentIDs(backend.ctx, pc)
			})(tx); err != nil {
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

func (backend *Backend) saveProperties(properties []*sbom.Property, opts ...func(*ent.PropertyCreate)) TxFunc {
	return func(tx *ent.Tx) error {
		builders := []*ent.PropertyCreate{}

		for _, prop := range properties {
			newProp := tx.Property.Create().
				SetProtoMessage(prop).
				SetName(prop.GetName()).
				SetData(prop.GetData())

			for _, fn := range opts {
				fn(newProp)
			}

			builders = append(builders, newProp)
		}

		err := tx.Property.CreateBulk(builders...).
			OnConflict().
			Ignore().
			Exec(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving property: %w", err)
		}

		return nil
	}
}

func (backend *Backend) savePurposes(purposes []sbom.Purpose, opts ...func(*ent.PurposeCreate)) TxFunc {
	return func(tx *ent.Tx) error {
		builders := []*ent.PurposeCreate{}

		for idx := range purposes {
			newPurpose := tx.Purpose.Create().
				SetPrimaryPurpose(purpose.PrimaryPurpose(purposes[idx].String()))

			for _, fn := range opts {
				fn(newPurpose)
			}

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

func (backend *Backend) saveSourceData(sourceData *sbom.SourceData, opts ...func(*ent.SourceDataCreate)) TxFunc {
	return func(tx *ent.Tx) error {
		newSourceData := tx.SourceData.Create().
			SetProtoMessage(sourceData).
			SetFormat(sourceData.GetFormat()).
			SetSize(sourceData.GetSize()).
			SetURI(sourceData.GetUri())

		for _, fn := range opts {
			fn(newSourceData)
		}

		id, err := newSourceData.OnConflict().Ignore().ID(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving source data: %w", err)
		}

		backend.saveHashes(sourceData.GetHashes(), func(hec *ent.HashesEntryCreate) {
			hec.AddSourceDatumIDs(id)
			addDocumentIDs(backend.ctx, hec)
		})

		return nil
	}
}

func (backend *Backend) saveTools(tools []*sbom.Tool, opts ...func(*ent.ToolCreate)) TxFunc {
	return func(tx *ent.Tx) error {
		builders := []*ent.ToolCreate{}

		for _, tool := range tools {
			newTool := tx.Tool.Create().
				SetProtoMessage(tool).
				SetName(tool.GetName()).
				SetVersion(tool.GetVersion()).
				SetVendor(tool.GetVendor())

			for _, fn := range opts {
				fn(newTool)
			}

			builders = append(builders, newTool)
		}

		err := tx.Tool.CreateBulk(builders...).
			OnConflict().
			Ignore().
			Exec(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("saving tool: %w", err)
		}

		return nil
	}
}

// GenerateUUID returns a deterministic UUID derived from the hash of a protobuf message.
func GenerateUUID(msg proto.Message) (uuid.UUID, error) {
	data, err := proto.MarshalOptions{Deterministic: true}.Marshal(msg)
	if err != nil {
		return uuid.Nil, fmt.Errorf("marshaling proto: %w", err)
	}

	return uuid.NewHash(sha256.New(), uuid.Max, data, int(uuid.Max.Version())), nil
}

func addDocumentIDs[T interface{ AddDocumentIDs(...uuid.UUID) T }](ctx context.Context, builder T) {
	if documentID, ok := ctx.Value(documentIDKey{}).(uuid.UUID); ok {
		builder.AddDocumentIDs(documentID)
	}
}
