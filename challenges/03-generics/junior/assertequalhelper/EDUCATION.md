# Generic Test Helper

## Intuition

Because both arguments share one type parameter, comparing an `int` against a `string` fails to compile — the mistake is caught before the test ever runs.

## Approach

1. Call `t.Helper()`.
2. Report with `t.Errorf` and return `false` when the values differ.
3. Return `true` otherwise.

## Solution

```go
func Equal[T comparable](t testing.TB, got, want T) bool {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
		return false
	}
	return true
}
```

## Walkthrough

`Equal(t, 1, 2)` records `got 1, want 2` against the caller's line and returns `false`.

## Pitfalls

- Taking `*testing.T`, which shuts benchmarks out.
- Forgetting `t.Helper()`, so every failure points at the helper.
- Calling `t.Fatalf`, which stops the test and prevents further checks.
