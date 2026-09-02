# Self-Referential Constraint

## Intuition

The recursive constraint is the standard Go idiom for user-defined ordering, and it is checked at compile time: a type whose `Less` takes something else simply will not instantiate.

## Approach

1. Return zero and `false` for an empty slice.
2. Seed `best` from `s[0]`.
3. Replace `best` whenever `best.Less(v)`.

## Solution

```go
func MaxOf[T Lesser[T]](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	best := s[0]
	for _, v := range s[1:] {
		if best.Less(v) {
			best = v
		}
	}
	return best, true
}
```

## Walkthrough

`MaxOf([]Version{{2}, {2}})` finds `Less` false for equal versions, so the earlier element stays.

## Pitfalls

- Writing `[T Lesser[any]]`, which loses the type link.
- Inverting the comparison and computing the minimum.
- Using `!v.Less(best)`, which lets later ties win.
