# The Clone That Shares Its Slices

## Intuition

`maps.Clone` is a shallow copy: it builds a new hash table and copies each value *as a value*. For a slice value that means copying the header — pointer, length, capacity — so both maps point at one backing array. Writing `copy[k][0]` writes through to `m[k][0]`.

## Approach

1. Return nil for a nil map so the shape is preserved.
2. Allocate a map of the same size.
3. Clone each slice value on the way in.

## Solution

```go
func CloneTags(m map[string][]string) map[string][]string {
	if m == nil {
		return nil
	}
	out := make(map[string][]string, len(m))
	for k, v := range m {
		out[k] = slices.Clone(v)
	}
	return out
}
```

## Walkthrough

`c := CloneTags(m); c["a"][0] = "z"` also sets `m["a"][0]` to `"z"`, because both headers address the same array.

## Pitfalls

- Believing `maps.Clone` is deep because the word "clone" suggests it; the doc says shallow.
- Appending to a cloned slice instead of writing to it — that usually *looks* fine, which makes the bug harder to find.
