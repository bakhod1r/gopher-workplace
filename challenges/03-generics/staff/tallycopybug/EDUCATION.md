# The Update That Half Survives

## Intuition

Ranging by value copies each `Tally`. The copy's `Counts` field is a map header pointing at the original table, so the increment is visible; `Total` is a plain `int` inside the copy and is discarded at the end of the iteration.

## Approach

1. Iterate by index.
2. Increment the count and the total through `ts[i]`.

## Solution

```go
func BumpAll[T comparable](ts []Tally[T], v T) {
	for i := range ts {
		ts[i].Counts[v]++
		ts[i].Total++
	}
}

func Consistent[T comparable](t Tally[T]) bool {
	sum := 0
	for _, n := range t.Counts {
		sum += n
	}
	return sum == t.Total
}
```

## Walkthrough

After two `BumpAll` calls, `Counts["x"]` is 2 but `Total` is 0, and `Consistent` reports false — the struct's two halves have drifted apart.

## Pitfalls

- Making the slice `[]*Tally[T]` to dodge the issue instead of understanding it.
- Assuming a value copy is harmless because "nothing else changed" — reference fields make the failure partial and therefore hard to spot.
