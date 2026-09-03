# A Sink That Sinks Nothing

## Intuition

`_ = v` says "I know about this value and I am choosing to drop it" — which is the exact opposite of what a benchmark sink is for.

## Approach

1. Assign the argument to the package-level variable.

## Solution

```go
func Consume(v int) int {
	prev := Sink
	Sink = v
	return prev
}
```

## Walkthrough

With `_ = v`, nothing outside the function ever observes the computed value, so the compiler is free to delete the call chain that produced it — and a benchmark of an empty loop is very fast indeed.

## Pitfalls

- `_ = result` used as a sink; it is the canonical *non*-solution.
- A function-local sink, which the optimiser can still see is unused.
- A sink whose store is expensive, so the benchmark measures the sink.
