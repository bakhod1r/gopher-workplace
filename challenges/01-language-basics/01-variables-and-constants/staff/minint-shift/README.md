# MinInt Signed Shift

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

Two's complement puts the most-negative value at `-1 << 63` for a 64-bit word.
`-1 << 62` is only a quarter of the way there. The asymmetry (`|MinInt| > MaxInt`)
is a memory-model fact the compiler encodes in the constant.

## Task

Fix the single line between the markers in [limits.go](limits.go) so `MinInt`
equals `math.MinInt64`.

## Examples

```go
MinInt            // => -9223372036854775808
int64(MinInt)     // => math.MinInt64
SymmetricTo()     // => true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Signed shift** | `-1 << 63` sets only the sign bit. |
| 2 | **Two's complement asymmetry** | `MinInt` has no positive mirror. |
| 3 | **Constant width** | The value must fit int64 exactly. |

## Hint

`-1 << 63`.

## Validate

```bash
make verify
```
