# Min Of Slice

## Intuition

Seeding from `s[0]` matters even more here: seeding from the zero value would report `0` as the minimum of any all-positive slice.

## Approach

1. Return `zero, false` for an empty slice.
2. Seed `best` from `s[0]`.
3. Replace `best` whenever a smaller element appears.

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

`MinOf([]int{4, 1, 3})` seeds `4`, drops to `1`, and keeps it.

## Pitfalls

- Seeding with `var best T`, which returns `0` for `[]int{4, 1, 3}`... no — worse, it returns `0` for any positive slice.
- Copying `MaxOf` without flipping the comparison.
- Returning `true` for an empty slice.
