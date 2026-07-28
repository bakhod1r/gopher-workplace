# Deriving MaxInt safely

## The idea

`^uint(0)` is all-bits-set: the largest **unsigned** value. Interpreting those
same bits as a **signed** int gives -1, because the top bit is the sign bit. To
get the largest positive signed value, clear that sign bit with `>> 1`:

```go
var allBits = ^uint(0)      // e.g. 2^64 - 1
MaxInt := int(allBits >> 1) // 2^63 - 1
```

The right shift happens in `uint` (logical shift, zero-filled), then convert.

## Why it matters

`int(^uint(0))` is a classic trap: it compiles as a *runtime* conversion (when
the operand is a variable) and yields -1, not the max. As a constant it will not
even compile (overflow). Understanding sign bits is what separates the two.

## Watch out

- Convert **after** shifting: `int(allBits >> 1)`. Shifting a signed -1 right is
  an *arithmetic* shift (sign-extends) and stays -1.
- A typed-int constant that overflows is a compile error; build wide values in
  `uint`.
- The stdlib exposes `math.MaxInt` (Go 1.17+) so you rarely hand-roll this.

## Try it yourself

```go
var u = ^uint8(0)     // 255
int8(u)               // -1
int8(u >> 1)          // 127
```
