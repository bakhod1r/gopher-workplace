# Integer limits from bit patterns

## Intuition

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

## Approach

1. `MaxUint = ^uint(0)` sets all bits.
2. `MaxInt = int(MaxUint >> 1)` clears the sign bit.
3. `MinInt = -MaxInt - 1`.

## Solution

```go
const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func FitsInInt(v uint) bool {
	return v <= uint(MaxInt)
}
```

## Walkthrough

Shifting `MaxUint` right by one drops the top bit, yielding the largest positive int; `FitsInInt` compares against it.

## Pitfalls

- `int(^uint(0))` alone is **-1**, not the max: all-bits *signed* is -1. You must
  shift first.
- `MinInt = -MaxInt - 1`: two's complement is asymmetric, one more negative than
  positive.
- A typed `int` constant that overflows is a **compile error**, so build wide
  values as `uint` and convert, or keep them untyped.
