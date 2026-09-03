# Where Does This Field Start

## Intuition

Fields are laid out in declaration order, but each one starts at an offset its type can be aligned to. That is why a byte followed by an int64 leaves seven bytes unused.

## Approach

1. Declare a zero `Rec`.
2. Return `unsafe.Offsetof` for each of its three fields.

## Solution

```go
import "unsafe"

// Rec is a record with mixed field widths.
type Rec struct {
	A byte
	B int64
	C byte
}

// Offsets returns the byte offset of each field of Rec from the start of
// the struct.
//
// Offsets are decided by the compiler from the field order and the
// alignment rules, not by the field sizes alone.
//
// Examples:
//
// 	Offsets() => 0, 8, 16 for the declared layout
func Offsets() (a, b, c uintptr) {
	var r Rec
	return unsafe.Offsetof(r.A), unsafe.Offsetof(r.B), unsafe.Offsetof(r.C)
}
```

## Walkthrough

`A` is at 0. `B` is an int64 and needs 8-byte alignment, so it starts at 8, not 1. `C` follows at 16, and the struct is padded to 24.

## Pitfalls

- Computing offsets as running sums of `Sizeof` — that ignores alignment.
- `Offsetof` needs the field selector `r.B`, not the type.
