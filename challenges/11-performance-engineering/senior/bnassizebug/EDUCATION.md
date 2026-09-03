# `b.N` Is Not The Input Size

## Intuition

`b.N` answers "how many times", never "how big". Using it as a size couples the measurement to a number the harness is free to change between runs.

## Approach

1. Multiply the iteration count by the input size.

## Solution

```go
func Work(n, size int) int {
	if n <= 0 || size <= 0 {
		return 0
	}
	return n * size
}
```

## Walkthrough

With the bug, `PerOp(n, 10)` returns `n`: the reported cost per operation is literally the iteration count, so raising `-benchtime` "proves" a scaling problem that exists only in the benchmark.

## Pitfalls

- `make([]byte, b.N)` inside the body, the most common form of this bug.
- Allocating the fixture once but slicing it to `b.N` elements, which is the same mistake wearing a hat.
- Varying the size with `b.Run` and *also* letting `b.N` in, so both dimensions move at once.
