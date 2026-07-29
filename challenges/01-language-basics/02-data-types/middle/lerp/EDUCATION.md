# Linear interpolation

## Intuition

`Lerp(a,b,t) = a + (b-a)*t` walks a straight line from `a` (at t=0) to `b` (at
t=1). Values of `t` outside [0,1] extrapolate.

## Approach

1. Return a + (b-a)*t.

## Solution

```go
func Lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}
```

## Walkthrough

Lerp(2,4,0.25): 2 + (4-2)*0.25 = 2 + 0.5 = 2.5.

## Pitfalls

- `a*(1-t) + b*t` is algebraically equal but can miss `b` at t=1 due to rounding.
- No clamping here: `t=2` extrapolates beyond `b`.
- Keep everything float64; integer operands would truncate.
