# Prefix-based higher-order operations

## Intuition

TakeWhile/DropWhile depend on order and stop at the boundary, unlike Filter which scans everything.

## Approach

1. Range until the predicate first fails.
2. Break at that point.

## Solution

```go
func TakeWhile(xs []int, pred func(int) bool) []int {
	var out []int
	for _, v := range xs {
		if !pred(v) {
			break
		}
		out = append(out, v)
	}
	return out
}
```

## Walkthrough

`[2 4 5 6]` with even: 2 and 4 pass, 5 fails and stops the take.

## Pitfalls

- `break` on the first failure — do not continue scanning like Filter.
- Empty result when the first element already fails.
