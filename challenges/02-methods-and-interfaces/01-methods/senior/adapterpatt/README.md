# Adapter Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `GetIntData` in [adapterpatt.go](adapterpatt.go):

1. Get string from `a.legacy.GetStringData()`.
2. Convert to `int` using `strconv.Atoi`. Return `0` on error.

## Validate

```bash
make verify
```
