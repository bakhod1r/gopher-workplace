# Deep-copying pointer fields

## Intuition

Struct copy duplicates a pointer field's address, sharing the pointee; a deep clone allocates and copies what the pointer references.

## Approach

1. `cp := *b` copies the struct but shares the `P` pointer.
2. Deep-copy the pointee: `v := *b.P; cp.P = &v`.

## Solution

```go
type Box struct {
	P *int
}

func Clone(b *Box) *Box {
	cp := *b
	v := *b.P
	cp.P = &v
	return &cp
}
```

## Walkthrough

The shallow copy shares `P`, so writing `*c.P` also changes `*b.P`. Allocating a fresh pointee for the clone breaks the sharing.

## Pitfalls

- `*b` copies the pointer value, not the pointee.
- Allocate a new pointee and copy the value into it.
