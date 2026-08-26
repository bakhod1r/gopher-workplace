# State Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Publish` in [statepattern.go](statepattern.go):

1. If `d.State == Draft`, change to `Moderation`.
2. If `d.State == Moderation`, change to `Published`.

## Validate

```bash
make verify
```
