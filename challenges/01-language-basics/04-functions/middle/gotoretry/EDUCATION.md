# The goto statement

## Intuition

`goto` transfers control to a label in the same function; it cannot jump into a block or over variable declarations.

## Approach

1. Label a point and `goto` it to loop.
2. Add elements while under the limit and within bounds.

## Solution

```go
func SumUntil(xs []int, limit int) (sum, used int) {
loop:
	if used < len(xs) && sum < limit {
		sum += xs[used]
		used++
		goto loop
	}
	return
}
```

## Walkthrough

For `[1 2 3 4]`, limit 5: add 1, 2, 3 (sum 6 ≥ 5), stop at used 3.

## Pitfalls

- `goto` can't jump over a variable's declaration into its scope.
- Prefer `for`; use `goto` only when it genuinely reads better.
