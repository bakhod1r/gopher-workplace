# Retry Logic

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A resilient client automatically retries transient failures.

## Task

Implement `DoWithRetry` on `*Client` in [retrylogic.go](retrylogic.go):

1. Loop up to `maxAttempts` times.
2. Call `c.Do()`. If it returns `nil`, return `nil` immediately.
3. If the loop finishes without success, return the last error.

## Validate

```bash
make verify
```
