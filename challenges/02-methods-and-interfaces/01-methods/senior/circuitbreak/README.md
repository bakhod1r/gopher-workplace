# Circuit Breaker

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A circuit breaker prevents a system from repeatedly trying an operation that is
likely to fail.

## Task

Implement `Call` on `*Breaker` in [circuitbreak.go](circuitbreak.go):

1. If `IsOpen`, return `errors.New("circuit open")`.
2. Execute `fn()`.
3. If `fn()` returns an error, increment `ConsecutiveFails`. If it hits `Threshold`, set `IsOpen = true`.
4. If `fn()` succeeds, reset `ConsecutiveFails = 0`.
5. Return `fn()`'s error (if any).

## Validate

```bash
make verify
```
