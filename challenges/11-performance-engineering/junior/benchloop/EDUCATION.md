# The Loop `b.N` Expects

## Intuition

A benchmark reports `elapsed / N`. That division is only honest if the body ran the work exactly `N` times.

## Approach

1. Loop `i` from `0` to `n-1`.
2. Call `work(i)` each pass.
3. Return the count; a negative `n` simply never enters the loop.

## Solution

```go
func Run(n int, work func(i int)) int {
	calls := 0
	for i := 0; i < n; i++ {
		work(i)
		calls++
	}
	return calls
}
```

## Walkthrough

`for i := 0; i < n; i++` already handles `n <= 0`: the condition fails on the first check, so no explicit guard is needed.

## Pitfalls

- `i <= n`, which runs `n+1` times and understates ns/op.
- Hoisting the work out of the loop so it runs once.
- Panicking on a negative count instead of returning `0`.
