# Timer Reset

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Ping` and `IsExpired` in [timerreset.go](timerreset.go):

1. `Ping`: update `lastPing` to `now`.
2. `IsExpired`: return `now.Sub(s.lastPing) > s.timeout`.

## Validate

```bash
make verify
```
