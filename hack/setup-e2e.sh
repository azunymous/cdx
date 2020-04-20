#!/usr/bin/env bash
set -euo pipefail
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_DIR="${SCRIPT_DIR}/../"

cd ${PROJECT_DIR}

echo "This installs cdx. The end to end tests assume cdx is on your \$PATH"
go install ./cmd/cdx/