# Grouping with map-of-slices

## Intuition

`append` to a map value works even when the key is absent, because a missing key
reads as a nil slice and appending to nil is valid:

```go
m := make(map[byte][]string)
for _, w := range words {
	if w == "" { continue }
	m[w[0]] = append(m[w[0]], w)
}
```

## Approach

1. Create map[byte][]string.
2. For each word, skip if empty.
3. Take first byte w[0] as key.
4. Append the word to that key's slice.
5. Return map.

## Solution

```go
func ByFirst(words []string) map[byte][]string {
	out := map[byte][]string{}
	for _, w := range words {
		if w == "" {
			continue
		}
		k := w[0]
		out[k] = append(out[k], w)
	}
	return out
}
```

## Walkthrough

"apple"->a:[apple]; "banana"->b:[banana]; "avocado"->a:[apple,avocado]; "cherry"->c:[cherry]; "" skipped.

## Pitfalls

- You must `make` the outer map; appending is a write to it.
- The *value* nil-append is safe; the outer map nil-write is not.
- Group order follows insertion order; map key order is random.
