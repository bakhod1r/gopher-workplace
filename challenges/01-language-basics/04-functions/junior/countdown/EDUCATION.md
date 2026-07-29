# The classic three-clause for loop

## Intuition

`for init; cond; post` is Go's only loop keyword; it covers counting up, down, and while-style conditions.

## Approach

1. Loop `i` from `n` down to 1.
2. Append each to the result.

## Solution

```go
func Countdown(n int) []int {
	var out []int
	for i := n; i >= 1; i-- {
		out = append(out, i)
	}
	return out
}
```

## Walkthrough

`Countdown(3)` appends 3, 2, 1.

## Pitfalls

- Off-by-one: use `i >= 1` to include 1 and stop before 0.
- Return a non-nil empty slice for `n <= 0` (len 0 is enough here).
