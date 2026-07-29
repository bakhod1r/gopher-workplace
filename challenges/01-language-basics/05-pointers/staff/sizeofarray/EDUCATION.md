# Sizeof over arrays

## Intuition

An array's size is length times element size and is a compile-time constant; measuring an element gives only one slot.

## Approach

1. `Sizeof(p[0])` is one element; the whole array is `Sizeof(*p)`.
2. Dereference the pointer to measure the array.

## Solution

```go
import "unsafe"

func TotalSize(p *[4]int32) uintptr {
	return unsafe.Sizeof(*p)
}
```

## Walkthrough

For `[4]int32`, `Sizeof(p[0])` is 4. `Sizeof(*p)` measures all four elements → 16.

## Pitfalls

- `Sizeof(*p)` is the array; `Sizeof(p[0])` is one element.
- Arrays carry their length in the type, so Sizeof knows the total.
