# Clearing pointer slices without leaking

## Intuition

Reducing length leaves element pointers in the backing array; for pointer/interface elements you must nil them (or use `clear`) so the referents can be collected.

## Approach

1. `s[:0]` alone leaves the pointers live in the backing array, leaking them.
2. Nil each slot first, then reslice to length 0.

## Solution

```go
func Clear(s []*int) []*int {
	for i := range s {
		s[i] = nil
	}
	return s[:0]
}
```

## Walkthrough

Reslicing keeps the element pointers reachable via capacity. Setting each to nil before `s[:0]` lets them be collected.

## Pitfalls

- `s[:0]` keeps pointers in the array's spare capacity.
- `clear(s)` (Go 1.21+) zeroes elements; then re-slice.
