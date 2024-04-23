// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package options

type StoreOptions struct {
	// NoClobber ensures documents with the same ID are never overwritten
	NoClobber bool

	// BackendOptions is a field to pipe system-specific options to the
	// modules implementing the storage backend interface
	BackendOptions interface{}
}

type RetrieveOptions struct {
	// BackendOptions is a field to pipe system-specific options to the
	// modules implementing the storage backend interface
	BackendOptions interface{}
}
