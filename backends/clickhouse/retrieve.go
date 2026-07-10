// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2026 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package clickhouse

import (
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"
)

// Retrieve implements the storage.Retriever interface.
//
// TODO: implemented in the next step (look up the document blob by id and
// unmarshal it).
func (*Backend) Retrieve(_ string, _ *storage.RetrieveOptions) (*sbom.Document, error) {
	return nil, errNotImplemented
}
