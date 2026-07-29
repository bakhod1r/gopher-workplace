# Comparison operator precision

## Intuition

`>` and `>=` differ only at the boundary, but that single equal value is exactly where specifications are easy to misread.

## Approach

1. "Above" means strictly greater.
2. The bug uses `>=`, including values equal to the threshold.
3. Use `> t`.

## Solution

```go
func AboveThreshold(xs []int, t int) []int {
	var out []int
	for _, v := range xs {
		if v > t {
			out = append(out, v)
		}
	}
	return out
}
```

## Walkthrough

With `>=`, the 5s are wrongly included. `> t` keeps only 8, which is strictly above 5.

## Pitfalls

- "strictly greater" ⇒ `>`; "at least" ⇒ `>=`.
- Always test the exact boundary value.
