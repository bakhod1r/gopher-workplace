# Slices as stacks

## Intuition

A slice makes a natural stack: push with `append`, pop by reading the last
element and re-slicing to drop it:

```go
top := s[len(s)-1]
return s[:len(s)-1], top, true
```

## Approach

1. Bug: return s, top, true returns the popped value but the same full-length slice, so the stack never shrinks. 2. Fix: return the truncated slice s[:len(s)-1]. 3. That drops the last element while returning its value.

## Solution

```go
func Pop(s []int) ([]int, int, bool) {
	if len(s) == 0 {
		return s, 0, false
	}
	top := s[len(s)-1]
	return s[:len(s)-1], top, true
}
```

## Walkthrough

s=[1,2,3], top=3. Buggy return keeps s=[1,2,3]. Correct s[:2]=[1,2] plus top=3 and true.

## Pitfalls

- `s[:len(s)-1]` keeps the popped element in the backing array; for pointer
  elements, zero it to avoid a memory leak.
- Guard the empty case before indexing `len(s)-1`.
- The caller must use the returned slice (`s, v, ok = Pop(s)`).
