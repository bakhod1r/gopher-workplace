# Rate Limit

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A basic rate limiter uses methods to check and consume tokens safely across
multiple goroutines.

## Task

Implement `Allow` and `Refill` on `*Limiter` in [ratelimit.go](ratelimit.go):

1. `Allow`: if `tokens > 0`, decrement and return `true`; else `false`.
2. `Refill`: add `n` to `tokens`.
3. Both must be thread-safe using `l.mu`.

## Validate

```bash
make verify
```
