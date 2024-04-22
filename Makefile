# ------------------------------------------------------------------------
# SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
# SPDX-FileName: Makefile
# SPDX-FileType: SOURCE
# SPDX-License-Identifier: Apache-2.0
# ------------------------------------------------------------------------
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# ------------------------------------------------------------------------
BASH := ${shell type -p bash}
SHELL := ${BASH}

# ANSI color escape codes
BOLD   := \033[1m
CYAN   := \033[36m
YELLOW := \033[33m
RESET  := \033[0m

.PHONY: help
help: # Display this help
	@awk 'BEGIN {FS = ":.*#"; printf "\n${YELLOW}Usage: make <target>${RESET}\n\n"} \
	  /^[a-zA-Z_0-9-]+:.*?#/ { printf "  ${CYAN}%-20s${RESET} %s\n", $$1, $$2 } \
	  /^#@/ { printf "\n${BOLD}%s${RESET}\n", substr($$0, 4) }' ${MAKEFILE_LIST} && echo

include backends/ent/Makefile

#@ Development Tools
lint: # Lint Golang code files
	golangci-lint run --verbose

lint-fix: # Fix linter findings
	golangci-lint run --fix --verbose
