# Integer Percent Order

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`part / total * 100` evaluates left to right: for `1/4` the integer division is
0 first, so the result is always 0. Multiply before dividing.

## Task

Fix the single line between the markers in [percent.go](percent.go) so integer
percentages compute correctly (floored).

## Examples

```go
Percent(1, 4) // => 25
Percent(1, 3) // => 33
Percent(3, 4) // => 75
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Integer division** | `1/4` is 0, not 0.25. |
| 2 | **Evaluation order** | `*` and `/` are left-associative, equal precedence. |
| 3 | **Scale first** | `part * 100 / total` keeps precision. |

## Hint

`return part * 100 / total`.

## Validate

```bash
make verify
```
