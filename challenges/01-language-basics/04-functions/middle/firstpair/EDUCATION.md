# Labeled break and continue

## The idea

Labels let `break`/`continue` target an enclosing loop, avoiding sentinel flags in nested iteration.

## Why it matters

Grid scans and multi-level searches read clearly with labels instead of goto-like flags.

## Watch out

- Only the FIRST matching pair should be returned; break immediately.
- Without the label, `break` exits just the inner loop.
