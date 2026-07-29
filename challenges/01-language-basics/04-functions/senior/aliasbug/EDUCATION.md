# Slice assignment copies only the header

## Intuition

Assigning a slice duplicates its (ptr,len,cap) header but not the backing array, so writes through either alias are shared.

## Approach

1. `cp := xs` copies the header but shares the backing array.
2. Copy the data: `cp := append([]int(nil), xs...)` before mutating.

## Solution

```go
func WithFirst(xs []int, v int) []int {
	cp := append([]int(nil), xs...)
	cp[0] = v
	return cp
}
```

## Walkthrough

Writing `cp[0]` through the shared array also changes the caller's slice. Copying into a fresh array isolates the mutation.

## Pitfalls

- `cp := xs` is NOT a copy of the elements.
- Use `append([]T(nil), xs...)` or `make`+`copy` for independence.
