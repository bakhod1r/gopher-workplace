# Exponential Backoff

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A client retrying a failing dependency must wait longer after each attempt, but
never longer than a ceiling. `Backoff` keeps the current delay as state; each
call to `Next` hands out the current delay and prepares the following one.

## Task

Implement `Next` on `*Backoff` in [backoff.go](backoff.go):

1. Remember the current delay — it is the value to return.
2. Double `b.current` for the next call.
3. Clamp `b.current` to `b.max` so it never exceeds the ceiling.
4. Return the remembered delay.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  New(5*time.Second); Next()
Output: 1s
```

**Example 2:**

```
Input:  third Next() on the same Backoff
Output: 4s
```

**Example 3:**

```
Input:  fourth and fifth Next() with max = 5s
Output: 5s, 5s
```

_Explanation:_ 8s would exceed the ceiling, so the state saturates at `max`.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer receiver mutation** | `Next` must be on `*Backoff`; a value receiver would double a copy. |
| 2 | **Return-then-advance** | Capture the value to return *before* mutating the field. |
| 3 | **`time.Duration` arithmetic** | A `Duration` is an int64 of nanoseconds — `b.current * 2` is ordinary multiplication. |

## Hint

Save `d := b.current` first. If you return `b.current` after doubling, every
delay is off by one step.

## Validate

```bash
make verify
```
