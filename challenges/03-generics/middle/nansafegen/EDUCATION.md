# Ordering And NaN

## Intuition

Because NaN loses every comparison, the usual `s[0]` seed can poison the whole scan — the `found` flag is what keeps the algorithm correct.

## Approach

1. Scan, skipping values that are not equal to themselves.
2. Record the first real value, then keep the smaller.
3. Report `false` when nothing was found.

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

`MinIgnoringNaN([]float64{3, NaN, 1})` skips the NaN, so `1` wins normally.

## Pitfalls

- Seeding `best` from `s[0]` when it may be NaN.
- Using `slices.Min`, which propagates NaN.
- Testing for NaN with `==` against a NaN constant, which is always false.
