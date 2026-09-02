# Minimum By Less

## Intuition

With a method-based ordering there is no `>`, so the only way to express "smaller" is to call `Less` in the other direction.

## Approach

1. Return zero and `false` for an empty slice.
2. Seed `best` from `s[0]`.
3. Replace `best` whenever `v.Less(best)`.

## Solution

```go
func MinOf[T Lesser[T]](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	best := s[0]
	for _, v := range s[1:] {
		if v.Less(best) {
			best = v
		}
	}
	return best, true
}
```

## Walkthrough

`MinOf([]Version{{3}, {1}})` finds `{1}.Less({3})` true and replaces the incumbent.

## Pitfalls

- Calling `best.Less(v)`, which computes the maximum.
- Negating `Less` to mean "greater or equal", which mishandles ties.
- Seeding from a zero value instead of `s[0]`.
