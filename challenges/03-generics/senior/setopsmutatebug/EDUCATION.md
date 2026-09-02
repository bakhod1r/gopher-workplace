# Union That Eats Its Input

## Intuition

A map value is a pointer to the same hash table, so writing to `out` writes to `a`. The caller's set is silently mutated.

## Approach

1. Allocate a fresh map.
2. Copy `a` into it.
3. Copy `b` into it.

## Solution

```go
func Union[T comparable](a, b map[T]struct{}) map[T]struct{} {
	out := make(map[T]struct{}, len(a)+len(b))
	for k := range a {
		out[k] = struct{}{}
	}
	for k := range b {
		out[k] = struct{}{}
	}
	return out
}
```

## Walkthrough

`Union(role, grant)` leaves `role` containing the grants, so the next check starts from the polluted set.

## Pitfalls

- The same mistake with slices, where `append` sometimes hides it.
- Returning `b` instead — equally wrong, just a different victim.
