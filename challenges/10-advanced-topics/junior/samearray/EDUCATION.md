# Do These Two Slices Share Storage

## Intuition

A slice header's first word is the address of its first element. Two slices alias when that address is the same, which is a pointer comparison the language will not let you write without `unsafe`.

## Approach

1. Return false when either slice is empty.
2. Compare `unsafe.SliceData(a)` with `unsafe.SliceData(b)`.

## Solution

```go
import "unsafe"

// SameArray reports whether a and b start at the same element of the
// same backing array.
//
// Comparing slices with == is not allowed; comparing their data pointers
// is.
//
// Examples:
//
// 	s := []int{1, 2}; SameArray(s, s[:1]) => true
func SameArray(a, b []int) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	return unsafe.SliceData(a) == unsafe.SliceData(b)
}
```

## Walkthrough

`s` and `s[:2]` both point at `&s[0]`, so the comparison is true. `s[1:]` points at `&s[1]`, so it is false — this is a same-start test, not an overlap test.

## Pitfalls

- Treating this as an overlap check; overlapping slices with different starts report false.
- Skipping the empty guard, where the data pointer of a nil slice is nil and two nils would compare equal.
