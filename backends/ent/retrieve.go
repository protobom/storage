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
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/metadata"
)

// Retrieve implements the storage.Retriever interface.
func (backend *Backend) Retrieve(id string, _opts *storage.RetrieveOptions) (*sbom.Document, error) {
	if backend.BackendOptions == nil {
		backend.BackendOptions = NewBackendOptions()
	}

	backend.init(backend.BackendOptions)

	defer backend.client.Close()

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

	doc := sbom.NewDocument()
	doc.Metadata = entMetadataToProtobom(entDoc.Edges.Metadata)
	doc.NodeList = entNodeListToProtobom(entDoc.Edges.NodeList)

	return doc, nil
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
	edges := []*sbom.Edge{}
	nodes := []*sbom.Node{}

	for _, n := range nl.Edges.Nodes {
		edgeType := sbom.Edge_Type_value[n.Type.String()]
		toIDs := []string{}

		for _, e := range n.Edges.Nodes {
			toIDs = append(toIDs, e.ID)
		}

		edges = append(edges, &sbom.Edge{
			Type: sbom.Edge_Type(edgeType),
			From: n.ID,
			To:   toIDs,
		})

		nodes = append(nodes, entNodeToProtobom(n))
	}

	return &sbom.NodeList{
		Nodes:        nodes,
		Edges:        edges,
		RootElements: nl.RootElements,
	}
}

func entNodeToProtobom(n *ent.Node) *sbom.Node {
	pbExtRefs := entExtRefsToProtobom(n.Edges.ExternalReferences)

	nodeType := sbom.Node_NodeType_value[n.Type.String()]

	return &sbom.Node{
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
		ReleaseDate:        timestamppb.New(n.ReleaseDate),
		BuildDate:          timestamppb.New(n.BuildDate),
		ValidUntilDate:     timestamppb.New(n.ValidUntilDate),
		ExternalReferences: pbExtRefs,
		Identifiers:        entIdentifiersToProtobom(n.Edges.Identifiers),
		Hashes:             entHashesToProtobom(n.Edges.Hashes),
		PrimaryPurpose:     entPurposesToProtobom(n.Edges.PrimaryPurpose),
	}
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
