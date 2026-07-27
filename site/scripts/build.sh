#!/usr/bin/env bash
# Build the static site. Local-runner-only: the in-browser wasm path is gone,
# so every puzzle runs against the real Go toolchain via site/cmd/localrunner.
# This step just regenerates problems.js from the challenges/ tree.
set -euo pipefail
cd "$(dirname "$0")/.."          # -> site/

echo "==> generating problems.js from challenges/"
( cd cmd/gencatalog && go run . -root ../../.. )

echo "==> done (start the runner with: cd site/cmd/localrunner && go run .)"
