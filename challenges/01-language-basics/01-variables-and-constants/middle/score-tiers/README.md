# Score Tiers

**Level:** middle
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

Reward thresholds climb in fixed steps. An `iota` expression writes the pattern
once and lets the block repeat it.

## Task

In [tiers.go](tiers.go):

1. Define `Bronze=100, Silver=200, Gold=300` using `(iota+1)*100` written once.
2. Implement `Rank(score)` returning the highest tier ≤ score, else 0.

## Examples

```go
Silver          // => 200
Rank(50)        // => 0
Rank(150)       // => 100 (Bronze)
Rank(350)       // => 300 (Gold)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **iota in expressions** | `(iota+1)*100` scales the counter. |
| 2 | **Implicit repetition** | Blank RHS lines reuse the previous expression. |
| 3 | **Ordered comparison** | Compare `Tier` values to bucket a score. |

## Hint

Only `Bronze` needs `Tier = (iota + 1) * 100`; `Silver` and `Gold` on bare lines
inherit it.

## Validate

```bash
make verify
```
