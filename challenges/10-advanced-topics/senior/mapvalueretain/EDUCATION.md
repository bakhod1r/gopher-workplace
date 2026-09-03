# The Map Entry That Pinned The Whole Buffer

## Intuition

`batch[off:off+n:off+n]` protects the batch from being appended over, and does nothing about retention — the header still points into the megabyte. The collector frees allocations, not the parts of them nobody uses.

## Approach

1. Validate the range as before.
2. Allocate `n` bytes, copy the slice into them, store the copy.

## Solution

```go
// Index stores the n bytes of batch at offset off under key.
//
// The map outlives the batch, so the stored value must own its bytes: a
// view keeps the entire batch reachable for as long as the entry lives.
//
// Examples:
//
// 	Index(m, "a", batch, 0, 4) => m["a"] is a 4-byte copy
func Index(m map[string][]byte, key string, batch []byte, off, n int) {
	if m == nil || off < 0 || n < 0 || off+n > len(batch) {
		return
	}
	owned := make([]byte, n)
	copy(owned, batch[off:off+n])
	m[key] = owned
}
```

## Walkthrough

Storing 8 bytes out of a 1 MiB batch keeps 1 MiB reachable with a view, and 8 bytes with a copy.

## Pitfalls

- Copying only when `n` is small; the entry's lifetime, not its size, is what matters.
- `append([]byte(nil), batch[off:off+n]...)` is also a copy — correct, just less explicit about the size.
