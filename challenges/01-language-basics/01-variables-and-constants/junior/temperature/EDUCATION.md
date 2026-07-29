# Untyped constants and constant kinds

## Intuition

An untyped constant has no type — it has a **kind**: integer, float, rune,
string, bool, complex. The kind decides how arithmetic behaves; the type is
chosen later, from the context where the constant is used.

```go
const ratio = 9 / 5      // integer kind: 1
const ratio = 9.0 / 5.0  // float kind: 1.8
```

Both operands of `9 / 5` are integer constants, so the division is integer
division and the fraction is gone before any variable is involved. Writing one
operand with a decimal point makes the whole expression float-kinded.

## The trap: it compiles either way

```go
func ToF(c float64) float64 {
	return c*9/5 + 32     // fine — c is float64, so the whole expression is
}

const ratio = 9 / 5       // 1
func ToF(c float64) float64 {
	return c*ratio + 32   // silently wrong: c * 1
}
```

In the first version the untyped constants join a `float64` expression and adopt
that type. In the second, the division happened *between two constants* first,
in integer kind. Nothing warns you: `1` is a perfectly good number.

## Constants are arbitrary precision

Untyped constant arithmetic is exact and unbounded until the value is assigned:

```go
const big = 1 << 100          // fine as a constant
var x int = big               // compile error: overflows int
const third = 1.0 / 3.0       // exact, far beyond float64 precision
var f float64 = third         // rounded here, at the point of use
```

That is why constant expressions can be written naturally without worrying about
intermediate overflow — the rounding happens once, when the value lands in a
typed variable.

## Kind conversion rules

| Expression | Kind | Value |
|------------|------|-------|
| `9 / 5` | integer | 1 |
| `9.0 / 5` | float | 1.8 |
| `float64(9) / 5` | typed float64 | 1.8 |
| `9 / 5.0` | float | 1.8 |

One float operand is enough: mixing kinds in a constant expression promotes to
the "wider" one (integer → rune → float → complex).

## Approach

1. Use floating-point arithmetic so `9/5` is 1.8, not 1.
2. Compute `c*9.0/5.0 + 32`.

## Solution

```go
const freezingF = 32

func CToF(c float64) float64 {
	return c*9.0/5.0 + freezingF
}
```

## Walkthrough

`CToF(100)`: `100*1.8 = 180`, plus 32 → 212. Integer `9/5` would truncate to 1 and break it.

## Pitfalls

- Give a constant a type only when you want to *restrict* it. `const ratio
  float64 = 1.8` is fine here; `const n int = 3` stops `n` being used as a
  `float64` later.
- `9 / 5` inside a `float64` expression is still 1 — the parentheses matter:
  `c * 9 / 5` is `(c*9)/5` (correct, because `c*9` is float64) but `c * (9/5)`
  is `c * 1`.
- Integer division truncates toward zero, it does not round: `7/2 == 3`,
  `-7/2 == -3`.
