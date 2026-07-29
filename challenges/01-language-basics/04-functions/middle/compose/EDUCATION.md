# Composing functions

## Intuition

A closure capturing two functions yields their composition; order matters — `f(g(x))` applies `g` first.

## Approach

1. Return a closure computing `f(g(x))`.
2. `g` runs first, then `f`.

## Solution

```go
func Compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}
```

## Walkthrough

`Compose(inc, double)(3)`: double(3)=6, inc(6)=7.

## Pitfalls

- `Compose(f,g)` is f-after-g, not g-after-f.
- Both captured functions live as long as the returned closure.
