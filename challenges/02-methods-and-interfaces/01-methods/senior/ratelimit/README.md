# Token Bucket

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A token bucket meters work: each admitted call consumes a token, and a refill
puts tokens back. The count is shared state, so both methods must hold the
mutex — a check followed by an unguarded decrement is a classic race.

## Task

Implement `Allow` and `Refill` on `*Limiter` in [ratelimit.go](ratelimit.go):

1. `Allow`: under `l.mu`, if `l.tokens > 0` decrement and return `true`;
   otherwise return `false`.
2. `Refill(n)`: under `l.mu`, add `n` to `l.tokens`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  NewLimiter(2); Allow(), Allow()
Output: true, true
```

**Example 2:**

```
Input:  a third Allow()
Output: false (bucket empty)
```

**Example 3:**

```
Input:  Refill(1); Allow(); Allow()
Output: true, false
```

_Explanation:_ a refill adds exactly the tokens it was given.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Check-and-decrement is one operation** | Both halves must happen under the same lock hold. |
| 2 | **`defer` unlock** | Releases on every return path, including the `false` branch. |
| 3 | **Pointer receiver** | A `sync.Mutex` must not be copied. |

## Hint

Do not decrement before checking — an empty bucket would go negative and then
never admit anything again after a refill.

## Validate

```bash
make verify
```
