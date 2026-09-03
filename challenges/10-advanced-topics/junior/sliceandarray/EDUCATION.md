# A Slice Header Versus An Array

## Intuition

This is the difference between the two types in one number. An array *is* its data; a slice is a small description of data that lives somewhere else — which is why passing an array copies everything and passing a slice copies three words.

## Approach

1. Declare `var a [8]int` and `var s []int`.
2. Return `unsafe.Sizeof` of each.

## Solution

```go
import "unsafe"

// Sizes returns the size of a [8]int array and of an []int slice header.
//
// The array's size is its contents; the slice's is three words, whatever it
// points at.
//
// Examples:
//
// 	Sizes() => 64, 24 on a 64-bit build
func Sizes() (arr, sl uintptr) {
	var (
		a [8]int
		s []int
	)
	return unsafe.Sizeof(a), unsafe.Sizeof(s)
}
```

## Walkthrough

On a 64-bit build an int is 8 bytes, so the array is 64 and the slice header is 24 — regardless of how many elements the slice has.

## Pitfalls

- Expecting `Sizeof` on a slice to grow with its length.
- Writing 64 and 24 as literals; both depend on the word size.
