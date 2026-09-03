# Minimum Seeded Wrong

## Intuition

Seeding with `var best T` silently inserts a zero into the comparison, so any slice whose values are all above zero reports the zero instead.

## Approach

1. Return early for an empty slice.
2. Seed `best` from `s[0]`.
3. Compare against the remaining elements.

## Solution

```go
func MinOf[T cmp.Ordered](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	best := s[0]
	for _, v := range s[1:] {
		if v < best {
			best = v
		}
	}
	return best, true
}
```

## Walkthrough

`MinOf([]int{4, 7})` starts from `0`, never finds anything smaller, and returns `0`.

## Pitfalls

- Seeding from a zero value or a hard-coded sentinel.
- Ranging over all of `s` after seeding from `s[0]` — harmless, but a wasted comparison.
- Assuming the bug only affects positive data: negatives break the maximum the same way.
