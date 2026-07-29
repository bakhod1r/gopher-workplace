# Padding and field offsets

## Intuition

Fields are placed at aligned offsets, so a field's position is not the sum of prior field sizes; `Offsetof` reports the true, padded offset.

## Approach

1. `Sizeof(r.Flag)` ignores the padding before `N`.
2. Use `unsafe.Offsetof(r.N)` to get N's true offset.

## Solution

```go
import "unsafe"

type Rec struct {
	Flag bool
	N    int64
}

func ReadN(r *Rec) int64 {
	base := unsafe.Pointer(r)
	off := unsafe.Offsetof(r.N)
	return *(*int64)(unsafe.Add(base, off))
}
```

## Walkthrough

A `bool` then an `int64` inserts 7 padding bytes; the size of the bool (1) is not the offset of N (8). `Offsetof` accounts for the padding.

## Pitfalls

- `N` follows 1 bool + 7 padding bytes -> offset 8.
- `Offsetof(r.N)` gives the correct position.
