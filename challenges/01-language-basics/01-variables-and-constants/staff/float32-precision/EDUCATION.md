# Float width and error growth

## The idea

`0.1` cannot be represented exactly in binary floating point; each store rounds.
`float32` has 24 bits of mantissa, `float64` has 53, so the per-operation error
is far larger in `float32` and accumulates faster:

```go
var t float32
for i := 0; i < 10; i++ { t += 0.1 }
// t drifts from 1.0 by ~1e-7, past a 1e-9 tolerance
```

The literal `0.1` is an *untyped* constant of full precision — it is the
**variable's** type that fixes how much precision survives each `+=`.

## Why it matters

Money, physics steps, and running sums accumulate thousands of operations. The
type of the accumulator, not the literals, decides whether the result stays
within tolerance. Choosing `float32` to "save memory" can silently blow a
correctness budget.

## Watch out

- Never compare floats with `==`; use an epsilon sized for the type and the
  operation count.
- Accumulate in the widest reasonable type, narrow only at the boundary.
- Summation order and techniques (Kahan) matter once N is large.

## Try it yourself

```go
var a float32 = 0
for i := 0; i < 100; i++ { a += 0.1 }
// compare a to 10.0 — note the drift
```
