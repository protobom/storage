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

define coverage-report
	@printf "${CYAN}+%s+${RESET}\n" $$(printf -- '-%.0s' {1..70})
	@printf "${CYAN}|%-70s|${RESET}\n" "    COVERAGE REPORT"
	@printf "${CYAN}+%s+${RESET}\n\n" $$(printf -- '-%.0s' {1..70})

	@go tool cover -func=coverage.out | \
	  awk -- '{ \
	    sub("github.com/protobom/storage/backends/ent/", "", $$1); \
	    percent = +$$3; sub("%", "", percent); \
	    if (percent < 50.00) color = "${RED}"; \
	    else if (percent < 80.00) color = "${YELLOW}"; \
	    else if (percent < 100.00) color = "${RESET}"; \
	    else color = "${GREEN}"; \
	    fmtstr = $$1 == "total:" ? "\n%s%s\t%s\t%s%s\n" : "%s%-24s %-36s %.1f%%%s\n"; \
	    printf fmtstr, color, $$1, $$2, $$3, "${RESET}" \
	  }'
endef

.PHONY: test-unit
test-unit: # Run unit tests
	@printf "Running tests for ${CYAN}backends/ent${RESET}..."
	@go test -failfast -v -coverprofile=coverage.out -covermode=atomic ./backends/ent/...
	@printf "${GREEN}DONE${RESET}\n\n"

	${call coverage-report}
