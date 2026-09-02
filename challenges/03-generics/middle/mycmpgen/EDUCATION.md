# Writing Your Own Ordered

## Intuition

Writing the set by hand shows why it is worth importing: miss one integer width and a caller's type silently fails to instantiate.

## Approach

1. Return zero and `false` for an empty slice.
2. Seed from `s[0]`.
3. Replace on a strictly larger element.

## Solution

```go
func Largest[T Ordered](s []T) (T, bool) {
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

`Largest([]string{"a", "c"})` works because `~string` is in the hand-written set, exactly as in `cmp.Ordered`.

## Pitfalls

- Forgetting a width such as `~int32`, which rejects valid callers.
- Omitting `~` and rejecting every named type.
- Shipping a private copy of `Ordered` when `cmp.Ordered` exists.
