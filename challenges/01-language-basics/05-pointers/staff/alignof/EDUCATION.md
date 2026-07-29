# Alignment vs size

## Intuition

`Alignof` reports the address boundary a type must sit on; `Sizeof` reports its width. They match for basic types but the concepts (and APIs) are distinct.

## Approach

1. Alignment and size are different properties.
2. The bug returns `Sizeof(s.B)`; use `unsafe.Alignof(s.B)`.

## Solution

```go
import "unsafe"

type S struct {
	A bool
	B [3]int32
}

func FieldAlign(s *S) uintptr {
	return unsafe.Alignof(s.B)
}
```

## Walkthrough

For an int64 both happen to be 8, but the question asks for alignment. `Alignof` is the correct operator and stays right for types where size != alignment.

## Pitfalls

- Use `Alignof` for alignment, `Sizeof` for width.
- Struct padding is driven by field alignment.
