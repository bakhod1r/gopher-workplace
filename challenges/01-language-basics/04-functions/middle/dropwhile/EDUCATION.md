# Skipping a matching prefix

## Intuition

DropWhile locates the first non-matching element and returns the remainder; copying the tail avoids sharing the caller's backing array.

## Approach

1. Advance an index while the predicate holds.
2. Return a copy of the remaining tail.

## Solution

```go
func DropWhile(xs []int, pred func(int) bool) []int {
	i := 0
	for i < len(xs) && pred(xs[i]) {
		i++
	}
	return append([]int(nil), xs[i:]...)
}
```

## Walkthrough

`[2 4 5 6]` with even: skip 2 and 4, return `[5 6]`.

## Pitfalls

- Returning `xs[i:]` directly would alias the caller's array; copy it.
- If everything matches, the result is empty.
