# Deriving MaxInt safely

## Intuition

`^uint(0)` is all-bits-set: the largest **unsigned** value. Interpreting those
same bits as a **signed** int gives -1, because the top bit is the sign bit. To
get the largest positive signed value, clear that sign bit with `>> 1`:

```go
var allBits = ^uint(0)      // e.g. 2^64 - 1
MaxInt := int(allBits >> 1) // 2^63 - 1
```

The right shift happens in `uint` (logical shift, zero-filled), then convert.

## Approach

1. `int(allBits)` is -1 (all bits including the sign bit).
2. Shift right by one to clear the sign bit: `int(allBits >> 1)`.

## Solution

```go
var allBits = ^uint(0)

var MaxInt = int(allBits >> 1)

func Overflows() bool { return MaxInt+1 < MaxInt }
```

## Walkthrough

All bits set as a signed int is -1, not the max. Clearing the top bit yields the largest positive value.

## Pitfalls

- Convert **after** shifting: `int(allBits >> 1)`. Shifting a signed -1 right is
  an *arithmetic* shift (sign-extends) and stays -1.
- A typed-int constant that overflows is a compile error; build wide values in
  `uint`.
- The stdlib exposes `math.MaxInt` (Go 1.17+) so you rarely hand-roll this.
