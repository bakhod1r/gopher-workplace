# String-Like Map Keys

## Intuition

The output key type is deliberately concrete: once normalised, the named type no longer carries meaning, and callers want a plain map.

## Approach

1. Allocate the result with capacity `len(m)`.
2. Convert and lowercase each key while copying the value.

## Solution

```go
func Normalize[K ~string, V any](m map[K]V) map[string]V {
	out := make(map[string]V, len(m))
	for k, v := range m {
		out[strings.ToLower(string(k))] = v
	}
	return out
}
```

## Walkthrough

Two headers `"A"` and `"a"` both normalise to `"a"`, so the result holds one entry.

## Pitfalls

- Using `string` instead of `~string` and rejecting the named type.
- Returning `map[K]V` and leaving the keys un-normalised.
- Returning a nil map for a nil input.
