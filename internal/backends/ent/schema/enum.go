// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2025 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package schema

import "google.golang.org/protobuf/reflect/protoreflect"

// enumValues returns the values of a protobuf enum type deterministically, preserving their order.
func enumValues(enum protoreflect.Enum) []string {
	values := []string{}

	enumValues := enum.Descriptor().Values()
	for idx := range enumValues.Len() {
		values = append(values, string(enumValues.Get(idx).Name()))
	}

	return values
}
