# IEEE-754 double layout

## The idea

A `float64` is 1 sign bit, 11 exponent bits, and 52 mantissa bits. The exponent
is stored with a **bias of 1023**, so the real exponent is:

```go
raw := int((math.Float64bits(x) >> 52) & 0x7FF)
exp := raw - 1023
```

`1.0` has raw exponent 1023 → unbiased 0; `2.0` → 1024 → 1.

## Why it matters

Reading float internals powers fast `frexp`, ULP comparisons, and custom
formatting. The bias constant is easy to misremember (1023, not 1024); off by one
and every exponent — and anything derived from it — is wrong.

## Watch out

- Bias is `2^(11-1) - 1 = 1023` for float64; float32's bias is 127.
- `raw == 0` is subnormal/zero; `raw == 0x7FF` is Inf/NaN — special cases.
- Mask the exponent with `0x7FF` (11 bits) after shifting right by 52.
