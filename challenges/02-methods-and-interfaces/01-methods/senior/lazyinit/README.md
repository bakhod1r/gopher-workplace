# Lazy Init

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `String` in [lazyinit.go](lazyinit.go):

1. If `l.val` is nil, call `l.init()`, store its address in `l.val`.
2. Return `*l.val`.

## Validate

```bash
make verify
```
