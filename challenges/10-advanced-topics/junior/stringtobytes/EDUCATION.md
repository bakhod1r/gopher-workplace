# Read A String's Bytes Without Copying

## Intuition

Strings are immutable by contract, not by the type system. `unsafe.Slice` gives you a mutable-looking view of immutable bytes, which is exactly why the result must stay read-only.

## Approach

1. Return nil for the empty string.
2. `unsafe.Slice(unsafe.StringData(s), len(s))`.

## Solution

```go
import "unsafe"

// Bytes returns a read-only byte view of s.
//
// The bytes belong to the string and may live in read-only memory, so the
// result must never be written to.
//
// Examples:
//
// 	Bytes("hi") => []byte("hi"), sharing the string's bytes
func Bytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
```

## Walkthrough

`unsafe.Slice(p, 4)` yields a slice of length and capacity 4 over the string's bytes — no allocation, and no room for `append` to overwrite anything after it.

## Pitfalls

- Writing through the result — for a string literal that is a segmentation fault, and for any other string it corrupts a value the whole program believes is immutable.
- Passing the result to a function that appends to it.
