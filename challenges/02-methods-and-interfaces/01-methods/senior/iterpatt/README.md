# Iterator Pattern (Linked List)

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `HasNext` and `Next` in [iterpatt.go](iterpatt.go):

1. `HasNext`: Return `true` if `current` is not nil.
2. `Next`: Get `current.Val`, set `current = current.Next`, return the value.

## Validate

```bash
make verify
```
