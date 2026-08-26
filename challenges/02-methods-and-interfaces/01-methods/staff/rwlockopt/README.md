# Lock Escalation

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `IncrementIfZero` in [rwlockopt.go](rwlockopt.go):

1. `RLock()`, check `v`. If `v != 0`, `RUnlock()` and return `v`.
2. `RUnlock()`.
3. `Lock()`, check `v == 0` again (double-checked locking). If so `v++`.
4. `Unlock()`, return `v`.

## Validate

```bash
make verify
```
