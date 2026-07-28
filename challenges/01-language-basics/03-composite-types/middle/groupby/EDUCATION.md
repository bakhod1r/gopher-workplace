# Grouping with map-of-slices

## The idea

`append` to a map value works even when the key is absent, because a missing key
reads as a nil slice and appending to nil is valid:

```go
m := make(map[byte][]string)
for _, w := range words {
	if w == "" { continue }
	m[w[0]] = append(m[w[0]], w)
}
```

## Why it matters

Group-by is a core data operation (histograms, indexes, adjacency lists). The
nil-append property makes it a clean one-liner with no "does the bucket exist?"
check.

## Watch out

- You must `make` the outer map; appending is a write to it.
- The *value* nil-append is safe; the outer map nil-write is not.
- Group order follows insertion order; map key order is random.
