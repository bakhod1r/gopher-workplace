# Integer limits from bit patterns

## The idea

The maximum and minimum of a machine integer are just bit patterns you can build
with constant operations — no `math` package needed:

```go
const MaxUint = ^uint(0)          // all bits set
const MaxInt  = int(^uint(0) >> 1) // clear the sign bit
const MinInt  = -MaxInt - 1        // two's complement
```

`^` is bitwise NOT. `^uint(0)` flips every bit of zero, giving the largest
unsigned value. Shifting right by one clears the top (sign) bit, giving the
largest **signed** value.

## Why it matters

These fold at compile time and adapt to the platform word size (32- or 64-bit),
because `uint`/`int` are platform-width. `math.MaxInt` in the stdlib is defined
exactly this way.

## Watch out

- `int(^uint(0))` alone is **-1**, not the max: all-bits *signed* is -1. You must
  shift first.
- `MinInt = -MaxInt - 1`: two's complement is asymmetric, one more negative than
  positive.
- A typed `int` constant that overflows is a **compile error**, so build wide
  values as `uint` and convert, or keep them untyped.

## Try it yourself

```go
const MaxUint8 = ^uint8(0)      // 255
const MaxInt8  = int8(^uint8(0) >> 1) // 127
const MinInt8  = -MaxInt8 - 1   // -128
```
