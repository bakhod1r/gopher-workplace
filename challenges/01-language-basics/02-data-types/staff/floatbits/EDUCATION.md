# IEEE-754 double layout

## Intuition

A `float64` is 1 sign bit, 11 exponent bits, and 52 mantissa bits. The exponent
is stored with a **bias of 1023**, so the real exponent is:

```go
raw := int((math.Float64bits(x) >> 52) & 0x7FF)
exp := raw - 1023
```

`1.0` has raw exponent 1023 → unbiased 0; `2.0` → 1024 → 1.

## Approach

1. Bug: subtracts bias 1024; float64 exponent bias is 1023.
2. Raw 11-bit exponent minus 1023 gives the unbiased exponent.
3. Fix: return raw - 1023.

## Solution

```go
import "math"

func Exponent(x float64) int {
	bits := math.Float64bits(x)
	raw := int((bits >> 52) & 0x7FF)
	return raw - 1023
}
```

## Walkthrough

x=1 has raw exponent 1023; 1023-1023=0. x=2 raw=1024 -> 1.

## Pitfalls

- Bias is `2^(11-1) - 1 = 1023` for float64; float32's bias is 127.
- `raw == 0` is subnormal/zero; `raw == 0x7FF` is Inf/NaN — special cases.
- Mask the exponent with `0x7FF` (11 bits) after shifting right by 52.
