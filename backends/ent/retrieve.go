// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"errors"
	"fmt"
	"time"

	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/externalreference"
)

var (
	errMultipleDocuments = errors.New("multiple documents matching ID")
	errMissingDocument   = errors.New("no documents matching IDs")
)

// Retrieve implements the storage.Retriever interface.
func (backend *Backend) Retrieve(id string, _ *storage.RetrieveOptions) (doc *sbom.Document, err error) {
	if backend.client == nil {
		return nil, fmt.Errorf("%w", errUninitializedClient)
	}

	if backend.Options == nil {
		backend.Options = NewBackendOptions()
	}

	switch documents, getDocsErr := backend.GetDocumentsByID(id); {
	case getDocsErr != nil:
		err = fmt.Errorf("querying documents: %w", getDocsErr)
	case len(documents) == 0:
		err = fmt.Errorf("%w %s", errMissingDocument, id)
	case len(documents) > 1:
		err = fmt.Errorf("%w %s", errMultipleDocuments, id)
	default:
		doc = documents[0]
	}

	return
}

func (backend *Backend) GetDocumentsByID(ids ...string) ([]*sbom.Document, error) {
	documents := []*sbom.Document{}

	query := backend.client.Document.Query().
		WithMetadata().
		WithNodeList()

	if len(ids) > 0 {
		query.Where(document.IDIn(ids...))
	}

	entDocs, err := query.All(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying documents table: %w", err)
	}

	for _, entDoc := range entDocs {
		entDoc.Edges.Metadata, err = entDoc.QueryMetadata().
			WithAuthors().
			WithDocumentTypes().
			WithTools().
			Only(backend.ctx)
		if err != nil {
			return nil, fmt.Errorf("eager loading metadata edges: %w", err)
		}

		entDoc.Edges.NodeList, err = entDoc.QueryNodeList().
			WithNodes().
			Only(backend.ctx)
		if err != nil {
			return nil, fmt.Errorf("eager loading node list edges: %w", err)
		}

		// Eager-load the nodes edges of the node list.
		entDoc.Edges.NodeList.Edges.Nodes, err = entDoc.Edges.NodeList.QueryNodes().
			WithEdgeTypes().
			WithExternalReferences().
			WithHashes().
			WithIdentifiers().
			WithNodes().
			WithOriginators().
			WithPrimaryPurpose().
			WithSuppliers().
			All(backend.ctx)
		if err != nil {
			return nil, fmt.Errorf("eager loading node edges: %w", err)
		}

		documents = append(documents, &sbom.Document{
			Metadata: entMetadataToProtobom(entDoc.Edges.Metadata),
			NodeList: entNodeListToProtobom(entDoc.Edges.NodeList),
		})
	}

	return documents, nil
}

func (backend *Backend) GetExternalReferencesByDocumentID(
	id string, types ...string,
) ([]*sbom.ExternalReference, error) {
	query := backend.client.Document.Query().
		Where(document.IDEQ(id)).
		QueryNodeList().
		QueryNodes().
		QueryExternalReferences()

	if len(types) > 0 {
		refTypes := []externalreference.Type{}

		for idx := range types {
			refType := externalreference.Type(types[idx])
			if err := externalreference.TypeValidator(refType); err != nil {
				return nil, fmt.Errorf("%s: %w", types[idx], err)
			}

			refTypes = append(refTypes, refType)
		}

		query.Where(externalreference.TypeIn(refTypes...))
	}

	entExtRefs, err := query.All(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying external references: %w", err)
	}

	return entExtRefsToProtobom(entExtRefs), nil
}

func entExtRefsToProtobom(entRefs ent.ExternalReferences) []*sbom.ExternalReference {
	refs := []*sbom.ExternalReference{}

	for _, ref := range entRefs {
		refType := sbom.ExternalReference_ExternalReferenceType_value[ref.Type.String()]

		refs = append(refs, &sbom.ExternalReference{
			Url:       ref.URL,
			Comment:   ref.Comment,
			Authority: ref.Authority,
			Type:      sbom.ExternalReference_ExternalReferenceType(refType),
			Hashes:    entHashesToProtobom(ref.Edges.Hashes),
		})
	}

	return refs
}

func entHashesToProtobom(entHashes ent.HashesEntries) map[int32]string {
	hashes := make(map[int32]string)

	for _, entry := range entHashes {
		alg := sbom.HashAlgorithm_value[entry.HashAlgorithmType.String()]
		hashes[alg] = entry.HashData
	}

	return hashes
}

func entIdentifiersToProtobom(idents ent.IdentifiersEntries) map[int32]string {
	pbIdents := make(map[int32]string)

	for _, entry := range idents {
		identType := sbom.SoftwareIdentifierType_value[entry.SoftwareIdentifierType.String()]
		pbIdents[identType] = entry.SoftwareIdentifierValue
	}

	return pbIdents
}

func entMetadataToProtobom(entmd *ent.Metadata) *sbom.Metadata {
	authors := entPersonsToProtobom(entmd.Edges.Authors)

	pbmd := &sbom.Metadata{
		Id:            entmd.ID,
		Version:       entmd.Version,
		Name:          entmd.Name,
		Date:          timestamppb.New(entmd.Date),
		Comment:       entmd.Comment,
		Tools:         []*sbom.Tool{},
		Authors:       authors,
		DocumentTypes: []*sbom.DocumentType{},
	}

	for _, dt := range entmd.Edges.DocumentTypes {
		sbomType := sbom.DocumentType_SBOMType(sbom.DocumentType_SBOMType_value[dt.Type.String()])

		pbmd.DocumentTypes = append(pbmd.DocumentTypes, &sbom.DocumentType{
			Name:        dt.Name,
			Description: dt.Description,
			Type:        &sbomType,
		})
	}

	for _, t := range entmd.Edges.Tools {
		pbmd.Tools = append(pbmd.Tools, &sbom.Tool{
			Name:    t.Name,
			Vendor:  t.Vendor,
			Version: t.Version,
		})
	}

	return pbmd
}

func entNodeListToProtobom(nl *ent.NodeList) *sbom.NodeList {
	pbnl := &sbom.NodeList{
		Nodes:        []*sbom.Node{},
		Edges:        []*sbom.Edge{},
		RootElements: nl.RootElements,
	}

	// Mapping of from ID and edge type to slice of to IDs.
	edgeMap := make(map[struct {
		fromID   string
		edgeType string
	}][]string)

	for _, n := range nl.Edges.Nodes {
		pbnl.Nodes = append(pbnl.Nodes, entNodeToProtobom(n))

		for _, e := range n.Edges.EdgeTypes {
			key := struct {
				fromID   string
				edgeType string
			}{e.NodeID, e.Type.String()}

			edgeMap[key] = append(edgeMap[key], e.ToNodeID)
		}
	}

	for typedEdge, toIDs := range edgeMap {
		if len(toIDs) > 0 {
			edgeType := sbom.Edge_Type_value[typedEdge.edgeType]

			pbnl.Edges = append(pbnl.Edges, &sbom.Edge{
				Type: sbom.Edge_Type(edgeType),
				From: typedEdge.fromID,
				To:   toIDs,
			})
		}
	}

	return pbnl
}

func entNodeToProtobom(n *ent.Node) *sbom.Node {
	pbExtRefs := entExtRefsToProtobom(n.Edges.ExternalReferences)

	nodeType := sbom.Node_NodeType_value[n.Type.String()]

	pbNode := &sbom.Node{
		Id:                 n.ID,
		Type:               sbom.Node_NodeType(nodeType),
		Name:               n.Name,
		Version:            n.Version,
		FileName:           n.FileName,
		UrlHome:            n.URLHome,
		UrlDownload:        n.URLDownload,
		Licenses:           n.Licenses,
		LicenseConcluded:   n.LicenseConcluded,
		LicenseComments:    n.LicenseComments,
		Copyright:          n.Copyright,
		SourceInfo:         n.SourceInfo,
		Comment:            n.Comment,
		Summary:            n.Summary,
		Description:        n.Description,
		Attribution:        n.Attribution,
		FileTypes:          n.FileTypes,
		Suppliers:          entPersonsToProtobom(n.Edges.Suppliers),
		Originators:        entPersonsToProtobom(n.Edges.Originators),
		ExternalReferences: pbExtRefs,
		Identifiers:        entIdentifiersToProtobom(n.Edges.Identifiers),
		Hashes:             entHashesToProtobom(n.Edges.Hashes),
		PrimaryPurpose:     entPurposesToProtobom(n.Edges.PrimaryPurpose),
	}

	epoch := time.Unix(0, 0).UTC()

	if n.ReleaseDate != epoch {
		pbNode.ReleaseDate = timestamppb.New(n.ReleaseDate)
	}

	if n.BuildDate != epoch {
		pbNode.BuildDate = timestamppb.New(n.BuildDate)
	}

	if n.ValidUntilDate != epoch {
		pbNode.ValidUntilDate = timestamppb.New(n.ValidUntilDate)
	}

	return pbNode
}

func entPersonsToProtobom(persons ent.Persons) []*sbom.Person {
	pbPersons := []*sbom.Person{}

	for _, p := range persons {
		pbPersons = append(pbPersons, &sbom.Person{
			Name:     p.Name,
			IsOrg:    p.IsOrg,
			Email:    p.Email,
			Url:      p.URL,
			Phone:    p.Phone,
			Contacts: entPersonsToProtobom(p.Edges.Contacts),
		})
	}

	return pbPersons
}

func entPurposesToProtobom(ps ent.Purposes) []sbom.Purpose {
	pbPurposes := []sbom.Purpose{}

	for _, p := range ps {
		purposeValue := sbom.Purpose_value[p.PrimaryPurpose.String()]
		pbPurposes = append(pbPurposes, sbom.Purpose(purposeValue))
	}

	return pbPurposes
}
