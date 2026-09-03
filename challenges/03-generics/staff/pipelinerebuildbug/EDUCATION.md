# Five Stages, Five Copies

## Intuition

The first `make` plus `copy` already gives the function a private buffer. Every stage after that can overwrite it in place; building a fresh slice per stage multiplies both the allocation count and the bytes copied by the number of stages.

## Approach

1. Copy the input into a buffer you own.
2. For each stage, overwrite every element of that buffer in place.
3. Return the buffer.

## Solution

```go
func Pipeline[T any](s []T, stages ...func(T) T) []T {
	out := make([]T, len(s))
	copy(out, s)
	for _, st := range stages {
		for i := range out {
			out[i] = st(out[i])
		}
	}
	return out
}
```

## Walkthrough

With five stages over two million elements the buggy version allocates six buffers and copies twelve million elements; the fixed version allocates one and copies two million.

## Pitfalls

- Skipping the initial copy and transforming `s` in place — cheaper still, but it corrupts the caller's slice.
- Assuming the allocator makes repeated large `make` calls free; they are the dominant cost here.
