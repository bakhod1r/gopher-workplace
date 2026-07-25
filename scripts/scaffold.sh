#!/usr/bin/env bash
# scaffold.sh — materialize a puzzle slot from _template/, pre-loaded with the
# covered set so the author reuses earlier topics (spiral rule, GENERATION.md §5c).
#
# Usage:
#   scripts/scaffold.sh <level>/<NN-topic>/<MM-subtopic> <name> <Func> [Title]
# Example:
#   scripts/scaffold.sh junior/03-error-handling/01-error-values wrap Wrap "Error Wrap"
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CH="$ROOT/challenges"
TMPL="$CH/_template"

SLOT="${1:-}"; NAME="${2:-}"; FUNC="${3:-}"; TITLE="${4:-$NAME}"
if [[ -z "$SLOT" || -z "$NAME" || -z "$FUNC" ]]; then
  echo "usage: scripts/scaffold.sh <level>/<topic>/<subtopic> <name> <Func> [Title]" >&2
  exit 2
fi
SLOT="${SLOT#challenges/}"; SLOT="${SLOT%/}"
IFS='/' read -r LEVEL TOPIC SUBTOPIC <<< "$SLOT"
[[ -n "$LEVEL" && -n "$TOPIC" && -n "$SUBTOPIC" ]] || { echo "bad slot: $SLOT" >&2; exit 2; }

PKG="$(echo "$NAME" | tr -d '-' | tr '[:upper:]' '[:lower:]')"
DEST="$CH/$SLOT/$NAME"
if [[ -e "$DEST" ]]; then echo "refuse: $DEST already exists" >&2; exit 1; fi

# covered set for this slot (target included)
COVERED="$("$ROOT/scripts/coverage.sh" "$SLOT" || true)"

mkdir -p "$DEST"
sed_all() { sed -e "s|{{LEVEL}}|$LEVEL|g" -e "s|{{TOPIC}}|$TOPIC|g" \
  -e "s|{{SUBTOPIC}}|$SUBTOPIC|g" -e "s|{{NAME}}|$NAME|g" \
  -e "s|{{PKG}}|$PKG|g" -e "s|{{FUNC}}|$FUNC|g" -e "s|{{TITLE}}|$TITLE|g" "$1"; }

sed_all "$TMPL/go.mod.tmpl"        > "$DEST/go.mod"
cp       "$TMPL/Makefile"           "$DEST/Makefile"
sed_all "$TMPL/puzzle.go.tmpl"     > "$DEST/$NAME.go"
sed_all "$TMPL/puzzle_test.go.tmpl">  "$DEST/${NAME}_test.go"
sed_all "$TMPL/README.tmpl.md"     > "$DEST/README.md"

# inject covered set as an author reference block (COVERED.txt) — not shipped code
{
  echo "# Covered set for $SLOT — concepts this puzzle MAY use and SHOULD reuse."
  echo "# Target is the last line. Delete this file before finishing."
  echo "$COVERED"
} > "$DEST/COVERED.txt"

# drop the slot's .keep now that it is authored
rm -f "$CH/$SLOT/.keep"

echo "scaffolded: challenges/$SLOT/$NAME"
echo "  covered: $(echo "$COVERED" | grep -c . ) subtopics (see COVERED.txt)"
echo "  next: plant bug in $NAME.go, write red tests, fill README, then 'make -C $DEST verify' and delete COVERED.txt"
