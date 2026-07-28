# Day Name via Switch

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _conditionals_

## Context

An expression `switch` compares one value against many case constants — the
natural fit for a small lookup.

## Task

Implement `DayName` in [dayname.go](dayname.go) with a value switch and a default.

Do **not** change the function signature or the tests.

## Examples

```go
DayName(0) // => "Sun"
DayName(6) // => "Sat"
DayName(9) // => "?"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Expression switch** | `switch d { case 0: ... }`. |
| 2 | **default** | Catches out-of-range input. |
| 3 | **Multiple constants** | Cases can list several values with commas. |

## Hint

`switch d { case 0: return "Sun"; ...; default: return "?" }`.

## Validate

```bash
make verify
```
