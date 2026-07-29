# Loop-variable capture semantics

## Intuition

Go 1.22 gives each iteration its own loop variable, so closures capture distinct values; pre-1.22 they all shared the final value.

## Approach

1. Go 1.22 gives each loop iteration its own `i`.
2. Each closure captures its own index.

## Solution

```go
func Handlers(n int) []func() int {
	var out []func() int
	for i := 0; i < n; i++ {
		out = append(out, func() int { return i })
	}
	return out
}
```

## Walkthrough

`Handlers(3)[1]()` returns 1 because the closure at position 1 captured `i == 1`.

## Pitfalls

- On Go <1.22 you would need `i := i` to shadow; on 1.26 it's automatic.
- Each closure returns its own captured index.
