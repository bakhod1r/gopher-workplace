# Fields + Join normalization

## The idea

`strings.Fields(s)` splits on runs of any whitespace and drops empties;
`strings.Join(..., " ")` rejoins with single spaces. Together they collapse and
trim in one step — no manual padding required:

```go
strings.Join(strings.Fields(s), " ")
```

## Why it matters

Normalizing text for search keys, dedup, and comparisons is everywhere. Adding a
stray leading/trailing space defeats the whole point — keys silently differ by an
invisible character.

## Watch out

- `Fields` already handles leading/trailing and multiple spaces; don't re-add
  any.
- `Fields("")` is empty, so `Join` yields `""` — correct.
- Tabs and newlines are whitespace too and collapse to a single space.
