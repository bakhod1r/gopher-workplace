# Hazard Pointers

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Protect` in [hazardptr.go](hazardptr.go):

1. Load `p` from `shared`.
2. Store `p` in `h.ptr`.
3. Check if `shared.Load()` is still `p`. If so, return `p`. Else `nil`.

## Validate

```bash
make verify
```
