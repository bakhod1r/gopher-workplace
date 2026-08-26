# Lock-Free Stack

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Push` in [lockfreestk.go](lockfreestk.go):

1. Write a `for` loop (spin lock).
2. Load `old := s.head.Load()`.
3. Set `n.next = old`.
4. If `s.head.CompareAndSwap(old, n)`, `break`.

## Validate

```bash
make verify
```
