# ARC Cache Simulation

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Access` in [arcpool.go](arcpool.go):

1. If `key` is in `T2`, do nothing.
2. If `key` is in `T1`, remove from `T1` and add to `T2`.
3. If not in either, add to `T1`.

## Validate

```bash
make verify
```
