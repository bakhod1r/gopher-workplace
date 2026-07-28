# Prefix-Sum Query

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

A prefix array `pre` satisfies `sum(xs[l:r]) == pre[r] - pre[l]`. The bug
returns only `pre[r]`, which is the sum from 0 to r — it omits subtracting the
left prefix.

## Task

Fix [prefixbug.go](prefixbug.go) so range sums are correct.

Do **not** change the function signature or the tests.

## Examples

```go
RangeSum([1 2 3 4 5], 1, 4) // => 9
RangeSum([1 2 3 4 5], 0, 5) // => 15
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Prefix sums** | `pre[i]` = sum of first i elements. |
| 2 | **Range formula** | `pre[r] - pre[l]`. |
| 3 | **Half-open interval** | l inclusive, r exclusive. |

## Hint

Subtract the left prefix: `return pre[r] - pre[l]`.

## Validate

```bash
make verify
```
