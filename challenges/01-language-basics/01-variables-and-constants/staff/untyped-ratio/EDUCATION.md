# Untyped constant division stays exact

## The idea

Constant division follows the operands' types **at compile time**. Two integer
operands give integer division — the fraction is gone before the value is ever a
float:

```go
const R = 233 / 144   // 1   (integer division, then used as float64 -> 1.0)
const R = 233.0 / 144.0 // 1.618...  (untyped float division, full precision)
```

Make at least one operand a floating literal and the whole constant expression
is evaluated in untyped floating point, keeping precision until it is assigned.

## Why it matters

Constant expressions look like plain math, so `233 / 144` reads as ≈1.6 — but the
integer rule silently floors it. The result only rounds to a concrete type at the
point of use, so the damage is done earlier, at the division.

## Watch out

- One float operand promotes the whole expression: `233.0 / 144` also works.
- Untyped constants keep arbitrary precision; the loss here is from integer
  *semantics*, not float rounding.
- The same rule applies to `/` in `const` and in runtime expressions alike.

## Try it yourself

```go
const a = 7 / 2     // 3
const b = 7.0 / 2   // 3.5
const c = 1 / 3.0   // 0.333...
```
