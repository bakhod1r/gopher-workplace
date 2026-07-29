# What unsafe.Sizeof measures

## Intuition

`Sizeof` returns the size of its operand's type; a pointer is always word-sized regardless of the pointee. Measure the element expression instead.

## Approach

1. `unsafe.Sizeof(p)` is the size of the **pointer** (8 on 64-bit), not the element.
2. Measure an element: `unsafe.Sizeof(p[0])`.

## Solution

```go
import "unsafe"

func ElemSize(p *[8]int32) uintptr {
	return unsafe.Sizeof(p[0])
}
```

## Walkthrough

The bug returns 8 for every pointer. `p[0]` is an `int32`, so its size is 4.

## Pitfalls

- `Sizeof(p)` is 8 (a pointer); `Sizeof(p[0])` is the element size.
- Sizeof is a compile-time constant of the type.
