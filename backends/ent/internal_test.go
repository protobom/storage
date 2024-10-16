// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent

import (
	"context"

	"github.com/protobom/storage/internal/backends/ent"
)

func (backend *Backend) Client() *ent.Client {
	return backend.client
}

func (backend *Backend) Context() context.Context {
	return backend.ctx
}
