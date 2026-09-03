# Let The Standard Library Size It

## Intuition

Joining is a two-pass algorithm: measure, then fill. `strings.Join` already does both, so writing the loop yourself only adds allocations.

## Approach

1. `JoinPath` delegates to `strings.Join`.
2. `SplitPath` special-cases the empty string, then delegates to `strings.Split`.

## Solution

```go
func JoinPath(parts []string) string {
	return strings.Join(parts, "/")
}

func SplitPath(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, "/")
}
```

## Walkthrough

`strings.Split("", "/")` returns `[""]` — a single empty segment — which would make `SplitPath` the inverse of nothing, so the guard is what preserves the round trip.

## Pitfalls

- Building the result with `+=` and trimming the trailing separator.
- Forgetting the empty-string case and returning a one-element slice.
- Reaching for `path.Join`, which also cleans the path and drops empty segments.
