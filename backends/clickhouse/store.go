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

// Store implements the storage.Storer interface.
//
// TODO: implemented in the next step (serialize the document blob and decompose
// the node list into wide nodes/edges batch inserts).
func (*Backend) Store(_ *sbom.Document, _ *storage.StoreOptions) error {
	return errNotImplemented
}
