# Closures capturing enclosing state

## Intuition

A function literal keeps a live reference to the variables it uses from the outer scope; those variables survive as long as the closure does.

## Approach

1. Declare `n := 0` in the enclosing scope.
2. Return a closure that increments and returns `n`.
3. Each call to `MakeCounter` gets its own `n`.

## Solution

```go
func MakeCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
```

## Walkthrough

The returned closure captures `n` by reference; successive calls yield 1, 2, 3. A fresh `MakeCounter` starts a new `n`.

## Pitfalls

- Each call to the factory creates a NEW captured variable — instances don't share state.
- The captured variable escapes to the heap; that's expected.
