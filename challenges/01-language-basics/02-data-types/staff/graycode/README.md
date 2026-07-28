# Gray Code Encoding

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A rotary encoder outputs Gray code so only one bit changes per step. The formula
is `x ^ (x >> 1)`, but the code shifts **left**, breaking the single-bit-change
property.

## Task

Fix the shift between the markers in [graycode.go](graycode.go).

## Examples

```go
ToGray(2) // => 3
ToGray(3) // => 2
ToGray(4) // => 6
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Gray code** | `g = x ^ (x >> 1)`. |
| 2 | **Single-bit change** | Consecutive codes differ by one bit. |
| 3 | **Shift direction** | Right shift folds high bits down. |

## Hint

`x ^ (x >> 1)`.

## Validate

```bash
make verify
```
