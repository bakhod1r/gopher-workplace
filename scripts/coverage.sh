#!/usr/bin/env bash
# coverage.sh — list the "covered set" for a puzzle slot.
#
# The learning path is one total order:  level -> topic(NN) -> subtopic(MM),
# with junior < middle < senior < staff. Given a slot, print every subtopic at
# or before it (the concepts a puzzle there is allowed to use). See
# challenges/GENERATION.md.
#
# Usage:
#   scripts/coverage.sh <level>/<NN-topic>/<MM-subtopic>
#   scripts/coverage.sh middle/03-error-handling/02-...
#   scripts/coverage.sh --count <slot>     # just the number
#
# Run from repo root (or anywhere; paths resolve to challenges/).
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CH="$ROOT/challenges"

COUNT_ONLY=0
if [[ "${1:-}" == "--count" ]]; then COUNT_ONLY=1; shift; fi

TARGET="${1:-}"
if [[ -z "$TARGET" ]]; then
  echo "usage: scripts/coverage.sh [--count] <level>/<NN-topic>/<MM-subtopic>" >&2
  exit 2
fi
TARGET="${TARGET#challenges/}"
TARGET="${TARGET%/}"

python3 - "$CH" "$TARGET" "$COUNT_ONLY" <<'PY'
import os, sys

ch, target, count_only = sys.argv[1], sys.argv[2], sys.argv[3] == "1"
LEVELS = {"junior": 0, "middle": 1, "senior": 2, "staff": 3}

def num(name):
    # leading NN- prefix -> int for ordering; fall back to name
    p = name.split("-", 1)[0]
    return int(p) if p.isdigit() else 1 << 30

def key(level, topic, sub):
    return (LEVELS[level], num(topic), num(sub))

# enumerate every subtopic slot: challenges/<level>/<topic>/<sub>
slots = []
for level in LEVELS:
    ld = os.path.join(ch, level)
    if not os.path.isdir(ld):
        continue
    for topic in sorted(os.listdir(ld)):
        td = os.path.join(ld, topic)
        if not os.path.isdir(td):
            continue
        for sub in sorted(os.listdir(td)):
            sd = os.path.join(td, sub)
            if os.path.isdir(sd):
                slots.append((level, topic, sub))

parts = target.split("/")
if len(parts) != 3:
    sys.exit(f"slot must be <level>/<topic>/<subtopic>, got: {target}")
tl, tt, ts = parts
if tl not in LEVELS:
    sys.exit(f"unknown level: {tl}")
tkey = key(tl, tt, ts)

covered = sorted([s for s in slots if key(*s) <= tkey], key=lambda s: key(*s))
if not any(key(*s) == tkey for s in slots):
    print(f"# warning: target slot not found on disk: {target}", file=sys.stderr)

if count_only:
    print(len(covered))
else:
    for level, topic, sub in covered:
        marker = "  <-- TARGET" if (level, topic, sub) == (tl, tt, ts) else ""
        print(f"{level}/{topic}/{sub}{marker}")
    print(f"# {len(covered)} subtopics covered (including target)", file=sys.stderr)
PY
