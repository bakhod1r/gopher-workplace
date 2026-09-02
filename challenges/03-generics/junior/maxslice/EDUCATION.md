# Max Of Slice

## Intuition

There is no generic `-Inf` you could seed with, so the first element is the only safe starting point — which is exactly why the empty case needs the `bool`.

## Approach

1. Return `zero, false` when the slice is empty.
2. Seed `best` from `s[0]`.
3. Replace `best` whenever a larger element appears.

## Solution

```go
func MaxOf[T cmp.Ordered](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	best := s[0]
	for _, v := range s[1:] {
		if v > best {
			best = v
		}
	}
	return best, true
}
```

## Walkthrough

`MaxOf([]int{1, 9, 3})` seeds `best = 1`, updates to `9`, then leaves it alone at `3`.

## Pitfalls

- Seeding with the zero value — negative-only slices then return `0`.
- Ranging over all of `s` including index 0 (harmless, but the comparison is wasted).
- Panicking on an empty slice instead of reporting `false`.
