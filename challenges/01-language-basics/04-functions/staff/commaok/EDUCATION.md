# The comma-ok idiom for maps

## The idea

`v, ok := m[k]` is the only way to tell a stored zero from a missing key; the single-value form collapses them.

## Why it matters

Treating absent as zero (or vice versa) is a real bug in caches, configs, and counters.

## Watch out

- `m[k]` alone can't distinguish missing from zero-valued.
- Always use `v, ok := m[k]` when absence matters.
