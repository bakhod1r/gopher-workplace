# Snapshot vs live capture in defer

## Intuition

Passing a value as a deferred argument freezes it at defer-time; referencing it in the closure body reads it at return-time.

## Approach

1. Deferred **arguments** are evaluated at the `defer` statement, snapshotting the value.
2. The bug captures `c` early, before later mutations.
3. Close over `c` instead so the final value is read at return.

## Solution

```go
func FinalCount(n int) (recorded int) {
	c := 0
	defer func() { recorded = c }()
	for i := 0; i < n; i++ {
		c++
	}
	return
}
```

## Walkthrough

`defer f(c)` freezes `c` at its value when defer runs. Closing over `c` in `defer func(){ recorded = c }()` reads it after all updates, so `FinalCount(5)` returns 5.

## Pitfalls

- `defer f(c)` uses c's value NOW; `defer func(){ use(c) }()` uses it at return.
- For final values, capture in the body.
