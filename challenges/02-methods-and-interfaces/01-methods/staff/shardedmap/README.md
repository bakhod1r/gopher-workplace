# Sharded Map

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Set` in [shardedmap.go](shardedmap.go):

1. Call `getShard(key)`.
2. Lock the shard's mutex.
3. Update the shard's `data`.
4. Unlock.

## Validate

```bash
make verify
```
