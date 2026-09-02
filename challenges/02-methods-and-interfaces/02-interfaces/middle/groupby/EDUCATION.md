# Group By

## Intuition

The bucketing algorithm never changes; only the key extractor does. Passing it as an interface value keeps `Group` at one pass and zero knowledge of the domain.

## Approach

1. `ByFirstLetter.Key` guards the empty string, then slices `s[:1]`.
2. `ByLength.Key` uses `strconv.Itoa(len(s))`.
3. `Group` ranges once, computing the key and appending into `out[key]`.
4. `SortedKeys` collects keys and sorts them.

## Solution

```go
func (ByFirstLetter) Key(s string) string {
	if s == "" {
		return ""
	}
	return s[:1]
}

func (ByLength) Key(s string) string { return strconv.Itoa(len(s)) }

func Group(values []string, k Keyer) map[string][]string {
	out := make(map[string][]string)
	for _, v := range values {
		key := k.Key(v)
		out[key] = append(out[key], v)
	}
	return out
}
```

## Walkthrough

`append(out[key], v)` on an absent key sees a nil slice, appends to it, and the result is stored back — no explicit initialisation needed.

## Pitfalls

- `s[0]` returns a `byte`, not a `string`; `s[:1]` keeps the type right.
- Indexing an empty string panics — guard it.
- Looping over the distinct keys and filtering the input per key, which is O(n·k).
