// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package storage

import (
	"github.com/bom-squad/protobom/pkg/sbom"

	"github.com/protobom/storage/pkg/options"
)

type Backend interface {
	Store(*sbom.Document, *options.StoreOptions) error
	Retrieve(string, *options.RetrieveOptions) (*sbom.Document, error)
}
