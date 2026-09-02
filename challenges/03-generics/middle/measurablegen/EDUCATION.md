# Constraint With A Method

## Intuition

Compared with taking `[]Measurable`, the type parameter keeps the caller's slice untouched and returns the concrete type, so no assertion is needed afterwards.

## Approach

1. Return zero and `false` for an empty slice.
2. Seed from `s[0]` and cache its value.
3. Replace on a strictly larger value.

## Solution

```go
func Heaviest[T Measurable](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	best := s[0]
	bestV := best.Value()
	for _, v := range s[1:] {
		x := v.Value()
		if x > bestV {
			best, bestV = v, x
		}
	}
	return best, true
}
```

## Walkthrough

`Heaviest([]Reading{{2}, {2}})` refuses the tie because the comparison is strict, keeping the first element.

## Pitfalls

- Taking `[]Measurable` and forcing callers to convert their slice.
- Calling `v.Value()` twice per iteration.
- Returning `Measurable` instead of `T`, losing the concrete type.
