# append returns a new header

## Intuition

A slice header carries length and capacity by value. `append` may write into the
existing backing array, but it always returns an updated header — you must assign
it back:

```go
out = append(out, x*2)
```

## Approach

1. Bug: _ = append(out, x*2) discards append's return value, so out never grows. 2. append may return a new slice header (new pointer/len/cap) that must be reassigned. 3. Fix: out = append(out, x*2).

## Solution

```go
func Doubled(xs []int) []int {
	out := make([]int, 0, len(xs))
	for _, x := range xs {
		out = append(out, x*2)
	}
	return out
}
```

## Walkthrough

out starts len 0 cap 3. Each iteration append returns a header with len+1; without reassignment out.len stays 0, so final out is empty. Reassigning captures the growing length -> [2,4,6].

## Pitfalls

- Always `s = append(s, ...)`.
- `append` may reallocate; the old header is stale afterward.
- The compiler doesn't require using the result (it's just a value), so this
  compiles.
