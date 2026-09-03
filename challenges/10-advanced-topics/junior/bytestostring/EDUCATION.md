# A String View Of Bytes You Own

## Intuition

A string and a byte slice have the same bytes underneath; only the promise differs. `unsafe.String` makes the string header point at bytes you already have — and hands you the job of keeping that promise.

## Approach

1. Return `""` for an empty input.
2. `unsafe.String(unsafe.SliceData(b), len(b))`.

## Solution

```go
import "unsafe"

// Str returns a string that shares b's bytes instead of copying them.
//
// The result is only valid while b is not written to again: a string is
// supposed to be immutable, and this one is not.
//
// Examples:
//
// 	Str([]byte("hi")) => "hi"
func Str(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(b), len(b))
}
```

## Walkthrough

For a 4096-byte buffer, `string(b)` allocates and copies 4096 bytes. `unsafe.String` writes a two-word header and copies nothing.

## Pitfalls

- Handing the result to a caller who outlives the buffer — that is a use-after-write, not a use-after-free, and it is silent.
- Skipping the empty-input guard; a nil slice's data pointer is nil, and a nil pointer with a non-zero length is invalid.
