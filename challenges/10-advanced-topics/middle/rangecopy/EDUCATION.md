# The Loop Variable Is A Copy

## Intuition

`for _, c := range s` binds `c` to a copy. Writing to it updates the copy and the copy is discarded — which is why the loop is silently a no-op rather than a compile error.

## Approach

1. Range over the indices only.
2. Write through `items[i]`.

## Solution

```go
// Counter is one element of the slice.
type Counter struct {
	N   int
	Pad [64]byte
}

// Bump increments every counter in items, in place.
//
// Ranging by value copies each element; the increment has to reach the
// slice's own storage.
//
// Examples:
//
// 	items := []Counter{{N: 1}}; Bump(items) => items[0].N == 2
func Bump(items []Counter) {
	for i := range items {
		items[i].N++
	}
}
```

## Walkthrough

The buggy loop copies 72 bytes per element, increments the copy, and drops it. Indexing writes into the caller's array directly and copies nothing.

## Pitfalls

- `for i, c := range items { items[i] = c }` after mutating `c` — correct, and it copies twice.
- Assuming a slice of pointers behaves the same; there the copy is a pointer you can write through.
