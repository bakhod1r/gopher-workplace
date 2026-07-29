# The Euclidean algorithm

## Intuition

`gcd(a,b) == gcd(b, a mod b)`, and `gcd(a,0) == a`. Iterating that recurrence
shrinks the pair fast until the second term hits 0:

```go
for b != 0 { a, b = b, a%b }
return a // (abs)
```

## Approach

1. Take absolute values of a and b. 2. Loop while b!=0: (a,b) = (b, a%b). 3. Return a.

## Solution

```go
func GCD(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```

## Walkthrough

GCD(12,8): a,b=12,8 -> 8,4 -> 4,0. b==0, return 4.

## Pitfalls

- Normalize signs: `%` can be negative in Go, so take `abs` of the result (or of
  the inputs).
- `GCD(0,0)` is defined here as 0.
- Multiple assignment `a, b = b, a%b` evaluates the right side before assigning.
