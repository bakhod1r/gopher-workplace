# Seeding a running extremum

## Intuition

Seeding from the last element with a strict `<` comparison can return a later tie; seed from the first element so the earliest minimum wins.

## Approach

1. Seed the running min at `&xs[0]`, then scan.
2. The bug seeds at the last element, biasing ties and starting the scan wrong.

## Solution

```go
func MinPtr(xs []int) *int {
	min := &xs[0]
	for i := range xs {
		if xs[i] < *min {
			min = &xs[i]
		}
	}
	return min
}
```

## Walkthrough

Seeding at `&xs[len-1]` makes the last element the tentative min; a leading true minimum is still found by value, but the seed should be `&xs[0]` for a correct left-most result.

## Pitfalls

- Seeding `&xs[len-1]` with `<` returns the last min on ties.
- Seed `&xs[0]` to honour the earliest-tie rule.
