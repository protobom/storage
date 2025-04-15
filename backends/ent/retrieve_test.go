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
	"github.com/stretchr/testify/suite"

	"github.com/protobom/storage/backends/ent"
)

type retrieveSuite struct {
	suite.Suite
	*ent.Backend
}

// SetupSuite runs before the tests in the suite are run.
func (rs *retrieveSuite) SetupSuite() {
	rs.Backend = ent.NewBackend(ent.WithDatabaseFile(":memory:"))
	rs.Require().NoError(rs.InitClient())

	cwd, err := os.Getwd()
	rs.Require().NoError(err)

	rdr := reader.New()
	testdataDir := filepath.Join(cwd, "testdata")

	entries, err := os.ReadDir(testdataDir)
	rs.Require().NoError(err)

	for idx := range entries {
		document, err := rdr.ParseFile(filepath.Join(testdataDir, entries[idx].Name()))
		rs.Require().NoError(err)

		rs.Require().NoError(rs.Store(document, nil))
	}
}

// TearDownSuite runs after all the tests in the suite have been run.
func (rs *retrieveSuite) TearDownSuite() {
	rs.CloseClient()
}

func (rs *retrieveSuite) TestBackend_GetExternalReferencesByDocumentID() {
	// juice-shop-11.1.2.cdx.json serial number
	id := "urn:uuid:1f860713-54b9-4253-ba5a-9554851904af"

	for _, data := range []struct {
		name     string
		errorMsg string
		types    []string
		expected int
	}{
		{
			name:     "issue-tracker",
			types:    []string{"ISSUE_TRACKER"},
			expected: 711,
		},
		{
			name:     "vcs",
			types:    []string{"VCS"},
			expected: 730,
		},
		{
			name:     "website",
			types:    []string{"WEBSITE"},
			expected: 714,
		},
		{
			name:     "all types",
			types:    []string{},
			expected: 2155,
		},
		{
			name:     "invalid-type",
			types:    []string{"INVALID"},
			errorMsg: `externalreference: invalid enum value for type field: "INVALID"`,
		},
	} {
		rs.Run(data.name, func() {
			extRefs, err := rs.GetExternalReferencesByDocumentID(id, data.types...)

			if data.errorMsg != "" {
				rs.Require().Error(err)
				rs.Require().Equal(data.errorMsg, err.Error())
			} else {
				rs.Require().NoError(err)
			}

			rs.Len(extRefs, data.expected)
		})
	}
}

func TestRetrieveSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(retrieveSuite))
}
