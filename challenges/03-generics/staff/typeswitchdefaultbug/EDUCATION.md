# The Default Branch That Eats The Value

## Intuition

A type switch over `any(v)` is the only way a generic function can special-case a concrete type, but every type it does not name must be passed through untouched. Returning the zero value there discards the input.

## Approach

1. Switch on `any(v)`.
2. Normalise the cases you know how to normalise.
3. Return `v` unchanged in the default branch.

## Solution

```go
func Normalize[T any](v T) T {
	switch x := any(v).(type) {
	case string:
		n := strings.ToLower(strings.TrimSpace(x))
		return any(n).(T)
	default:
		return v
	}
}
```

## Walkthrough

`Normalize(42)` matches no case, falls into `default`, and returns `0`.

## Pitfalls

- Writing `return zero` as a placeholder and never coming back to it.
- Assuming the compiler warns about an unreachable-looking branch — a type switch on `any` has no such check.
