# Full-slice expressions and capacity

## Intuition

`s[low:high:max]` sets capacity to `max-low`; using it to cap a sub-slice prevents appends from reaching the parent and lets the unused tail be garbage-collected.

## Approach

1. `xs[:k]` retains the original capacity, so appending overwrites `xs[k]`.
2. Use a full-slice bound `xs[:k:k]`.

## Solution

```go
func Head(xs []int, k int) []int {
	return xs[:k:k]
}
```

## Walkthrough

`xs[:2]` keeps capacity 5, so `append(h, 99)` writes into `xs[2]`. `xs[:2:2]` caps capacity, forcing a fresh array on append.

## Pitfalls

- `xs[:k]` keeps cap == cap(xs); appends spill into the parent.
- `xs[:k:k]` (or a copy) isolates the head.
