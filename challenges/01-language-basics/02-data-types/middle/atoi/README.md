# Manual Atoi

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Parsing "123" means folding digits into a running total: `total = total*10 +
digit`. A character's digit value is `c - '0'`.

## Task

Implement `Parse(s)` (optional leading `-`), returning `(value, ok)`; `ok=false`
on any non-digit. Don't use `strconv`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Parse("42")
Output: (42, true)
```

_Explanation:_ digit-by-digit: 4, then 4*10+2=42

**Example 2:**

```
Input:  Parse("-17")
Output: (-17, true)
```

_Explanation:_ leading '-' sets sign, magnitude 17 negated

**Example 3:**

```
Input:  Parse("1a")
Output: (0, false)
```

_Explanation:_ 'a' is not a digit -> invalid

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Digit value** | `c - '0'` maps '0'..'9' to 0..9. |
| 2 | **Horner fold** | `n = n*10 + d` accumulates left to right. |
| 3 | **Validation** | Reject empty and non-digit input. |

## Hint

Range over bytes, check `c >= '0' && c <= '9'`, fold, apply sign at the end.

## Validate

```bash
make verify
```
