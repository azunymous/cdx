#!/usr/bin/env bash
set -euo pipefail
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_DIR="${SCRIPT_DIR}/../"

COMMAND=${1:-build}

go ${COMMAND} -ldflags "-X main.version=`cdx tag latest -n cdx` -X main.buildTime=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.gitHash=`git rev-parse HEAD`" ./cmd/cdx/
