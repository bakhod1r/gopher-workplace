# Two's complement and MinInt

## The idea

In two's complement the most-negative value has only the **sign bit** set. For a
64-bit word that is `-1 << 63`:

```go
const MinInt = -1 << 63 // -9223372036854775808
```

`-1 << 62` is only a quarter as far and is a *valid but wrong* value — it still
fits the type, so nothing complains.

## Why it matters

`|MinInt| > MaxInt`: the range is asymmetric by exactly one. Negating `MinInt`
overflows (there is no `+MinInt`), which breaks naive `abs`, division, and
`-x` all over. Getting the constant itself right is the foundation.

## Watch out

- `-1 << 63` sets the sign bit; the compiler represents it exactly as an int.
- The asymmetry means `MinInt = -MaxInt - 1`, never `-MaxInt`.
- Signed right-shift is arithmetic (sign-extending); do bit tricks in the
  unsigned domain when you need logical shifts.

## Try it yourself

```go
const MinInt8 = -1 << 7 // -128
const MaxInt8 = 1<<7 - 1 // 127
-MinInt8 == MinInt8      // true (int8): overflow, no positive mirror
```
