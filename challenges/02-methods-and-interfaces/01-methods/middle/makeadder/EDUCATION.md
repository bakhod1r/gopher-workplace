# Method Closure Factories

## Intuition

A method can return an anonymous function. That function forms a closure over
the method's arguments and receiver. Since `Number` is passed by value, the
closure captures a copy of `n`.

## Approach

1. Return an anonymous function.

## Solution

```go
func (n Number) Adder() func(int) int {
	return func(x int) int {
		return n.Val + x
	}
}
```

## Walkthrough

- `n = Number{5}`.
- `Adder()` returns `func(x int) int { return 5 + x }`.

## Pitfalls

- If `Number` used a pointer receiver `*Number`, the closure would capture the
  pointer, and later changes to `n.Val` would affect the closure's output.
