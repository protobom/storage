//go:build ignore
// +build ignore

// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	// Parse the template file.
	tmpl := gen.MustParse(gen.NewTemplate("header").ParseFiles("template/header.tmpl"))

	if err := entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureExecQuery,
			gen.FeatureIntercept,
			gen.FeatureUpsert,
			gen.FeatureVersionedMigration,
		},
		Templates: []*gen.Template{tmpl},
	}); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
