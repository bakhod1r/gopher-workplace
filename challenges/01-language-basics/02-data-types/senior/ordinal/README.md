# English Ordinals

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A UI prints ranks as `1st, 2nd, 3rd, 4th…`. The code keys only on the last digit,
so it wrongly prints `11st, 12nd, 13rd` — but 11–13 are always `th`.

## Task

Fix the `switch` between the markers in [ordinal.go](ordinal.go) so 11–13 (and
111–113, etc.) use `"th"`.

## Examples

```go
Format(11)  // => "11th"
Format(21)  // => "21st"
Format(113) // => "113th"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Last-digit rule** | 1→st, 2→nd, 3→rd, else th. |
| 2 | **Teens exception** | `n%100` in 11..13 → th. |
| 3 | **Switch on two keys** | Guard the teens before the last digit. |

## Hint

First: `if n%100 >= 11 && n%100 <= 13 { return "th" }`, then switch on `n%10`.

## Validate

```bash
make verify
```
