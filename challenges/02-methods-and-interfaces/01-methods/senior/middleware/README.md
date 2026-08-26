# Middleware Chain

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Chain` in [middleware.go](middleware.go):

1. Given `mws ...Middleware`, return a `Middleware` that wraps a `Handler`.
2. To make `mws[0]` the outermost wrapper, you must apply them in reverse order
   (from `len(mws)-1` down to `0`).
3. `next = mws[i](next)` inside the reverse loop.

## Validate

```bash
make verify
```
