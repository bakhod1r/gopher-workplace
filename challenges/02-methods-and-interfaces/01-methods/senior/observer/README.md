# Observer Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `SetState` in [observer.go](observer.go):

1. Set `s.state = val`.
2. Loop through `s.observers` and call each with `val`.

## Validate

```bash
make verify
```
