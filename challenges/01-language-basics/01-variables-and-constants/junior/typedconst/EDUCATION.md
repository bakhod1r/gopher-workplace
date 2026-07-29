# Typed vs untyped constants

## Intuition

A constant may be declared with a type or without one, and the two behave very
differently at the point of use.

```go
const MaxBatch byte = 200   // typed:   a byte, everywhere, always
const Retries = 3           // untyped: an integer *kind*, type decided later
```

An **untyped** constant is converted implicitly to whatever its context needs.
A **typed** constant is not converted at all — mixing it with another type is a
compile error until you say so explicitly.

```go
var f float64 = Retries     // fine: adopts float64
var r rune = Retries        // fine: adopts rune
var n int = MaxBatch        // compile error: byte is not int
var n int = int(MaxBatch)   // fine: explicit
```

## Which to reach for

Default to **untyped**. It composes with any numeric type without ceremony,
which is why the standard library's constants (`math.Pi`, `time.Nanosecond`'s
multipliers) are mostly untyped.

Choose **typed** when the type is part of the meaning and you want the compiler
to enforce it — an enum member (`const Free Tier = iota`), a value that must
stay inside a byte, a duration that should never be confused with a plain
integer.

## Conversions truncate, silently

Narrowing a value into a smaller type wraps modulo its range. No panic, no
warning:

```go
byte(255)   // 255
byte(256)   // 0
byte(456)   // 200
int8(200)   // -56
```

So the *direction* of a conversion is a correctness decision, not a style
choice. Widening (`int(aByte)`) can never lose information; narrowing can. When
comparing values of two sizes, widen the small one:

```go
n <= int(MaxBatch)     // right: nothing wraps
byte(n) <= MaxBatch    // wrong: n == 256 becomes 0 and passes
```

## Constant conversions are checked at compile time

Because constants are known to the compiler, an out-of-range *constant* is
caught rather than wrapped:

```go
const c byte = 256      // compile error: constant 256 overflows byte
var n = 300
b := byte(n)            // no error: runtime conversion, wraps to 44
```

Same-looking code, opposite outcome — the difference is whether the value is a
constant.

## Approach

1. `MaxBatch byte = 200` is typed; comparing with an `int` needs `int(MaxBatch)`.
2. `Retries = 3` is untyped; it adapts to float64 in `Budget`.
3. `Fits` guards negatives.

## Solution

```go
const (
	MaxBatch byte = 200
	Retries       = 3
)

func Fits(n int) bool {
	return n >= 0 && n <= int(MaxBatch)
}

func Budget(base float64) float64 {
	return base * Retries
}
```

## Walkthrough

`Fits(200)` converts `MaxBatch` to int and compares; `Budget(1.5)` multiplies by the untyped 3 → 4.5 with no conversion.

## Pitfalls

- `byte` is an alias for `uint8`, `rune` for `int32`. Aliases are the *same*
  type, so no conversion is needed between `byte` and `uint8`.
- A typed constant restricts arithmetic too: `MaxBatch * 2` is a `byte`
  expression and overflows at 400 → 144.
- Do not sprinkle conversions to silence the compiler. Each one is a claim that
  the value fits; make sure it does.
