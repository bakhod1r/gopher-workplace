# Reserve Once, Fill Later

## Intuition

Growing is allocate-and-copy. Doing it once, to the size you actually need, beats doing it repeatedly in doubling steps.

## Approach

1. Return `s` when `cap(s)` is already at least `n`.
2. Otherwise `make` a slice with `len(s)` elements and capacity `n`, copy, return it.

## Solution

```go
func GrowTo(s []int, n int) []int {
	if n <= cap(s) {
		return s
	}
	out := make([]int, len(s), n)
	copy(out, s)
	return out
}
```

## Walkthrough

`make([]int, len(s), n)` keeps the length so the copied elements stay addressable, while the capacity jumps straight to `n`.

## Pitfalls

- `make([]int, n)`, which changes the length and appends zeros to the caller's data.
- Comparing against `len(s)` instead of `cap(s)`, reallocating a slice that already had room.
- Assuming the returned slice still aliases `s` — after growth it does not.
