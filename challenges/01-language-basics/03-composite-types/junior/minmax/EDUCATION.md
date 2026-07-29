# Seed-and-scan for extremes

## Intuition

Don't seed min/max with 0 or huge sentinels — seed with the **first element**,
then scan the rest:

```go
if len(xs) == 0 { return 0, 0, false }
mn, mx := xs[0], xs[0]
for _, x := range xs[1:] { if x < mn { mn = x }; if x > mx { mx = x } }
```

## Approach

1. If xs is empty, return 0, 0, false.
2. Seed both min and max with xs[0].
3. Scan xs[1:], lowering min or raising max as needed.
4. Return min, max, true.

## Solution

```go
func MinMax(xs []int) (min, max int, ok bool) {
	if len(xs) == 0 {
		return 0, 0, false
	}
	min, max = xs[0], xs[0]
	for _, x := range xs[1:] {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return min, max, true
}
```

## Walkthrough

MinMax([3,1,4,1,5]): seed min=max=3; see 1 -> min=1; 4 -> max=4; 1 -> no change; 5 -> max=5. Return 1,5,true.

## Pitfalls

- Guard empty before indexing `xs[0]`.
- Seeding with 0 is a classic bug for all-negative or all-large inputs.
- `slices.Min`/`slices.Max` (Go 1.21+) panic on empty — the `ok` form is safer.
