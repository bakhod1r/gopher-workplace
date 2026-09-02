# Set Operations

## Intuition

Both operations build a fresh map, so neither input is disturbed — important because callers usually still need their original sets afterwards.

## Approach

1. `Union`: copy every key from `a`, then every key from `b`.
2. `Intersect`: for each key of `a`, keep it when `b` has it too.

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

func Intersect[T comparable](a, b map[T]struct{}) map[T]struct{} {
	out := make(map[T]struct{})
	for k := range a {
		if _, ok := b[k]; ok {
			out[k] = struct{}{}
		}
	}
	return out
}
```

## Walkthrough

`Intersect({1,2}, {2,3})` probes `b` for `1` (absent) and `2` (present), so only `2` survives.

## Pitfalls

- Writing results into `a`, which mutates the caller's set.
- Returning a nil map, which panics if the caller adds to it.
- Using a nested loop over both sets instead of a map probe.
