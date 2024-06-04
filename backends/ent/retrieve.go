// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"fmt"
	"time"

	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/metadata"
)

// Retrieve implements the storage.Retriever interface.
func (backend *Backend) Retrieve(id string, _ *storage.RetrieveOptions) (*sbom.Document, error) {
	if backend.client == nil {
		return nil, fmt.Errorf("%w", errUninitializedClient)
	}

	if backend.Options == nil {
		backend.Options = NewBackendOptions()
	}

	entDoc, err := backend.client.Document.Query().
		Where(document.HasMetadataWith(metadata.IDEQ(id))).
		WithMetadata().
		WithNodeList().
		Only(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	entDoc.Edges.Metadata, err = entDoc.QueryMetadata().
		WithAuthors().
		WithDocumentTypes().
		WithTools().
		Only(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	entDoc.Edges.NodeList, err = entDoc.QueryNodeList().
		WithNodes().
		Only(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
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
		return nil, fmt.Errorf("%w", err)
	}

	return &sbom.Document{
		Metadata: entMetadataToProtobom(entDoc.Edges.Metadata),
		NodeList: entNodeListToProtobom(entDoc.Edges.NodeList),
	}, nil
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

	edgeMap := make(map[string]*sbom.Edge)

	for _, n := range nl.Edges.Nodes {
		pbnl.Nodes = append(pbnl.Nodes, entNodeToProtobom(n))

		for _, e := range n.Edges.EdgeTypes {
			if edgeMap[e.NodeID] == nil {
				edgeMap[e.NodeID] = &sbom.Edge{From: e.NodeID}
			}

			edgeType := sbom.Edge_Type_value[e.Type.String()]
			edgeMap[e.NodeID].To = append(edgeMap[e.NodeID].To, e.ToNodeID)
			edgeMap[e.NodeID].Type = sbom.Edge_Type(edgeType)
		}
	}

	for _, edge := range edgeMap {
		if len(edge.To) > 0 {
			pbnl.Edges = append(pbnl.Edges, edge)
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
