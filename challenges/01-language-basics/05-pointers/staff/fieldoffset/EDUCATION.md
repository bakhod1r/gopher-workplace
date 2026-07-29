# Locating fields with unsafe.Offsetof

## Intuition

`Offsetof` returns a field's byte position within its struct; `Sizeof` returns a type's width — different quantities.

## Approach

1. To reach field B you need its **offset**, not the size of B.
2. Use `unsafe.Offsetof(p.B)`.

## Solution

```go
import "unsafe"

type Pair struct {
	A int64
	B int32
}

func SecondField(p *Pair) int32 {
	base := unsafe.Pointer(p)
	off := unsafe.Offsetof(p.B)
	return *(*int32)(unsafe.Add(base, off))
}
```

## Walkthrough

`Sizeof(p.B)` happens to equal the offset only by luck of layout; `Offsetof(p.B)` is the correct byte distance from the struct start to B.

## Pitfalls

- `Sizeof(p.B)` is 4 (width), `Offsetof(p.B)` is 4 (position) — they coincide here by luck of layout but mean different things; with a wider first field they diverge.
- Always use Offsetof for a field's position.
