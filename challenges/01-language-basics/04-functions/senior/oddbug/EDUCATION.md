# Signed modulo and parity

## Intuition

Because `%` follows the dividend's sign, `n%2` is -1 for negative odds; only `!= 0` reliably detects oddness.

## Approach

1. In Go, `v % 2` is `-1` for negative odds, so `== 1` misses them.
2. Test `v % 2 != 0`.

## Solution

```go
func CountOdd(xs []int) int {
	c := 0
	for _, v := range xs {
		if v%2 != 0 {
			c++
		}
	}
	return c
}
```

## Walkthrough

`-3 % 2` is `-1`, so `== 1` fails to count negative odds. `!= 0` correctly identifies all odd numbers.

## Pitfalls

- `-3 % 2` is `-1`, so `== 1` fails.
- Prefer `v%2 != 0` for parity across all integers.
