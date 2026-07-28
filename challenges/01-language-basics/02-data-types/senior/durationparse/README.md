# Duration Minutes

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A scheduler parses timeouts like `1h30m`. Minutes are converted with the wrong
factor (×6 instead of ×60), so every timeout with minutes is ten times too
short.

## Task

Fix the minute multiplier between the markers in
[durationparse.go](durationparse.go).

## Examples

```go
Seconds("2m")     // => 120
Seconds("1h30m")  // => 5400
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Unit multipliers** | h=3600, m=60, s=1 seconds. |
| 2 | **Digit fold** | Accumulate the number before its unit. |
| 3 | **Trailing unit** | A number with no unit is invalid. |

## Hint

`total += num * 60`.

## Validate

```bash
make verify
```
