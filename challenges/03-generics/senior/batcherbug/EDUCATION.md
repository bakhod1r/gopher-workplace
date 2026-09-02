# The Batch That Never Ships

## Intuition

The accumulator is only appended when it fills up, so anything left over when the input runs out is dropped on the floor.

## Approach

1. Accumulate into `cur`.
2. Emit and reset when `cur` is full.
3. After the loop, emit `cur` if it is non-empty.

## Solution

```go
func Batches[T any](s []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}
	out := make([][]T, 0)
	cur := make([]T, 0, size)
	for _, v := range s {
		cur = append(cur, v)
		if len(cur) == size {
			out = append(out, cur)
			cur = make([]T, 0, size)
		}
	}
	if len(cur) > 0 {
		out = append(out, cur)
	}
	return out
}
```

## Walkthrough

`Batches([]int{1,2,3}, 2)` emits `{1,2}` and then leaves `{3}` in `cur` unemitted.

## Pitfalls

- Testing only with inputs that divide evenly.
- Emitting an empty final batch when the input does divide evenly.
