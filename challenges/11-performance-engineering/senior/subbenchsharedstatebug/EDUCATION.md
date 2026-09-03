# Sub-Benchmarks That Contaminate Each Other

## Intuition

A shared scratch buffer is an optimisation right up until someone forgets that "shared" also means "still full of the last run's data".

## Approach

1. Reset the buffer's length at the start of each run.

## Solution

```go
func (r *Runner) RunSize(size int) int {
	r.buf = r.buf[:0]
	for i := 0; i < size; i++ {
		r.buf = append(r.buf, 1)
	}
	sum := 0
	for _, v := range r.buf {
		sum += v
	}
	return sum
}
```

## Walkthrough

With the bug, `RunAll([1, 10, 100])` returns `[1, 11, 111]`: each run inherits everything the previous ones appended, so the "cost" of a size includes every size before it.

## Pitfalls

- `r.buf = nil`, which is correct but throws away the capacity the buffer existed to provide.
- Sharing any mutable state across `b.Run` cases — a map, a counter, a connection.
- Concluding a size sweep shows superlinear behaviour without checking the reverse order first.
