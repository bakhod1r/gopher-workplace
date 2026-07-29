# Slice header vs backing data

## Intuition

A slice value is a header; `&s` addresses that header, while `&s[0]` (or `unsafe.SliceData`) addresses the actual data.

## Approach

1. `&s` is the address of the slice **header**, not the data.
2. Point at the first element: `unsafe.Pointer(&s[0])`.

## Solution

```go
import "unsafe"

func SetFirst(s []int, v int) {
	p := unsafe.Pointer(&s[0])
	*(*int)(p) = v
}
```

## Walkthrough

Writing through `&s` corrupts the header. `&s[0]` addresses the backing array, so the store updates `s[0]`.

## Pitfalls

- `&s` writes over the slice header fields, not the elements.
- `&s[0]` is the data pointer.
