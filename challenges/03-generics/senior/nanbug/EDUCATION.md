# NaN Poisons The Minimum

## Intuition

Skipping NaN inside the loop is correct, but seeding from `s[0]` can put NaN into `best`, after which every comparison fails and the NaN is returned.

## Approach

1. Track a `found` flag instead of seeding from an element.
2. Skip values unequal to themselves.
3. Take the first real value, then the smaller of each pair.

## Solution

```go
func MinIgnoringNaN[T Float](s []T) (T, bool) {
	var best T
	found := false
	for _, v := range s {
		if v != v {
			continue
		}
		if !found || v < best {
			best, found = v, true
		}
	}
	if !found {
		var zero T
		return zero, false
	}
	return best, true
}
```

## Walkthrough

`MinIgnoringNaN([NaN, 2])` seeds `best = NaN`; `2 < NaN` is false, so NaN survives to the return.

## Pitfalls

- Assuming a guard inside the loop protects the seed.
- Reaching for `math.IsNaN` when `v != v` is available and constraint-friendly.
- Reporting `true` for an all-NaN slice.
