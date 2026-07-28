# Fill Two Outputs

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Output parameters via pointers let a function return several results by writing
through caller-provided addresses.

## Task

Implement `MinMax` in [fillptr.go](fillptr.go).

Do **not** change the function signature or the tests.

## Examples

```go
MinMax([]int{3,-1,7}, &lo, &hi) // lo=-1, hi=7
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Output pointers** | `*min`, `*max` receive results. |
| 2 | **Single pass** | Track both extremes. |
| 3 | **Write through** | Assign `*min = ...`. |

## Hint

Seed `*min, *max = xs[0], xs[0]`, then scan updating through the pointers.

## Validate

```bash
make verify
```
