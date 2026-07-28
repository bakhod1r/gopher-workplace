# Float Tolerance Equality

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

`0.1 + 0.2 != 0.3` in binary floats. Comparing floats needs a tolerance, not
`==`.

## Task

Implement `Equal(a, b, eps)` = true when `|a-b| <= eps`; NaN never equal.

## Examples

```go
Equal(0.1+0.2, 0.3, 1e-9) // => true
Equal(1.0, 1.001, 1e-9)   // => false
Equal(NaN, NaN, 1)        // => false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Float inexactness** | Most decimals aren't exact in binary. |
| 2 | **Absolute tolerance** | `math.Abs(a-b) <= eps`. |
| 3 | **NaN** | `NaN` fails every comparison; exclude it. |

## Hint

`return math.Abs(a-b) <= eps` — NaN makes `Abs` NaN, and `NaN <= eps` is false,
so NaN is handled naturally.

## Validate

```bash
make verify
```
