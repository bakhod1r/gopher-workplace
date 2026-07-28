# Overlapping copy direction

## The idea

`copy` handles overlapping ranges as if through a temporary, but the direction you choose determines whether you compact or duplicate.

## Why it matters

Left/right shift confusion corrupts ring buffers, gap arrays, and in-place compaction.

## Watch out

- Drop-first shifts LEFT: `copy(xs, xs[1:])`.
- Insert shifts RIGHT: `copy(xs[i+1:], xs[i:])`.
