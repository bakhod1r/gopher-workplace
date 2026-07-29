# Fully resetting pooled objects

## Intuition

Object-pool reuse requires clearing every field; truncating a slice with `[:0]` keeps its capacity for reuse while dropping the length.

## Approach

1. Resetting must clear both the length counter and the data view.
2. The bug sets only `b.Len = 0`; also `b.Data = b.Data[:0]`.

## Solution

```go
type Buf struct {
	Data []byte
	Len  int
}

func Reset(b *Buf) {
	b.Len = 0
	b.Data = b.Data[:0]
}
```

## Walkthrough

Leaving the data slice full keeps stale bytes reachable. Reslicing `b.Data[:0]` alongside `Len = 0` fully resets the buffer.

## Pitfalls

- Zeroing one field leaves the rest stale.
- `b.Data[:0]` empties the slice but keeps its backing array for reuse.
