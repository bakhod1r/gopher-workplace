# Once Init

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Get` in [onceinit.go](onceinit.go):

1. Use `l.once.Do(...)` to call `l.init()` and set `l.data`.
2. Return `l.data`.

## Validate

```bash
make verify
```
