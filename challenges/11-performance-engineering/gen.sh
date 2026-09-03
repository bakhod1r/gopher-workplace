#!/bin/bash
# Scaffold 11-performance-engineering puzzles.
# Usage: bash gen.sh <level> <name> [<name> ...]
# Creates <level>/<name>/{go.mod,Makefile} and nothing else; the Go source,
# tests, README.md and EDUCATION.md are authored per puzzle.
set -e
cd "$(dirname "$0")"

level="$1"; shift || true
case "$level" in
junior | middle | senior | staff) ;;
*)
	echo "usage: bash gen.sh <junior|middle|senior|staff> <name>..." >&2
	exit 1
	;;
esac

for name in "$@"; do
	dir="$level/$name"
	mkdir -p "$dir"
	cat >"$dir/go.mod" <<EOF
module github.com/gopher-workplace/challenges/11-performance-engineering/$level/$name

go 1.26
EOF
	cp ../03-generics/senior/clipbug/Makefile "$dir/Makefile"
	echo "scaffolded $dir"
done
