# Compare Bytes To A String Without Converting

## Intuition

Zero-copy is a question of lifetime. A string view built purely to feed a comparison never escapes the function, so nothing can ever see that it aliased a mutable slice.

## Approach

1. Return false on a length mismatch, true for two empties.
2. Build a string view over `b` and compare it with `s`.

## Solution

```go
import "unsafe"

// EqualString reports whether b's bytes are exactly s.
//
// Neither side may be converted: a conversion in either direction copies
// the payload just to throw the copy away.
//
// Examples:
//
// 	EqualString([]byte("hi"), "hi") => true
func EqualString(b []byte, s string) bool {
	if len(b) != len(s) {
		return false
	}
	if len(b) == 0 {
		return true
	}
	return unsafe.String(unsafe.SliceData(b), len(b)) == s
}
```

## Walkthrough

For a 3584-byte payload, `string(b) == s` allocates and copies 3584 bytes. The view compares the same bytes in place and allocates nothing.

## Pitfalls

- Returning or storing the view — the safety argument depends on it dying here.
- Skipping the empty guard; a nil data pointer with a non-zero length is invalid.
