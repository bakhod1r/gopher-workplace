#!/usr/bin/env bash
# Production-style local server: gzip-serves the wasm (Content-Encoding: gzip).
set -euo pipefail
cd "$(dirname "$0")/.."          # -> site/
PORT="${1:-8080}"
( cd server && go build -o server . )
PORT="$PORT" ROOT="web" exec ./server/server
