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

**Example 1:**

```
Input:  MinMax([3 -1 7])
Output: -1, 7
```

**Example 2:**

```
Input:  MinMax([5])
Output: 5, 5
```

**Example 3:**

```
Input:  MinMax([2 2 2])
Output: 2, 2
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
