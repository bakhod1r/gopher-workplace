#!/usr/bin/env bash
# mirror-roadmap.sh — mirror the roadmap source dir tree into challenges/,
# recursively at every depth, EXCLUDING 01-introduction-to-go.
#
# - Top-level topics are renumbered from 01 in source order after the excluded
#   intro (source 02 -> 01, 03 -> 02, ... 18 -> 17). Deeper dirs keep their
#   source numbering verbatim.
# - Dirs only (no .md content). Each LEAF dir (no child dir in source) gets a
#   .keep, unless the destination already holds a puzzle (go.mod).
# - Idempotent, never deletes an existing puzzle.
set -euo pipefail

SRC="${SRC:-/Users/mrb/Desktop/SeniorProject/roadmap/Programming/languages/golang}"
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
OUT="$ROOT/challenges"

[[ -d "$SRC" ]] || { echo "source not found: $SRC" >&2; exit 1; }

python3 - "$SRC" "$OUT" <<'PY'
import os, sys

SRC, OUT = sys.argv[1], sys.argv[2]
EXCLUDE_TOP = "01-introduction-to-go"

tops = sorted(d for d in os.listdir(SRC)
              if os.path.isdir(os.path.join(SRC, d)) and d != EXCLUDE_TOP)
remap = {}
for i, base in enumerate(tops, start=1):
    slug = base.split("-", 1)[1]           # drop leading NN-
    remap[base] = "%02d-%s" % (i, slug)

def has_child_dir(d):
    return any(os.path.isdir(os.path.join(d, x)) for x in os.listdir(d))

def has_gomod(d):
    for _dp, _dn, fs in os.walk(d):
        if "go.mod" in fs:
            return True
    return False

made = keeps = 0
for dirpath, dirnames, _files in os.walk(SRC):
    rel = os.path.relpath(dirpath, SRC)
    if rel == ".":
        continue
    top = rel.split(os.sep, 1)[0]
    if top == EXCLUDE_TOP:
        dirnames[:] = []
        continue
    rest = rel[len(top):]                    # leading sep + remainder, or ""
    dest = os.path.join(OUT, remap[top] + rest)
    if not os.path.isdir(dest):
        os.makedirs(dest); made += 1
    # leaf in source -> candidate for .keep
    if not has_child_dir(dirpath):
        if not (os.path.isdir(dest) and has_gomod(dest)):
            k = os.path.join(dest, ".keep")
            if not os.path.exists(k):
                open(k, "w").close(); keeps += 1

print("mirrored: created %d dirs, %d new .keep leaves" % (made, keeps))
print("top-level renumber map:")
for base in tops:
    print("  %s -> %s" % (base, remap[base]))
PY
