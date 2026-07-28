# The Euclidean algorithm

## The idea

`gcd(a,b) == gcd(b, a mod b)`, and `gcd(a,0) == a`. Iterating that recurrence
shrinks the pair fast until the second term hits 0:

```go
for b != 0 { a, b = b, a%b }
return a // (abs)
```

## Why it matters

It is the canonical example of an integer algorithm driven by `%`, and it
reduces fractions, computes LCM (`a/gcd*b`), and underpins modular arithmetic.

## Watch out

- Normalize signs: `%` can be negative in Go, so take `abs` of the result (or of
  the inputs).
- `GCD(0,0)` is defined here as 0.
- Multiple assignment `a, b = b, a%b` evaluates the right side before assigning.
