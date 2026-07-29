# When defer evaluates its arguments

## Intuition

The deferred function's arguments are evaluated when `defer` executes; only the call is postponed.

## Approach

1. A deferred argument is evaluated at the `defer` statement.
2. `defer f(x)` with `x == 1` snapshots 1 before `x = 100`.

## Solution

```go
func Snapshot() (r int) {
	x := 1
	defer func(v int) { r = v }(x)
	x = 100
	return
}
```

## Walkthrough

`x` is 1 when the defer registers, so the snapshot is 1; changing `x` to 100 afterward does not affect it.

## Pitfalls

- `defer f(x)` snapshots `x`; `defer func(){ use(x) }()` reads it at return.
- Here the argument form gives 1.
