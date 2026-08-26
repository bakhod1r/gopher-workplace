# Exponential Backoff

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Next` in [backoff.go](backoff.go):

1. Save the current delay to return it.
2. Double `current`.
3. If `current > max`, set `current = max`.
4. Return the saved delay.

## Validate

```bash
make verify
```
