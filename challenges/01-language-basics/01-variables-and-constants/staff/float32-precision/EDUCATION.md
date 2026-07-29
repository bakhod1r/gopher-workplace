# Float width and error growth

## Intuition

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

## Approach

1. `float32` accumulates 0.1 with visible error.
2. Use `float64` for a tighter total.

## Solution

```go
func SumTenths(n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += 0.1
	}
	return float64(total)
}
```

## Walkthrough

Summing 0.1 in float32 drifts; float64 keeps the running total close to the expected value.

## Pitfalls

- Never compare floats with `==`; use an epsilon sized for the type and the
  operation count.
- Accumulate in the widest reasonable type, narrow only at the boundary.
- Summation order and techniques (Kahan) matter once N is large.
