# Did That Append Reallocate

## Intuition

Whether `append` copied is invisible in the result's contents — the length changes either way. The capacity is the observable that separates "wrote into spare room" from "allocated a bigger array".

## Approach

1. Save `cap(s)`.
2. Append.
3. Return the result and whether the capacity differs.

## Solution

```go
// Appended appends v to s and reports whether the append had to grow the
// capacity.
//
// Growing means a new array was allocated and the old contents copied.
//
// Examples:
//
// 	Appended(make([]int, 0, 4), 1) => a one-element slice, false
func Appended(s []int, v int) ([]int, bool) {
	before := cap(s)
	out := append(s, v)
	return out, cap(out) != before
}
```

## Walkthrough

A slice with len 0 and cap 4 has room, so the capacity stays 4 and `grew` is false. A slice with len 1 and cap 1 has none, so `append` allocates a larger array and the capacity changes.

## Pitfalls

- Comparing lengths, which always differ by one.
- Comparing `cap(s)` after the append with itself — `s` still holds the old header.
