# Concat That Returns Its Input

## Intuition

The one-slice fast path hands the caller's own slice back, so writing to the result writes to the input.

## Approach

1. Sum the lengths.
2. Allocate once with that capacity.
3. Append every slice — including when there is only one.

## Solution

```go
func Concat[T any](ss ...[]T) []T {
	n := 0
	for _, s := range ss {
		n += len(s)
	}
	out := make([]T, 0, n)
	for _, s := range ss {
		out = append(out, s...)
	}
	return out
}
```

## Walkthrough

`out := Concat(src); out[0] = 99` also sets `src[0]` to 99.

## Pitfalls

- Keeping the fast path but returning a clone — fine, but the general path already does that.
- Assuming the empty case must return `nil`; an empty non-nil slice is the friendlier result.
