# --------------------------------------------------------------
# SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
# SPDX-FileType: SOURCE
# SPDX-License-Identifier: Apache-2.0
# --------------------------------------------------------------
ENT_ROOT_DIR := ${abspath ${patsubst %/,%,${dir ${lastword ${MAKEFILE_LIST}}}}}

.PHONY: generate-ent describe-schemas
.SILENT: generate-ent describe-schemas

#@ Ent Utilities
generate-ent: # Generate ent database types from schema
	printf "Generating ${CYAN}ent${RESET} database types with ${CYAN}go generate${RESET}..."
	go generate ${ENT_ROOT_DIR}
	printf "${GREEN}DONE${RESET}\n"

describe-schemas: # Print descriptions of generated ent schema types
	go run -mod=mod entgo.io/ent/cmd/ent describe ${ENT_ROOT_DIR}/schema
