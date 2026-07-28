# "Copy" that isn't

## The idea

`out := xs` copies the slice header (pointer/len/cap) — both names point at the
same backing array. `sort.Ints(out)` therefore sorts `xs` too. Duplicate the data
to isolate:

```go
out := append([]int{}, xs...) // independent copy
sort.Ints(out)
```

## Why it matters

Functions advertised as non-mutating ("SortedCopy", "Cleaned") must not alter
their inputs. The shared-header trap makes an in-place sort masquerade as a copy —
a surprising, action-at-a-distance bug.

## Watch out

- Assigning a slice never copies its elements.
- `slices.Sorted`/`slices.Clone` avoid the footgun.
- `sort.Ints`, `sort.Slice`, and `slices.Sort` all mutate in place.
