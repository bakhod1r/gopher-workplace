# Partition That Reverses One Half

## Intuition

One branch appends and the other prepends, so the failing half comes out reversed — and in quadratic time.

## Approach

1. Allocate both halves.
2. Append to whichever half the predicate selects.

## Solution

```go
func Partition[T any](s []T, pred func(T) bool) (yes, no []T) {
	yes, no = make([]T, 0), make([]T, 0)
	for _, v := range s {
		if pred(v) {
			yes = append(yes, v)
		} else {
			no = append(no, v)
		}
	}
	return yes, no
}
```

## Walkthrough

For `[1 2 3 4]` the odd half is built as `[1]`, then `[3 1]`.

## Pitfalls

- Assuming order does not matter because the caller sorts anyway.
- Returning `nil` halves, which prints as `[]` but is not equal to an empty slice under some checks.
