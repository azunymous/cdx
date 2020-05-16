#!/usr/bin/env bash
set -euo pipefail
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_DIR="${SCRIPT_DIR}/../"

export KO_DOCKER_REPO=$1

ko publish ./cmd/cdx -B --tags `cdx tag latest -n cdx --head --fallback`,latest