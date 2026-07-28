# Finite Float Check

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Floating point has special values: `NaN` (not-a-number) and `±Inf`. They arise
from `0.0/0.0`, `math.Sqrt(-1)`, overflow, etc. and must often be rejected.

## Task

Implement `Finite(x)` returning true only for normal finite numbers.

## Examples

```go
Finite(3.14)          // => true
Finite(math.NaN())    // => false
Finite(math.Inf(1))   // => false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **NaN** | Not equal to anything, even itself. |
| 2 | **±Inf** | Result of overflow / divide by zero. |
| 3 | **math.IsNaN / IsInf** | The correct way to test these. |

## Hint

`return !math.IsNaN(x) && !math.IsInf(x, 0)`.

## Validate

```bash
make verify
```
