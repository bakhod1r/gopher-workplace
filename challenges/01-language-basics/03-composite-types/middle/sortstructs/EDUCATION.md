# Sorting records

## The idea

`sort.Slice` takes a "less" function. For a stable, deterministic order, compare
the primary field, then break ties on a secondary field:

```go
c := append([]Person{}, people...)
sort.Slice(c, func(i, j int) bool {
	if c[i].Age != c[j].Age { return c[i].Age < c[j].Age }
	return c[i].Name < c[j].Name
})
```

## Why it matters

Table/report sorting is ubiquitous. Copying first respects the caller; a tie-break
makes output deterministic across runs and platforms.

## Watch out

- `sort.Slice` is **not** stable; encode tie-breaks in the comparator (or use
  `sort.SliceStable`).
- Sort mutates in place — copy if the input must survive.
- The comparator must be a strict weak ordering (no `<=`).
