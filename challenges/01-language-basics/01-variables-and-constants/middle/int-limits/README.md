# Integer Limits

**Level:** middle
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

The max/min of a machine int come straight from its bit pattern — no `math`
package needed, just constant bit operations the compiler folds.

## Task

In [limits.go](limits.go):

1. Define `MaxUint = ^uint(0)`, `MaxInt` (all bits but the sign), `MinInt`.
2. Implement `FitsInInt(v)` reporting whether `v <= MaxInt`.

## Examples

```go
MaxUint            // => 18446744073709551615 (64-bit)
MaxInt             // => 9223372036854775807
FitsInInt(0)       // => true
FitsInInt(MaxUint) // => false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Bit complement** | `^uint(0)` sets every bit → the max unsigned. |
| 2 | **Sign bit shift** | `int(^uint(0) >> 1)` clears the top bit → max signed. |
| 3 | **Two's complement** | `MinInt = -MaxInt - 1`. |

## Hint

`MaxInt = int(^uint(0) >> 1)`; then `MinInt = -MaxInt - 1`.

## Validate

```bash
make verify
```
