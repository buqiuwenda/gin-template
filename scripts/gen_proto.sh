#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

GOOGLEAPIS="${ROOT}/third_party/googleapis"
if [[ ! -f "${GOOGLEAPIS}/google/api/annotations.proto" ]]; then
  echo "==> fetching googleapis (one-time)..."
  mkdir -p "${ROOT}/third_party"
  git clone --depth 1 https://github.com/googleapis/googleapis.git "${GOOGLEAPIS}"
fi
PROTOC_INCLUDE="${PROTOC_INCLUDE:-${GOOGLEAPIS}}"

protoc \
  --proto_path="${PROTOC_INCLUDE}" \
  --proto_path=. \
  --go_out=. --go_opt=module=github.com/buqiuwenda/gin-template \
  --go-grpc_out=. --go-grpc_opt=module=github.com/buqiuwenda/gin-template \
  api/v1/user/user.proto

echo "proto generated under api/gen"
