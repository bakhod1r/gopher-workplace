# Session Timeout

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A session dies if it goes quiet for too long. Rather than reading the clock
inside the type — which would make it untestable — both methods take `now`
explicitly, so tests can drive time by hand.

## Task

Implement `Ping` and `IsExpired` on `*Session` in [timerreset.go](timerreset.go):

1. `Ping(now)` sets `s.lastPing` to `now`.
2. `IsExpired(now)` reports whether more than `s.timeout` has elapsed since `s.lastPing`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  lastPing = t0, timeout = 5s; IsExpired(t0+4s)
Output: false
```

**Example 2:**

```
Input:  IsExpired(t0+6s)
Output: true
```

**Example 3:**

```
Input:  Ping(t0+4s) then IsExpired(t0+6s)
Output: false   (only 2s since the ping)
```

_Explanation:_ pinging moves the deadline forward.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`Time.Sub`** | `now.Sub(s.lastPing)` yields a `time.Duration` you can compare directly. |
| 2 | **Injected clock** | Passing `now` in keeps the method deterministic — no sleeping in tests. |
| 3 | **Strict vs inclusive** | The contract says *more than* the timeout, so use `>`, not `>=`. |

## Hint

`return now.Sub(s.lastPing) > s.timeout`. Do not compare `time.Time` values with
`>` — that does not compile; subtract first.

## Validate

```bash
make verify
```
