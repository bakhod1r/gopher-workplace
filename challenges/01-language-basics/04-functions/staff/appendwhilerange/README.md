# Append While Ranging

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

The loop bound must be over the INPUT length, `len(xs)`, not the growing output
`len(out)`. Bounding on `len(out)` (which starts at 0) skips the whole loop, and
any self-referential growing bound is a footgun.

## Task

Fix the loop bound in [appendwhilerange.go](appendwhilerange.go).

Do **not** change the function signature or the tests.

## Examples

```go
DupAll([3 5]) // => [3 6 5 10]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Iterate the source length** | Bound on `len(xs)`, not `len(out)`. |
| 2 | **Growing collections** | Don't bound a loop on the slice it grows. |
| 3 | **Snapshot the length** | Capture the fixed count before the loop. |

## Hint

Iterate over the input: `for i := 0; i < len(xs); i++` (or `for i := range xs`).

## Validate

```bash
make verify
```
