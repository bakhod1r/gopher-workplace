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

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Bronze, Silver, Gold
Output: 100, 200, 300
```

**Example 2:**

```
Input:  Rank(250)
Output: Silver
```

**Example 3:**

```
Input:  Rank(50)
Output: 0
```

_Explanation:_ Below Bronze.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
