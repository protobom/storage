# --------------------------------------------------------------
# SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
# SPDX-FileType: SOURCE
# SPDX-License-Identifier: Apache-2.0
# --------------------------------------------------------------

# ANSI color escape codes
BOLD   := \033[1m
CYAN   := \033[36m
GREEN  := \033[32m
YELLOW := \033[33m
RESET  := \033[0m

.PHONY: help
help: # Display this help
	@awk 'BEGIN {FS = ":.*#"; printf "\n${YELLOW}Usage: make <target>${RESET}\n\n"} \
	  /^[a-zA-Z_0-9-]+:.*?#/ { printf "  ${CYAN}%-20s${RESET} %s\n", $$1, $$2 } \
	  /^#@/ { printf "\n${BOLD}%s${RESET}\n", substr($$0, 4) }' ${MAKEFILE_LIST} && echo

include internal/backends/ent/Makefile

#@ Development Tools
.PHONY: lint
lint: # Lint Golang code files
	golangci-lint run --verbose

.PHONY: lint-fix
lint-fix: # Fix linter findings
	golangci-lint run --fix --verbose

.PHONY: test-unit
.SILENT: test-unit
test-unit: # Run unit tests
	printf "Running tests for ${CYAN}backends/ent${RESET}..."
	go test -failfast -coverprofile=coverage.out -covermode=atomic ./backends/ent 1> /dev/null
	printf "${GREEN}DONE${RESET}\n\n"

	printf "${CYAN}"
	echo "+--------------------------------+"
	echo "|         COVERAGE REPORT        |"
	echo "+--------------------------------+"
	printf "${RESET}\n"

	go tool cover -func=coverage.out | \
	  awk '{ gsub(/^github.com\/protobom\/storage\/backends\/ent\//, "", $$1) } \
	    { printf ($$1 == "total:") ? "\n\t%s\t%s\t%s\n" : "%-24s %-36s %s\n", $$1, $$2, $$3 }'
