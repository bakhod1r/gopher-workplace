# Slices as stacks

## The idea

A slice makes a natural stack: push with `append`, pop by reading the last
element and re-slicing to drop it:

```go
top := s[len(s)-1]
return s[:len(s)-1], top, true
```

## Why it matters

Forgetting to shrink means the "pop" doesn't advance — a stack that never empties.
Because the slice is returned by value, the caller must reassign.

## Watch out

- `s[:len(s)-1]` keeps the popped element in the backing array; for pointer
  elements, zero it to avoid a memory leak.
- Guard the empty case before indexing `len(s)-1`.
- The caller must use the returned slice (`s, v, ok = Pop(s)`).
