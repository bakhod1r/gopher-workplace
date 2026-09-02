# Partition

## Intuition

Calling a filter twice would evaluate `pred` on every element twice. A single pass with two accumulators is both faster and guarantees the two halves stay consistent.

## Approach

1. Allocate `yes` and `no`, each with capacity `len(s)`.
2. Append each element to `yes` or `no` depending on `pred`.
3. Return both slices.

## Solution

```go
func Partition[T any](s []T, pred func(T) bool) ([]T, []T) {
	yes := make([]T, 0, len(s))
	no := make([]T, 0, len(s))
	for _, e := range s {
		if pred(e) {
			yes = append(yes, e)
		} else {
			no = append(no, e)
		}
	}
	return yes, no
}
```

## Walkthrough

`Partition([]int{1, 2, 3}, isEven)` sends 1 to `no`, 2 to `yes`, 3 to `no`, yielding `[2]` and `[1 3]`.

## Pitfalls

- Returning the rejected half first — the signature fixes the order.
- Calling `pred` twice per element by filtering the slice twice.
- Returning `nil` halves when nothing lands in one of them, if the tests expect empty slices.
