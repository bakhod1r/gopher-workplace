# Manual Atoi

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Parsing "123" means folding digits into a running total: `total = total*10 +
digit`. A character's digit value is `c - '0'`.

## Task

Implement `Parse(s)` (optional leading `-`), returning `(value, ok)`; `ok=false`
on any non-digit. Don't use `strconv`.

## Examples

```go
Parse("42")  // => 42, true
Parse("-17") // => -17, true
Parse("1a")  // => 0, false
Parse("")    // => 0, false
```

## Topics to Master

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
