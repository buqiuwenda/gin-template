#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

PROTOC_INCLUDE="${PROTOC_INCLUDE:-}"

protoc \
  ${PROTOC_INCLUDE:+--proto_path="$PROTOC_INCLUDE"} \
  --proto_path=. \
  --go_out=. --go_opt=module=github.com/buqiuwenda/gin-template \
  api/v1/user/user.proto

echo "proto generated under api/gen/go"
