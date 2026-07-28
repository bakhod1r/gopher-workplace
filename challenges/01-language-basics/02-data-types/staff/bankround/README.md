# Banker's Rounding

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A financial report must round half-to-even to avoid the upward bias of
round-half-away. The code uses `math.Round`, which rounds half **away** from zero
(`2.5 -> 3`), skewing totals.

## Task

Fix the body between the markers in [bankround.go](bankround.go) to round ties to
the even neighbor. `math.RoundToEven` does exactly this.

## Examples

```go
Round(2.5) // => 2
Round(3.5) // => 4
Round(0.5) // => 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Round modes** | Half-away vs half-to-even. |
| 2 | **Bias** | Half-away skews sums upward. |
| 3 | **math.RoundToEven** | The stdlib half-to-even. |

## Hint

`return math.RoundToEven(x)`.

## Validate

```bash
make verify
```
