# Sequential Merge

## Intuition

"Merge" in concurrency usually means interleaving with `select`, which is
nondeterministic. When the spec fixes the order — all of primary, then all
of standby — the simplest correct answer is two sequential drains. Reach
for `select` only when you genuinely must take whichever is ready first.

## Approach

1. Start with an empty, non-nil slice.
2. `range` over `primary`, appending.
3. `range` over `standby`, appending.
4. Return the slice.

## Solution

```go
func MergeFeeds(primary, standby <-chan int) []int {
	out := []int{}
	for id := range primary {
		out = append(out, id)
	}
	for id := range standby {
		out = append(out, id)
	}
	return out
}
```

## Walkthrough

`primary` holding `1, 2` and `standby` holding `3`: the first loop appends
`1` and `2` and ends at the primary close; the second appends `3` —
`[1 2 3]`.

## Pitfalls

- If `primary` is never closed, `standby` is never read — the whole report stalls on the first loop.
- A `select`-based merge here would produce a nondeterministic order and break the report format.
- Each feed must be closed by its own producer.
