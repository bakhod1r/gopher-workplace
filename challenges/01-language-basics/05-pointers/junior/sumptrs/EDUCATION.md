# Reading pointer collections safely

## Intuition

Dereferencing each non-nil pointer and summing avoids the nil-deref panic on empty slots.

## Approach

1. Range the slice.
2. Add `*p` only when `p != nil`.
3. Return the running total.

## Solution

```go
func SumPtrs(ps []*int) int {
	total := 0
	for _, p := range ps {
		if p != nil {
			total += *p
		}
	}
	return total
}
```

## Walkthrough

`[]*int{&a, nil, &b}` with `a = 1`, `b = 2`: add `1`, skip nil, add `2` → `3`.

## Pitfalls

- Dereferencing a nil element panics.
- Skip nils, sum the rest.
