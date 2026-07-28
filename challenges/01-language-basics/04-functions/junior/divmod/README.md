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

```go
DivMod(7, 3)  // => 2, 1
DivMod(10, 5) // => 2, 0
DivMod(9, 4)  // => 2, 1
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
