# Min and Max

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _multiple-return_

## Context

A Go function can return more than one value. Scanning a slice once and
returning both extremes is cheaper and clearer than two passes.

## Task

Implement `MinMax` in [minmax.go](minmax.go) so it returns the smallest and
largest element of `xs` in one pass.

Do **not** change the function signature or the tests.

## Examples

```go
MinMax([]int{5, -2, 9, 0}) // => -2, 9
MinMax([]int{3})          // => 3, 3
MinMax([]int{1, 2, 3})    // => 1, 3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Multiple return values** | A function may return a tuple `(min, max int)`. |
| 2 | **Seed from first element** | Initialise both extremes to `xs[0]`, not to 0. |
| 3 | **Single pass** | One `for range` updates both bounds. |

## Hint

Start `min, max = xs[0], xs[0]`, then loop from index 1 comparing each element.

## Validate

```bash
make verify
```
