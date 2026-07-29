# Manual Itoa

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Formatting is the inverse of parsing: peel digits with `%10`, prepend `'0'+d`,
divide by 10, reverse.

## Task

Implement `Format(n)` (decimal, leading `-` for negatives). No `strconv`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Format(42)
Output: "42"
```

_Explanation:_ digits extracted then reversed

**Example 2:**

```
Input:  Format(-17)
Output: "-17"
```

_Explanation:_ negative gets a leading '-'

**Example 3:**

```
Input:  Format(0)
Output: "0"
```

_Explanation:_ zero special case

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Digit extraction** | `n%10` is the last digit. |
| 2 | **Char mapping** | `byte('0'+d)`. |
| 3 | **Sign + reverse** | Handle negatives; digits build in reverse. |

## Hint

Work with the absolute value, collect digits, reverse, prepend `-` if negative.

## Validate

```bash
make verify
```
