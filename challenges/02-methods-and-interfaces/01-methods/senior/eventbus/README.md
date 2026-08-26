# Event Bus

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `On` and `Emit` in [eventbus.go](eventbus.go):

1. `On`: Append the `listener` to the slice for `eventType` in the map.
2. `Emit`: Loop over all listeners for `eventType` and call them with `data`.

## Validate

```bash
make verify
```
