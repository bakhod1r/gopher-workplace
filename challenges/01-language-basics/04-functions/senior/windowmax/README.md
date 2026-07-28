# Sliding Window Bound

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

A sliding window of width k advances by adding the entering element `xs[i]` and
removing the leaving one `xs[i-k]`. The planted bug removes `xs[i-k+1]` — off by
one — so the running sum drifts.

## Task

Fix the window update in [windowmax.go](windowmax.go).

Do **not** change the function signature or the tests.

## Examples

```go
MaxWindow([]int{1,2,3,4,5}, 2)     // => 9
MaxWindow([]int{2,1,5,1,3,2}, 3)   // => 9
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Sliding window** | Add entering, subtract leaving element. |
| 2 | **Index arithmetic** | The element leaving is `xs[i-k]`. |
| 3 | **Off-by-one** | `i-k+1` drops the wrong element. |

## Hint

The element leaving the window is `xs[i-k]`, not `xs[i-k+1]`.

## Validate

```bash
make verify
```
