# Struct layout and padding

## Intuition

The compiler pads fields to satisfy alignment and rounds the struct to its largest alignment; ordering fields widest-first minimises wasted padding.

## Approach

1. Struct fields are laid out in declaration order with alignment padding.
2. `bool, int64, bool` forces 7 bytes of padding twice.
3. Reorder to `int64, bool, bool` so the two bools share the tail.

## Solution

```go
type Record struct {
	B int64
	A bool
	C bool
}
```

## Walkthrough

With `A bool, B int64, C bool`, the int64 must start at offset 8, wasting padding after A, then C pads to 24. Putting `B int64` first packs A and C into one 8-byte tail → 16.

## Pitfalls

- `bool,int64,bool` -> 24 bytes; `int64,bool,bool` -> 16.
- Alignment is platform-dependent; this targets 64-bit.
