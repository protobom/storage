// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

//go:build integration

package ent_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/protobom/protobom/pkg/reader"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"google.golang.org/protobuf/proto"

	"github.com/protobom/storage/backends/ent"
)

// PostgreSQL Integration Test Suite
// Run with: go test -tags=integration -v ./backends/ent -run TestPostgresSuite
type postgresSuite struct {
	suite.Suite
	backend   *ent.Backend
	container testcontainers.Container
	documents []*sbom.Document
}

func (ps *postgresSuite) SetupSuite() {
	ctx := context.Background()

	// Start PostgreSQL container
	pgContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpass"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second)),
	)
	ps.Require().NoError(err)
	ps.container = pgContainer

	// Get connection string
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	ps.Require().NoError(err)

	// Initialize backend
	ps.backend = ent.NewBackend(ent.WithPostgresConnection(connStr))
	ps.Require().NoError(ps.backend.InitClient())

	// Load test documents
	cwd, err := os.Getwd()
	ps.Require().NoError(err)

	rdr := reader.New()
	testdataDir := filepath.Join(cwd, "testdata")

	entries, err := os.ReadDir(testdataDir)
	ps.Require().NoError(err)

	for idx := range entries {
		document, err := rdr.ParseFile(filepath.Join(testdataDir, entries[idx].Name()))
		ps.Require().NoError(err)
		ps.documents = append(ps.documents, document)
	}
}

func (ps *postgresSuite) TearDownSuite() {
	if ps.backend != nil {
		ps.backend.CloseClient()
	}
	if ps.container != nil {
		ctx := context.Background()
		ps.Require().NoError(testcontainers.TerminateContainer(ps.container))
	}
}

func (ps *postgresSuite) TestPostgresBackendStore() {
	for idx := range ps.documents {
		ps.T().Run(ps.documents[idx].GetMetadata().GetId(), func(t *testing.T) {
			err := ps.backend.Store(ps.documents[idx], nil)
			ps.Require().NoError(err)

			retrieved, err := ps.backend.Retrieve(ps.documents[idx].GetMetadata().GetId(), nil)
			ps.Require().NoError(err)

			// Remove source data URI for comparison
			retrieved.GetMetadata().GetSourceData().Uri = nil
			ps.documents[idx].GetMetadata().GetSourceData().Uri = nil

			ps.Require().True(proto.Equal(ps.documents[idx], retrieved))
		})
	}
}

func (ps *postgresSuite) TestPostgresDuplicateStorage() {
	// Test that storing the same document twice doesn't fail (NodeList fix)
	document := ps.documents[0]
	
	// Store first time
	err := ps.backend.Store(document, nil)
	ps.Require().NoError(err)

	// Store second time - should not fail
	err = ps.backend.Store(document, nil)
	ps.Require().NoError(err)

	// Verify document can still be retrieved
	retrieved, err := ps.backend.Retrieve(document.GetMetadata().GetId(), nil)
	ps.Require().NoError(err)
	ps.Require().NotNil(retrieved)
}

func (ps *postgresSuite) TestPostgresConfiguration() {
	// Test different configuration methods
	ctx := context.Background()
	connStr, err := ps.container.(*postgres.PostgresContainer).ConnectionString(ctx, "sslmode=disable")
	ps.Require().NoError(err)

	// Method 1: WithPostgresConnection
	backend1 := ent.NewBackend(ent.WithPostgresConnection(connStr))
	ps.Equal(ent.PostgresDialect, backend1.Options.Dialect)
	ps.Equal(connStr, backend1.Options.DatabaseURL)

	// Method 2: Separate options
	backend2 := ent.NewBackend(
		ent.WithDialect(ent.PostgresDialect),
		ent.WithDatabaseURL(connStr),
	)
	ps.Equal(ent.PostgresDialect, backend2.Options.Dialect)
	ps.Equal(connStr, backend2.Options.DatabaseURL)

	// Test initialization
	ps.Require().NoError(backend1.InitClient())
	ps.Require().NoError(backend2.InitClient())

	backend1.CloseClient()
	backend2.CloseClient()
}

func TestPostgresSuite(t *testing.T) {
	suite.Run(t, new(postgresSuite))
}
