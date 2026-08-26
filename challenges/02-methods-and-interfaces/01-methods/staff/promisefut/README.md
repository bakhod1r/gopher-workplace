# Future Pattern

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Complete` and `Get` in [promisefut.go](promisefut.go):

1. `Complete`: send `val` to `f.ch`, then `close(f.ch)`.
2. `Get`: return `<-f.ch`.

## Validate

```bash
make verify
```
