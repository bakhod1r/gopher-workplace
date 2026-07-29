# Data pointers of empty slices

## Intuition

`unsafe.SliceData` of an empty slice can be nil; you must check the length before dereferencing it, just like any pointer.

## Approach

1. `unsafe.SliceData` of an empty/nil slice yields a nil pointer; dereferencing panics.
2. Guard `len(s) == 0` before reading.

## Solution

```go
import "unsafe"

func FirstOr(s []int, def int) int {
	if len(s) == 0 {
		return def
	}
	return *unsafe.SliceData(s)
}
```

## Walkthrough

`FirstOr(nil, 7)` should return the default, but `*SliceData(nil)` dereferences nil. The length guard returns `def` safely.

## Pitfalls

- Empty slices may yield a nil data pointer.
- Guard `len(s) == 0` before reading element 0.
