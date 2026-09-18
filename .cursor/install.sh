#!/usr/bin/env bash
# Idempotent bootstrap for the LeetCode Go solutions repository.
# There are no external dependencies (no go.mod), so this simply verifies the
# Go toolchain and compiles every solution to catch build errors up front.
set -euo pipefail

echo "Go toolchain:"
go version

# Compile each solution's main.go. Directory names may contain spaces, so use a
# NUL-delimited loop and build each file individually (module-wide ./... does
# not work with spaces in package paths).
files=$(find . -path ./.git -prune -o -name '*.go' -print)
if [ -z "$files" ]; then
  echo "No Go source files found yet; nothing to build."
  exit 0
fi

while IFS= read -r -d '' f; do
  echo "Building: $f"
  go build -o /dev/null "$f"
done < <(find . -path ./.git -prune -o -name 'main.go' -print0)

echo "All solutions compiled successfully."
