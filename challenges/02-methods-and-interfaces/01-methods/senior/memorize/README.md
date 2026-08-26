# Memorize

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Get` in [memorize.go](memorize.go):

1. Check if `key` is in `m.cache`. If so, return it.
2. Otherwise, call `m.fn(key)`, store it in `cache`, and return it.

## Validate

```bash
make verify
```
