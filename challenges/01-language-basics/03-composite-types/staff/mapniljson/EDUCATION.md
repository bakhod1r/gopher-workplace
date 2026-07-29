# Nil and empty maps at the boundary

## Intuition

Like slices, a nil map and an empty map both have length 0 but serialize
differently: `encoding/json` emits `null` for nil and `{}` for an allocated empty
map:

```go
m := make(map[string]int) // marshals to {}, even when empty
```

## Approach

1. Bug: `var m map[string]int` starts nil; with empty input the loop never initializes it, so a nil map is returned and JSON-encodes to null.
2. Fix: `m := map[string]int{}` starts non-nil so empty input encodes to {}.

## Solution

```go
func Counts(xs []string) map[string]int {
	m := map[string]int{}
	for _, x := range xs {
		if m == nil {
			m = map[string]int{}
		}
		m[x]++
	}
	return m
}
```

## Walkthrough

xs=[]: fix returns an initialized empty map -> json.Marshal gives {}. Buggy nil map -> null. Non-empty inputs behave the same either way.

## Pitfalls

- `m == nil` is the only way to tell nil from empty.
- Reading nil maps is safe; writing panics — allocate before the loop.
- Prefer `make` at the top when the emptiness must be observable.
