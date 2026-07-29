# Pointer arithmetic in bytes

## Intuition

`unsafe.Add` works in bytes; indexing into a typed array requires multiplying the index by the element size.

## Approach

1. `unsafe.Add` advances by **bytes**, not elements.
2. The bug adds `i` bytes; multiply by the element size.
3. `unsafe.Add(base, uintptr(i)*unsafe.Sizeof(arr[0]))`.

## Solution

```go
import "unsafe"

func At(arr *[4]int32, i int) int32 {
	base := unsafe.Pointer(arr)
	p := unsafe.Add(base, uintptr(i)*unsafe.Sizeof(arr[0]))
	return *(*int32)(p)
}
```

## Walkthrough

For `i = 2` on `int`s, the byte offset must be `2 * 8 = 16`. Adding just `2` bytes lands mid-element and reads garbage.

## Pitfalls

- `unsafe.Add(base, i)` moves i bytes, not i int32s.
- Multiply by `unsafe.Sizeof(elem)` for element steps.
