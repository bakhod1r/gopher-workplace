# Linear Interpolation

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Lerp blends two values by a factor `t`: at `t=0` you get `a`, at `t=1` you get
`b`.

## Task

Implement `Lerp(a, b, t)` = `a + (b-a)*t`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Lerp(0, 10, 0.5)
Output: 5
```

_Explanation:_ halfway between 0 and 10

**Example 2:**

```
Input:  Lerp(2, 4, 0.25)
Output: 2.5
```

_Explanation:_ a quarter of the way from 2 to 4

**Example 3:**

```
Input:  Lerp(0, 10, 1)
Output: 10
```

_Explanation:_ t=1 gives b

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Interpolation formula** | `a + (b-a)*t`. |
| 2 | **Float arithmetic** | All operands float64. |
| 3 | **Endpoints** | t=0→a, t=1→b exactly. |

## Hint

`return a + (b-a)*t`.

## Validate

```bash
make verify
```
