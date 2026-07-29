# Computing multiple aggregates in one pass

## Intuition

Accumulating several results in a single loop avoids re-scanning and keeps related outputs in sync.

## Approach

1. Range the variadic args.
2. For each even value, bump `count` and add to `total`.

## Solution

```go
func EvenStats(xs ...int) (count, total int) {
	for _, x := range xs {
		if x%2 == 0 {
			count++
			total += x
		}
	}
	return
}
```

## Walkthrough

`EvenStats(1,2,3,4)`: 2 and 4 are even, so count 2 and total 6.

## Pitfalls

- `n%2 != 0` for negative odds too (`-3%2 == -1`); evenness test still holds.
- Both accumulators start at 0.
