# Session Timeout

## Intuition

Expiry is a comparison between two moments and a budget. The type stores one
moment (`lastPing`) and the budget (`timeout`); the caller supplies the other
moment. That inversion — clock as a parameter, not a hidden dependency — is what
makes the logic testable at all.

## Approach

1. `Ping` overwrites the stored moment.
2. `IsExpired` subtracts to get an elapsed `Duration` and compares it to the budget.

## Solution

```go
func (s *Session) Ping(now time.Time) {
	s.lastPing = now
}

func (s *Session) IsExpired(now time.Time) bool {
	return now.Sub(s.lastPing) > s.timeout
}
```

## Walkthrough

The test starts at `time.Unix(0, 0)` with a 5s timeout.

- At `+4s`: elapsed 4s, not greater than 5s → not expired.
- At `+6s`: elapsed 6s → expired.
- `Ping(+4s)` moves `lastPing` to `+4s`; at `+6s` elapsed is only 2s → alive again.

## Pitfalls

- **`s.lastPing.Add(s.timeout).Before(now)`.** Equivalent, but longer, and it
  invites an off-by-one on the boundary.
- **Comparing `time.Time` with `<`.** Not allowed — `time.Time` is a struct.
  Use `Sub`, `Before` or `After`.
- **Calling `time.Now()` inside the methods.** The signature already gives you
  `now`; reading the real clock makes the third assertion flaky.
- **`>=` instead of `>`.** Flips the exact-boundary case.

## Monotonic time

A `time.Time` from `time.Now()` carries a monotonic reading, and `Sub` uses it —
so elapsed time stays correct across wall-clock adjustments. `time.Unix(0, 0)`
in the test has no monotonic part, which is fine: `Sub` falls back to wall time.
