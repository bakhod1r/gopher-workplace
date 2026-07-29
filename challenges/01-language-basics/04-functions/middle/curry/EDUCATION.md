# Currying with nested closures

## Intuition

Each closure captures its argument and returns the next; the innermost sees the whole accumulated environment.

## Approach

1. Return nested closures, each capturing one argument.
2. The innermost sums all three.

## Solution

```go
func Add3() func(int) func(int) func(int) int {
	return func(a int) func(int) func(int) int {
		return func(b int) func(int) int {
			return func(c int) int {
				return a + b + c
			}
		}
	}
}
```

## Walkthrough

`Add3()(1)(2)(3)` captures 1, then 2, then adds 3 → 6.

## Pitfalls

- Every inner closure captures its enclosing arguments by reference.
- The type signature nests exactly as the calls do.
