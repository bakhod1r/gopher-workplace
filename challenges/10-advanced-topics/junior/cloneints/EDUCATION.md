# A Copy That Owns Its Memory

## Intuition

A slice value is a small header pointing at an array. Handing it out hands out the array. To break the link you must allocate a new array and copy the elements into it.

## Approach

1. Allocate a destination of `len(s)`.
2. `copy` the elements across.
3. Return the destination.

## Solution

```go
// Clone returns a copy of s that shares no storage with s.
//
// Writes to the result must not be visible through s, and writes to s must
// not be visible through the result.
//
// Examples:
//
// 	Clone([]int{1, 2}) => []int{1, 2}
func Clone(s []int) []int {
	out := make([]int, len(s))
	copy(out, s)
	return out
}
```

## Walkthrough

`s := []int{1,2,3}` allocates one array. `c := s` yields a second header over the *same* array, so `s[0] = 99` shows up in `c`. After `make` + `copy` there are two arrays and the write is invisible.

## Pitfalls

- `s[:]` is not a copy.
- `copy(out, s)` into an out of length 0 copies nothing — the destination's length is what limits it.
