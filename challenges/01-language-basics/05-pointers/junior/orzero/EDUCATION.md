# Optional values as pointers

## Intuition

A `*int` doubles as an optional int; the nil case maps to the zero value.

## Approach

1. Guard: `if p == nil { return 0 }`.
2. Otherwise `return *p`.

## Solution

```go
func DerefOrZero(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}
```

## Walkthrough

`DerefOrZero(nil)` returns `0` without touching memory; `DerefOrZero(&n)` with `n = 8` dereferences to `8`.

## Pitfalls

- Reading `*p` before the nil check panics.
- Returning 0 collapses nil and a real 0 — fine when that's intended.
