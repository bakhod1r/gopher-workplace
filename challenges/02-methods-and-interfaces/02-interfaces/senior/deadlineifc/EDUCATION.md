# Deadline Guard

## Intuition

A deadline is only respected if it is consulted at a point where stopping is cheap. Checking before each unit of work means the reported count is always exactly what completed.

## Approach

1. `CountingOp.Do` increments `Runs` and advances the injected clock by `Cost`.
2. `RunUntil` loops while `clock.Now().Before(deadline)`.
3. Each iteration runs the operation and counts it.
4. A deadline already reached produces zero iterations.

## Solution

```go
func (c *CountingOp) Do() {
	c.Runs++
	if c.Clock != nil {
		c.Clock.Advance(c.Cost)
	}
}

func RunUntil(clock Clock, deadline time.Time, op Op) int {
	done := 0
	for clock.Now().Before(deadline) {
		op.Do()
		done++
	}
	return done
}
```

## Walkthrough

With a 2-second cost and a 5-second budget: runs start at t=0, 2, 4 (all before the deadline) and the clock lands on 6, where the check fails. Three operations, and the overshoot is one operation's cost — inherent to any pre-check scheme.

## Pitfalls

- Checking the deadline after the run, which can report work that started too late.
- `!Now().After(deadline)`, which lets a run start exactly at the deadline.
- Using `time.Now()` instead of the injected clock, making the test non-deterministic.
