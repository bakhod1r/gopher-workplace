# Reducing several values

## Intuition

A running accumulator seeded from the first input generalises to any number of comparisons.

## Approach

1. Start `m := a`.
2. Raise `m` if `b` or `c` is larger.

## Solution

```go
func Max3(a, b, c int) int {
	m := a
	if b > m {
		m = b
	}
	if c > m {
		m = c
	}
	return m
}
```

## Walkthrough

`Max3(1,2,3)`: m starts 1, becomes 2, then 3.

## Pitfalls

- Seeding `m := 0` fails when every argument is negative.
- Use `>` (or `>=`) consistently.
