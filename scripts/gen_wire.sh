#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

go install github.com/google/wire/cmd/wire@latest
wire ./cmd/web

echo "wire generated: cmd/web/wire_gen.go"
