# Two's complement and MinInt

## Intuition

In two's complement the most-negative value has only the **sign bit** set. For a
64-bit word that is `-1 << 63`:

```go
const MinInt = -1 << 63 // -9223372036854775808
```

`-1 << 62` is only a quarter as far and is a *valid but wrong* value — it still
fits the type, so nothing complains.

## Approach

1. A 64-bit signed minimum is `-1 << 63`.
2. The bug shifts by 62, quartering the magnitude.

## Solution

```go
const MinInt = -1 << 63

func SymmetricTo() bool { return MinInt < -(MinInt + 1) }
```

## Walkthrough

`-1 << 62` is not the true minimum; `-1 << 63` sets the sign bit alone, the most negative value.

## Pitfalls

- `-1 << 63` sets the sign bit; the compiler represents it exactly as an int.
- The asymmetry means `MinInt = -MaxInt - 1`, never `-MaxInt`.
- Signed right-shift is arithmetic (sign-extending); do bit tricks in the
  unsigned domain when you need logical shifts.
