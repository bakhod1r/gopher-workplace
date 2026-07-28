#!/usr/bin/env bash
# Materialize the new-layout skeleton from challenges/HIERARCHY.md.
# Layout: challenges/<NN-topic>/<MM-subtopic>/  (level dirs added at gen time).
# Idempotent: only creates missing dirs + .keep. Never deletes.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
HIER="$ROOT/challenges/HIERARCHY.md"
OUT="$ROOT/challenges"

topic=""
while IFS= read -r line; do
  # topic heading:  "## 01 · language-basics  (roadmap 2)"
  if [[ "$line" =~ ^##[[:space:]]+([0-9]{2})[[:space:]]+·[[:space:]]+([a-z0-9-]+) ]]; then
    topic="${BASH_REMATCH[1]}-${BASH_REMATCH[2]}"
    continue
  fi
  # subtopic heading: "### 01-variables-and-constants"
  if [[ "$line" =~ ^###[[:space:]]+([0-9]{2}-[a-z0-9-]+) ]]; then
    sub="${BASH_REMATCH[1]}"
    dir="$OUT/$topic/$sub"
    mkdir -p "$dir"
    # Only drop a .keep when the subtopic holds no puzzle yet.
    if [[ -z "$(find "$dir" -name go.mod -print -quit)" ]]; then
      [[ -e "$dir/.keep" ]] || : > "$dir/.keep"
    else
      [[ -e "$dir/.keep" ]] && rm -f "$dir/.keep"
    fi
  fi
done < "$HIER"

echo "skeleton synced under $OUT"
find "$OUT" -maxdepth 2 -type d -name '[0-9][0-9]-*' | wc -l | xargs echo "topic+subtopic dirs:"
