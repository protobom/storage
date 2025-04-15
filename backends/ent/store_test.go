// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/protobom/protobom/pkg/reader"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"

	"github.com/protobom/storage/backends/ent"
)

type storeSuite struct {
	suite.Suite
	*ent.Backend
	documents []*sbom.Document
}

func (ss *storeSuite) SetupSuite() {
	cwd, err := os.Getwd()
	ss.Require().NoError(err)

	rdr := reader.New()
	testdataDir := filepath.Join(cwd, "testdata")

	entries, err := os.ReadDir(testdataDir)
	ss.Require().NoError(err)

	for idx := range entries {
		document, err := rdr.ParseFile(filepath.Join(testdataDir, entries[idx].Name()))
		ss.Require().NoError(err)

		ss.documents = append(ss.documents, document)
	}
}

func (ss *storeSuite) BeforeTest(_suiteName, _testName string) {
	ss.Backend = ent.NewBackend(ent.WithDatabaseFile(":memory:"))
	ss.Require().NoError(ss.InitClient())
}

func (ss *storeSuite) AfterTest(_suiteName, _testName string) {
	ss.CloseClient()
}

func (ss *storeSuite) TestBackend_Store() {
	messages := [][]byte{}

	for _, document := range ss.documents {
		msg, err := proto.MarshalOptions{Deterministic: true}.Marshal(document)
		ss.Require().NoError(err)

		messages = append(messages, msg)

		ss.Require().NoError(ss.Store(document, nil))
	}

	results, err := ss.Backend.Ent().Document.
		Query().
		WithMetadata().
		WithNodeList().
		All(ss.Context())

	ss.Require().NoError(err)
	ss.Len(results, len(ss.documents))

	for _, result := range results {
		document := &sbom.Document{
			Metadata: result.Edges.Metadata.ProtoMessage,
			NodeList: result.Edges.NodeList.ProtoMessage,
		}

		msg, err := proto.MarshalOptions{Deterministic: true}.Marshal(document)
		ss.Require().NoError(err)

		ss.Contains(messages, msg)
	}
}

func TestStoreSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(storeSuite))
}
