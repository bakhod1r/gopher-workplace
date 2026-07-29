# Modulo with signed integers

## Intuition

Go's `%` keeps the sign of the dividend, so `-3 % 2 == -1`; test evenness with `== 0`, never `== 1`.

## Approach

1. `n % 2 == 0` → even, else odd.

## Solution

```go
func Parity(n int) string {
	if n%2 == 0 {
		return "even"
	}
	return "odd"
}
```

## Walkthrough

`Parity(4)`: 4 mod 2 is 0 → "even".

## Pitfalls

- `-3 % 2` is `-1`, not `1`; only `== 0` is reliable for parity.
- 0 is even.
