# Banker's Rounding

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A financial report must round half-to-even to avoid the upward bias of
round-half-away. The code uses `math.Round`, which rounds half **away** from zero
(`2.5 -> 3`), skewing totals.

## Task

Fix the body between the markers in [bankround.go](bankround.go) to round ties to
the even neighbor. `math.RoundToEven` does exactly this.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  2.5
Output: 2
```

_Explanation:_ Tie rounds to the even neighbor 2, not 3.

**Example 2:**

```
Input:  3.5
Output: 4
```

_Explanation:_ Tie rounds to even 4.

**Example 3:**

```
Input:  -2.5
Output: -2
```

_Explanation:_ Tie rounds to even -2, not -3.

**Example 4:**

```
Input:  2.6
Output: 3
```

_Explanation:_ Non-tie rounds normally to nearest.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Round modes** | Half-away vs half-to-even. |
| 2 | **Bias** | Half-away skews sums upward. |
| 3 | **math.RoundToEven** | The stdlib half-to-even. |

## Hint

`return math.RoundToEven(x)`.

## Validate

```bash
make verify
```
