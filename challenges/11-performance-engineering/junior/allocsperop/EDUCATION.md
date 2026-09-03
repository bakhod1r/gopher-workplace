# Counting Allocations Without A Profiler

## Intuition

`testing.AllocsPerRun` brackets `runs` calls of `f` with the runtime's allocation counter and divides. It is the measurement a "must not allocate" test should assert on.

## Approach

1. Clamp `runs` to at least 1.
2. Call `testing.AllocsPerRun` and round the result.

## Solution

```go
func AllocsOf(runs int, f func()) int {
	if runs < 1 {
		runs = 1
	}
	return int(math.Round(testing.AllocsPerRun(runs, f)))
}
```

## Walkthrough

`AllocsPerRun` also does one warm-up call, so lazily initialised state inside `f` is not counted as a per-call allocation.

## Pitfalls

- `int(...)` alone, which truncates `0.999999` to `0`.
- Passing `runs = 0`, which divides by zero inside the helper.
- Asserting on allocated *bytes* instead of count; byte totals vary with size classes.
