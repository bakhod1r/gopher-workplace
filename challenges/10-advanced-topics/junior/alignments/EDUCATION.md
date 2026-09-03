# What Each Type Must Line Up On

## Intuition

Alignment is a hardware requirement the compiler enforces: a wide load wants an address divisible by its width. Every layout question — offsets, padding, struct size — follows from these numbers.

## Approach

1. Declare a variable of each type.
2. Return `unsafe.Alignof` of each.

## Solution

```go
import "unsafe"

// Alignments returns the alignment requirement of a byte, an int32, an
// int64 and a string.
//
// A type's alignment is the boundary its address must be a multiple of.
//
// Examples:
//
// 	Alignments() => 1, 4, 8, 8 on a 64-bit build
func Alignments() (b, i32, i64, s uintptr) {
	var (
		vb  byte
		v32 int32
		v64 int64
		vs  string
	)
	return unsafe.Alignof(vb), unsafe.Alignof(v32), unsafe.Alignof(v64), unsafe.Alignof(vs)
}
```

## Walkthrough

A string is a pointer plus a length, so its alignment is the pointer's — 8 on a 64-bit build, even though it occupies 16 bytes.

## Pitfalls

- Confusing `Alignof` with `Sizeof`; a string is 16 bytes and 8-aligned.
- Hard-coding the numbers, which are architecture-specific.
