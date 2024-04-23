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
	if err := entc.Generate("./schema", &gen.Config{
		Features:  []gen.Feature{gen.FeatureUpsert},
		Templates: []*gen.Template{gen.MustParse(gen.NewTemplate("header").ParseFiles("template/header.tmpl"))},
	}); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
