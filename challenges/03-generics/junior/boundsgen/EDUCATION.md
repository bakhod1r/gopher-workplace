# Bounds

## Intuition

A one-element slice must report the same value twice, which falls out naturally from seeding both accumulators from `s[0]`.

## Approach

1. Return zero values and `false` when the slice is empty.
2. Seed `lo` and `hi` from `s[0]`.
3. Update each accumulator independently while scanning the rest.

## Solution

```go
func Bounds[T cmp.Ordered](s []T) (T, T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, zero, false
	}
	lo, hi := s[0], s[0]
	for _, v := range s[1:] {
		if v < lo {
			lo = v
		}
		if v > hi {
			hi = v
		}
	}
	return lo, hi, true
}
```

## Walkthrough

`Bounds([]int{5})` seeds `lo = hi = 5`, skips the loop, and returns `5, 5, true`.

## Pitfalls

- Using `else if` between the two updates, which is fine here but breaks if you seed differently.
- Seeding from zero values, which is wrong for all-positive or all-negative data.
- Scanning twice by calling `MinOf` and `MaxOf` separately.
