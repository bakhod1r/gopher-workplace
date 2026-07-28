# Euclidean algorithm as a while-loop

## The idea

Go's `for cond {}` is the while loop; parallel assignment `a, b = b, a%b` advances both values without a temporary.

## Why it matters

It's a compact example of loop-until-condition with simultaneous update, common in numeric routines.

## Watch out

- `GCD(0, n)` should return n; the loop handles it since b!=0 runs once.
- Parallel assignment evaluates the whole right side before assigning.
