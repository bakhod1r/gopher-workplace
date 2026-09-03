# Keeping The Optimizer Honest

## Intuition

The compiler removes computations whose results nothing reads. Writing the result to a package-level variable makes it read-able from anywhere, so the work must survive.

## Approach

1. Sum with a plain loop, guarded for non-positive input.
2. Save the old `Sink`, overwrite it, return the saved value.

## Solution

```go
func Consume(v int) int {
	prev := Sink
	Sink = v
	return prev
}

func SumTo(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}
```

## Walkthrough

`Sink = v` before `prev := Sink` would return the value just stored, which is why the read comes first.

## Pitfalls

- A function-local sink, which the optimizer can still discard.
- `_ = result`, which is explicitly a no-op and provides no protection.
- A sink whose store is expensive (a map insert, a channel send), which is then what you are actually measuring.
