# Deferring within loops

## Intuition

Every `defer` in a loop is stacked; they all run at function return, not at each iteration's end — reverse order overall.

## Approach

1. Each iteration defers a call capturing `i` as an argument.
2. They fire in reverse order at function exit.

## Solution

```go
func CloseOrder(n int) (out []int) {
	for i := 0; i < n; i++ {
		defer func(v int) { out = append(out, v) }(i)
	}
	return
}
```

## Walkthrough

`CloseOrder(3)` queues 0, 1, 2 and drains them 2, 1, 0.

## Pitfalls

- Defers do NOT run at the end of each iteration — only at function return.
- For per-iteration cleanup, wrap the body in its own function.
