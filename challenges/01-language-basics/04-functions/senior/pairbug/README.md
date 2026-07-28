# Adjacent-Pair Bound

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

Reading `xs[i+1]` requires the loop to stop at `len(xs)-1`; iterating to
`len(xs)` reads one element past the slice and panics.

## Task

Fix the loop bound in [pairbug.go](pairbug.go).

Do **not** change the function signature or the tests.

## Examples

```go
SumDiffs([1 4 9]) // => 8
SumDiffs([5])     // => 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Look-ahead bound** | Accessing i+1 needs `i < len-1`. |
| 2 | **Empty/single guard** | Zero pairs for <2 elements. |
| 3 | **Index safety** | Out-of-range indexing panics. |

## Hint

Loop `for i := 0; i < len(xs)-1; i++`.

## Validate

```bash
make verify
```
