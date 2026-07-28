# Celsius to Fahrenheit

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

`c*9/5` in integers truncates. Converting to `float64` first keeps the fraction,
so `37°C` gives `98.6`, not `98`.

## Task

Implement `ToF(c)` returning `c*9/5 + 32` as a `float64`.

## Examples

```go
ToF(0)   // => 32
ToF(100) // => 212
ToF(37)  // => 98.6
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Type conversion** | `float64(c)` widens the int before dividing. |
| 2 | **Integer vs float division** | `9/5` in ints is 1; in floats 1.8. |
| 3 | **Mixed arithmetic** | Go has no implicit int↔float; convert explicitly. |

## Hint

`return float64(c)*9/5 + 32`.

## Validate

```bash
make verify
```
