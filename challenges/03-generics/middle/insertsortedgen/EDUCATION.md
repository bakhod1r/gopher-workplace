# Insert Keeping Order

## Intuition

Choosing the upper bound rather than the lower one keeps insertion stable with respect to arrival time, which is what a leaderboard wants.

## Approach

1. Clone and normalise nil.
2. Advance `i` past every element `<= v`.
3. Return `slices.Insert(out, i, v)`.

## Solution

```go
func InsertSorted[T cmp.Ordered](s []T, v T) []T {
	out := slices.Clone(s)
	if out == nil {
		out = []T{}
	}
	i := 0
	for i < len(out) && out[i] <= v {
		i++
	}
	return slices.Insert(out, i, v)
}
```

## Walkthrough

`InsertSorted([]int{1,2}, 2)` walks past both elements and appends `2` at index 2.

## Pitfalls

- Using `<`, which inserts before equal elements.
- Inserting into the caller's slice.
- Appending then sorting, which throws away the cheap insert.
