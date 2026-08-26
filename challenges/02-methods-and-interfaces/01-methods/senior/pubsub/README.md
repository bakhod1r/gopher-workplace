# Pub Sub

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Publish` in [pubsub.go](pubsub.go):

1. Read-lock the mutex (`ps.mu.RLock()`).
2. Iterate over the channels for `topic`.
3. Send `msg` to each channel.
4. Unlock.

## Validate

```bash
make verify
```
