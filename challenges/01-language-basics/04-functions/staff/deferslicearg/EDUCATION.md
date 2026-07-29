# Slice headers as deferred arguments

## Intuition

A deferred argument snapshots the slice header (including length) at defer-time; since append reassigns the variable, only a body capture sees the final slice.

## Approach

1. A deferred argument snapshots the slice header at `defer` time.
2. If `xs` grows afterward, the snapshot's length is stale.
3. Close over `xs` instead of passing it.

## Solution

```go
func BuildAndReport(n int) (reported int) {
	var xs []int
	defer func() { reported = len(xs) }()
	for i := 0; i < n; i++ {
		xs = append(xs, i)
	}
	return
}
```

## Walkthrough

`defer f(xs)` captures the length before the builds finish. A parameterless defer reading `len(xs)` at return sees the final length.

## Pitfalls

- `defer f(xs)` freezes xs's header (len 0) before the appends.
- Reference `xs` in the closure body for the final value.
