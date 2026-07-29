# Quotient and Remainder

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _multiple-return_

## Context

Integer division and remainder are separate operators (`/` and `%`).
Returning both at once is a classic use of multiple return values.

## Task

Implement `DivMod` in [divmod.go](divmod.go) returning quotient then remainder.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  DivMod(17, 5)
Output: 3, 2
```

**Example 2:**

```
Input:  DivMod(10, 2)
Output: 5, 0
```

**Example 3:**

```
Input:  DivMod(4, 5)
Output: 0, 4
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Integer `/` and `%`** | `/` truncates toward zero; `%` gives the remainder. |
| 2 | **Multiple return** | Return `(q, r)` in one statement. |
| 3 | **Order matters** | Quotient first, remainder second, matching the signature. |

## Hint

Return `a / b, a % b`.

## Validate

```bash
make verify
```
