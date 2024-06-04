#!/usr/bin/env bash

set -euo pipefail

readonly ENT_ROOT_DIR="${GITHUB_WORKSPACE}/internal/backends/ent"

# Install GNU make if not installed
if ! command -v make &> /dev/null; then
  (apt-get update && apt-get install --yes make) &> /dev/null
fi

# Generate ent schemas and code
make generate-ent

git diff --exit-code -- "${ENT_ROOT_DIR}"/**/*.go || {
  echo "The ent schemas and database types are not up to date." \
    "Check the docs and run 'make generate-ent'"

  exit 1
}
