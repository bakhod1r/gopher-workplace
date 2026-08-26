# Actor Router

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Route` in [actorrouter.go](actorrouter.go):

1. Send `msg` to `r.workers[r.idx].Inbox`.
2. `r.idx = (r.idx + 1) % len(r.workers)`.

## Validate

```bash
make verify
```
