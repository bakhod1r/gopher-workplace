# String header vs data pointer

## Intuition

A string is a (data, len) header; `&s` addresses that header, while `unsafe.StringData(s)` addresses the bytes.

## Approach

1. `&s` is the address of the string **header**, not its bytes.
2. Use `unsafe.StringData(s)` to reach the first byte.

## Solution

```go
import "unsafe"

func FirstByte(s string) byte {
	p := unsafe.Pointer(unsafe.StringData(s))
	return *(*byte)(p)
}
```

## Walkthrough

Dereferencing `&s` reads header words, not text. `unsafe.StringData(s)` points at the actual byte data, so the first byte is `'X'`.

## Pitfalls

- `&s` reads the header fields, not the characters.
- `unsafe.StringData(s)` is the data pointer (bytes are immutable).
