# Order-preserving dedup

## The idea

Keep a set of already-emitted values; emit an element the **first** time you see
it (when it is NOT in the set):

```go
if _, ok := seen[x]; !ok {
	out = append(out, x); seen[x] = struct{}{}
}
```

## Why it matters

Unlike sorting-based dedup, this preserves first-seen order — needed for stable
lists, ordered menus, and log deduplication. The condition polarity (`!ok`) is the
crux.

## Watch out

- Emit on absence, then record; recording first would suppress everything.
- The map gives O(1) membership; a slice scan would be O(n²).
- `slices.Compact` only removes *consecutive* dups — different semantics.
