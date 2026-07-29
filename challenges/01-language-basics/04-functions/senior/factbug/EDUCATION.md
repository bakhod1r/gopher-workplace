# Base cases in recursion

## Intuition

A recursive definition needs a base that both terminates and returns the identity for the combining operation — 1 for products, 0 for sums.

## Approach

1. The recursion base case for factorial is `0! = 1`.
2. The bug returns 0, zeroing every product.
3. Return 1 at the base.

## Solution

```go
func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}
```

## Walkthrough

`return 0` makes `Fact(5) = 5*4*3*2*1*0 = 0`. Returning 1 at `n == 0` gives the correct 120.

## Pitfalls

- `0! == 1` by definition; the multiplicative identity is 1.
- Wrong base values are hard to spot because the recursion still terminates.
