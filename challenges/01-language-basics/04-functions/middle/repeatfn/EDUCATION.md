# Iterated application

## Intuition

A closure over a function and a count applies it repeatedly; zero applications is the identity.

## Approach

1. Return a closure that applies `f` to `x` `n` times.

## Solution

```go
func Repeat(f func(int) int, n int) func(int) int {
	return func(x int) int {
		for i := 0; i < n; i++ {
			x = f(x)
		}
		return x
	}
}
```

## Walkthrough

`Repeat(inc, 3)(0)`: inc applied 3 times to 0 → 3.

## Pitfalls

- n==0 must return the argument unchanged.
- The closure captures f and n by reference; don't mutate them after.
