# Pruning That Walks A Stale Key List

## Intuition

`maps.DeleteFunc` is a whole-map sweep. Driving it from inside a loop over a pre-taken key list means the very first iteration removes every doomed entry; the remaining doomed keys are then skipped as absent, so they never reach the report — while surviving keys are recorded as removed because the loop credits the sweep to whichever key it happens to be on.

## Approach

1. Snapshot and sort the keys so the report is deterministic.
2. For each key, test its own value against the limit.
3. Record it and `delete` just that one entry.

## Solution

```go
func Prune(m map[string]int, limit int) []string {
	keys := slices.Sorted(maps.Keys(m))
	removed := make([]string, 0)
	for _, k := range keys {
		if m[k] < limit {
			removed = append(removed, k)
			delete(m, k)
		}
	}
	return removed
}
```

## Walkthrough

`{a:1, b:5, c:2}` with limit 3: the `a` iteration sweeps out both `a` and `c`; `b` survives the sweep but is appended anyway; `c` is now absent and skipped. The report reads `[a b]` instead of `[a c]`.

## Pitfalls

- Deleting from a map while ranging it is actually legal in Go — the trap here is the bulk sweep, not the mutation.
- Returning `maps.Keys(m)` before and after and diffing them, which allocates two full key sets to answer a question one pass already knew.
