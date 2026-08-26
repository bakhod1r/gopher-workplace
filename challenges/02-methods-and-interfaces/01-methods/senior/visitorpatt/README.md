# Visitor Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Accept` in [visitorpatt.go](visitorpatt.go):

1. Base case: if `n == nil`, return.
2. Call `visitor(n.Val)`.
3. Call `Accept(visitor)` on `Left` and `Right`.

## Validate

```bash
make verify
```
