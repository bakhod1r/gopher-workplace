# unsafe.Slice length semantics

## Intuition

`unsafe.Slice(ptr, n)` builds a slice of n ELEMENTS; passing a byte length creates an out-of-bounds view.

## Approach

1. `unsafe.Slice(ptr, n)` takes an **element count**, not a byte size.
2. The bug passes `Sizeof(*p)`; pass `len(p)`.

## Solution

```go
import "unsafe"

func View(p *[4]int32) []int32 {
	return unsafe.Slice(&p[0], len(p))
}
```

## Walkthrough

`Sizeof(*p)` is the array's byte size (16), producing a wildly oversized slice. `len(p)` yields the correct 4-element view.

## Pitfalls

- `unsafe.Slice`'s length is in elements, not bytes.
- An over-long length yields a slice that reads past the array.
