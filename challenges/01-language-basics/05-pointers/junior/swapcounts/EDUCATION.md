# Swapping fields across instances

## Intuition

Parallel assignment on two struct pointers exchanges a field without a temporary, mutating both callers' structs.

## Approach

1. Access each field through its pointer.
2. Parallel-assign `a.Count, b.Count = b.Count, a.Count`.

## Solution

```go
type Cart struct{ Count int }

func SwapCounts(a, b *Cart) {
	a.Count, b.Count = b.Count, a.Count
}
```

## Walkthrough

`SwapCounts(&a, &b)` with counts `1` and `2`: the right side reads `2, 1`, then both fields are stored at once.

## Pitfalls

- Swapping the pointers (`a, b = b, a`) wouldn't affect the caller.
- Field-level swap touches the real instances.
