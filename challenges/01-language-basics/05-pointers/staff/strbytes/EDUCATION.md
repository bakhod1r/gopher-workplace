# String header fields

## Intuition

A string is a two-word header: a data pointer and an int length. Reading the wrong word returns an address instead of the length.

## Approach

1. The byte length is the header's `Len` field.
2. The bug returns the data pointer as an int; return `h.Len`.

## Solution

```go
import "unsafe"

type strHeader struct {
	Data unsafe.Pointer
	Len  int
}

func ByteLen(s string) int {
	h := (*strHeader)(unsafe.Pointer(&s))
	return h.Len
}
```

## Walkthrough

`uintptr(h.Data)` is an address, not a length. `h.Len` is the string's byte count, so `"héllo"` reports 6.

## Pitfalls

- The first word is the data pointer; the second is the length.
- `len(s)` is the safe, idiomatic way; this exercises the header layout.
