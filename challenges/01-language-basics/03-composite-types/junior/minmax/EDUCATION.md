# Seed-and-scan for extremes

## The idea

Don't seed min/max with 0 or huge sentinels — seed with the **first element**,
then scan the rest:

```go
if len(xs) == 0 { return 0, 0, false }
mn, mx := xs[0], xs[0]
for _, x := range xs[1:] { if x < mn { mn = x }; if x > mx { mx = x } }
```

## Why it matters

Seeding from the data avoids sentinel bugs (a list of all-negative numbers breaks
`max := 0`). The `(v, ok)` return handles the empty case without a fake value.

## Watch out

- Guard empty before indexing `xs[0]`.
- Seeding with 0 is a classic bug for all-negative or all-large inputs.
- `slices.Min`/`slices.Max` (Go 1.21+) panic on empty — the `ok` form is safer.
