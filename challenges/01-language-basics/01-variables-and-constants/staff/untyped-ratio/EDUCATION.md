# Untyped constant division stays exact

## Intuition

Constant division follows the operands' types **at compile time**. Two integer
operands give integer division — the fraction is gone before the value is ever a
float:

```go
const R = 233 / 144   // 1   (integer division, then used as float64 -> 1.0)
const R = 233.0 / 144.0 // 1.618...  (untyped float division, full precision)
```

Make at least one operand a floating literal and the whole constant expression
is evaluated in untyped floating point, keeping precision until it is assigned.

## Approach

1. `233 / 144` with integer operands truncates to 1.
2. Make one operand a float: `233.0 / 144`.

## Solution

```go
const GoldenApprox = 233.0 / 144

func Value() float64 { return GoldenApprox }
```

## Walkthrough

Integer division loses the fraction; `233.0 / 144` keeps full precision → ~1.618.

## Pitfalls

- One float operand promotes the whole expression: `233.0 / 144` also works.
- Untyped constants keep arbitrary precision; the loss here is from integer
  *semantics*, not float rounding.
- The same rule applies to `/` in `const` and in runtime expressions alike.
