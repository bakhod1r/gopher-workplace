# Reach A Field Through Its Offset

## Intuition

`unsafe.Add` is pointer arithmetic that stays legal: the result is still a pointer into the same object, so the garbage collector keeps tracking it. Doing the same arithmetic on a `uintptr` would not be safe.

## Approach

1. Convert `p` to `unsafe.Pointer`.
2. `unsafe.Add` it by `unsafe.Offsetof(p.Seq)`.
3. Convert to `*int64`, increment through it, and return the value.

## Solution

```go
import "unsafe"

// Rec is a record whose Seq field is reached by offset.
type Rec struct {
	Tag  byte
	Seq  int64
	Name string
}

// BumpSeq increments the Seq field of the record p points at, using the
// field's offset rather than the field selector, and returns the new value.
//
// This is what a generic marshaller does when it only knows the offset.
//
// Examples:
//
// 	r := &Rec{Seq: 1}; BumpSeq(r) => 2, r.Seq is 2
func BumpSeq(p *Rec) int64 {
	q := (*int64)(unsafe.Add(unsafe.Pointer(p), unsafe.Offsetof(p.Seq)))
	*q++
	return *q
}
```

## Walkthrough

`Tag` is at 0, so `Seq` is at offset 8 after padding. Adding 8 to the struct's address and reading it as `*int64` is exactly `&p.Seq`.

## Pitfalls

- Doing the arithmetic in `uintptr` across separate statements — a moving collector would invalidate the number.
- Hard-coding the offset instead of asking `Offsetof`, which breaks the moment the struct changes.
