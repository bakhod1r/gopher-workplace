# The Minimum That NaN Poisons

## Intuition

`cmp.Ordered` includes the float types, and floats carry a value that is neither less than, greater than, nor equal to anything — itself included. Seeding the running extremes from such a value freezes them: no later comparison can ever displace it.

## Approach

1. Skip any element that is not equal to itself — that is the portable NaN test.
2. Seed both extremes from the first usable element.
3. Widen the range with the ordinary comparisons and report whether anything was seen.

## Solution

```go
func MinMaxSkipNaN[T cmp.Ordered](xs []T) (T, T, bool) {
	var mn, mx T
	seen := false
	for _, v := range xs {
		if v != v {
			continue
		}
		if !seen {
			mn, mx, seen = v, v, true
			continue
		}
		if v < mn {
			mn = v
		}
		if v > mx {
			mx = v
		}
	}
	return mn, mx, seen
}
```

## Walkthrough

With `[NaN 3 1 2]` the first iteration sets `mn = mx = NaN`. Every subsequent `v < mn` and `v > mx` is false, so the function answers `NaN, NaN, true`.

## Pitfalls

- Testing for NaN with `math.IsNaN`, which does not compile for a `cmp.Ordered` type parameter.
- Assuming the guard costs integer instantiations something — `v != v` is false for every non-float.
