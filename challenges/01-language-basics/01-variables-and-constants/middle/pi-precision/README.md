# High-Precision Pi

**Level:** middle
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

Untyped constants carry arbitrary precision until assigned to a typed value.
A `Pi` written to 20+ digits rounds correctly to whichever float type uses it.

## Task

In [geo.go](geo.go):

1. Define `Pi` as an untyped constant with ≥20 significant digits.
2. Implement `Area(r)` returning `Pi*r*r`.

## Examples

```go
Area(1)   // => 3.14159...
Area(2)   // => 12.566...
Area(0)   // => 0
Area(0.5) // => 0.785...
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Untyped constants** | Hold arbitrary precision; no type until used. |
| 2 | **Constant rounding** | Converted to float64 only at the use site. |
| 3 | **Constant vs var** | A typed `var Pi float32` would lose digits immediately. |

## Hint

Leave the type off: `const Pi = 3.14159265358979323846`. The extra digits are
free — they only round when multiplied into a `float64`.

## Validate

```bash
make verify
```
