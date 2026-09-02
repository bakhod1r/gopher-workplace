# Rate Limiter

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An outbound client must not exceed a burst budget. The limiter is an interface so tests can drive time deterministically.

## Task

Implement the stub(s) in [ratelimitifc.go](ratelimitifc.go):

1. Implement `Allow` on `*TokenBucket` — consume a token when available, refilling by elapsed time from the injected `Clock`.
2. Implement `AllowN`, which reports how many of n attempts were allowed.
3. Constraint: no wall-clock sleeping in tests — the clock is injected, and the limiter must never allow more than `Burst` in a window.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Burst 2, no time passing; three Allow calls
Output: true, true, false
```

**Example 2:**

```
Input:  advance the clock by the refill interval
Output: one more token
```

**Example 3:**

```
Input:  AllowN(5) with Burst 2
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Clock injection** | Time behind an interface makes rate limits testable. |
| 2 | **Token bucket** | Capacity is the burst; refill rate is the steady-state limit. |
| 3 | **Clamping** | Refill must not exceed the burst ceiling. |

## Hint

Refill = elapsed / interval, capped at `Burst`. Advance `last` by the tokens actually granted.

## Validate

```bash
make verify
```
