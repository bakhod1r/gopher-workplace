# Coalesce

## Intuition

Layered defaults are a fold over the arguments: scan in priority order and stop at the first meaningful value.

## Approach

1. Declare `var zero T`.
2. Return the first argument not equal to `zero`.
3. Return `zero` after the loop.

## Solution

```go
func Coalesce[T comparable](vals ...T) T {
	var zero T
	for _, v := range vals {
		if v != zero {
			return v
		}
	}
	return zero
}
```

## Walkthrough

`Coalesce("", "a", "b")` skips the empty string, returns `"a"`, and never examines `"b"`.

## Pitfalls

- Returning the last non-zero value instead of the first.
- Panicking or returning nothing when all arguments are zero.
- Treating a deliberately-set zero as unset — that is inherent to this API, and why `GetOr` exists too.
