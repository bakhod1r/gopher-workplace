# Linear interpolation

## The idea

`Lerp(a,b,t) = a + (b-a)*t` walks a straight line from `a` (at t=0) to `b` (at
t=1). Values of `t` outside [0,1] extrapolate.

## Why it matters

Animation, gradients, and numeric blending all use lerp. The form `a+(b-a)*t` is
preferred over `a*(1-t)+b*t` because it hits the endpoints exactly.

## Watch out

- `a*(1-t) + b*t` is algebraically equal but can miss `b` at t=1 due to rounding.
- No clamping here: `t=2` extrapolates beyond `b`.
- Keep everything float64; integer operands would truncate.
