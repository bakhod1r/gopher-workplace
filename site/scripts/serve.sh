#!/usr/bin/env bash
# Serve the static site locally.
set -euo pipefail
cd "$(dirname "$0")/../web"
PORT="${1:-8080}"
echo "serving http://localhost:$PORT  (Ctrl-C to stop)"
exec python3 -m http.server "$PORT"
