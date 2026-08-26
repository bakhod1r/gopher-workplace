# HyperLogLog Simulation

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Add` in [hyperlogl.go](hyperlogl.go):

1. Compute `zeros = leadingZeros(hash)`.
2. Update `h.maxZeros` if `zeros` is larger.

## Validate

```bash
make verify
```
