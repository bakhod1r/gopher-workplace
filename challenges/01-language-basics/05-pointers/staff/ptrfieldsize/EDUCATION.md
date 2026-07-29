# Sizing pointer-holding structs

## Intuition

Each pointer field is word-sized; the struct's size is the sum (plus any padding). Measure the struct type, not a single pointer.

## Approach

1. The bug measures a single pointer (8), not the struct.
2. Measure the struct: `unsafe.Sizeof(Pair{})`.

## Solution

```go
import "unsafe"

type Pair struct {
	A *int
	B *int
}

func Size() uintptr {
	return unsafe.Sizeof(Pair{})
}
```

## Walkthrough

`Sizeof((*int)(nil))` is one pointer word. The `Pair` holds two pointer fields, so its size is 16.

## Pitfalls

- One pointer is 8 bytes; two make 16.
- `Sizeof(Pair{})` measures the whole struct.
