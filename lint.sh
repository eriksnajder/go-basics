#!/usr/bin/env bash
set -euo pipefail

ARTIFACTS=".artifacts"
mkdir -p "$ARTIFACTS"

echo "==> go mod tidy check"
go mod tidy
git diff --exit-code go.mod go.sum || {
  echo "go.mod/go.sum not tidy"
  exit 1
}

echo "==> build"
go build ./...

echo "==> vet"
go vet ./...

echo "==> format check"
UNFORMATTED=$(gofmt -l .)
if [ -n "$UNFORMATTED" ]; then
  echo "Unformatted files:"
  echo "$UNFORMATTED"
  exit 1
fi

echo "==> lint"
golangci-lint run

echo "==> test with race + coverage"
go test ./... -race -covermode=atomic -coverprofile="$ARTIFACTS/coverage.out"

echo "==> done"