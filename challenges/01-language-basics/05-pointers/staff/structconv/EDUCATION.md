# Reinterpreting compatible structs

## Intuition

Structs with identical memory layout can be converted with a single unsafe reinterpret, copying all fields at once.

## Approach

1. Reinterpreting `*Point` as `*Vec` gives both fields at once.
2. The bug rebuilds a `Vec` keeping only `X`; return the whole reinterpreted value.

## Solution

```go
import "unsafe"

type Point struct{ X, Y int32 }
type Vec struct{ X, Y int32 }

func ToVec(p *Point) Vec {
	return *(*Vec)(unsafe.Pointer(p))
}
```

## Walkthrough

`Vec{X: v.X}` drops the Y field, yielding `{3 0}`. Returning `*(*Vec)(unsafe.Pointer(p))` preserves both coordinates.

## Pitfalls

- `*(*Vec)(unsafe.Pointer(p))` copies X and Y together.
- Rebuilding by hand risks dropping fields.
