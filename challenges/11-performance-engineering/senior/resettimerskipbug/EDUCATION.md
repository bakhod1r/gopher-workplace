# A Reset That Resets Nothing

## Intuition

The timer never stops accumulating; a reset only moves the line that measurement is counted from. Setting that line to zero puts it back at the beginning of time, which is the opposite of a reset.

## Approach

1. Move the mark to the current total, so `Elapsed` reports zero immediately after.

## Solution

```go
func (t *Timer) Reset() {
	t.mark = t.total
}
```

## Walkthrough

`Benchmark(1_000_000, 7, 1)` returns `1_000_007` with the bug and `7` without it — the setup dominates the entire measurement, and the reset that was supposed to prevent it is right there in the code.

## Pitfalls

- `t.mark = 0`, which reads like "clear the state" and actually un-resets the timer.
- Zeroing `total` instead of the mark, which breaks any code holding the lifetime total.
- Resetting a copy of the timer rather than the receiver — a value receiver here would fail the same way.
