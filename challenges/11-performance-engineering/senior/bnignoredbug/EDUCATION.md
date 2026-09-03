# The Body That Runs Once

## Intuition

The contract between the harness and the body is one number. The harness says "run this `N` times"; the body's loop is the only thing that honours it.

## Approach

1. Bound the loop by `n` instead of by the constant.

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

With the bug, `PerOp(300, 3, noop)` reports 300 ns/op instead of 100: the elapsed time covers three iterations' worth of work, and the divisor is one.

## Pitfalls

- `for i := 0; i < 1; i++` left behind from debugging a single iteration.
- `for i := range n` in newer Go, which is correct — but only if you remember to change the bound at all.
- Reporting a benchmark whose ns/op is suspiciously stable across `-benchtime` values; that is the symptom.
