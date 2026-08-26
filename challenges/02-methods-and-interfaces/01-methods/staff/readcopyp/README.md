# Read Copy Update

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Update` in [readcopyp.go](readcopyp.go):

1. `r.mu.Lock()`, defer `r.mu.Unlock()`.
2. Get the old `Config` (though you don't strictly need it to make a new one, this is the pattern).
3. `r.ptr.Store(&Config{Data: newData})`.

## Validate

```bash
make verify
```
