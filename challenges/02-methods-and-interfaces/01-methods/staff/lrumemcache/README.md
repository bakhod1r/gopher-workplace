# LRU Cache

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Task

Implement `Get` and `Put` in [lrumemcache.go](lrumemcache.go):

1. `Get`: Look up map. If found, call `remove(n)` and `insert(n)` (to move it to front). Return value.
2. `Put`: Look up map. If found, update `val`, `remove`, `insert`. If not found, create new node, `insert`, add to map. If `len(l.cache) > l.cap`, find `l.tail.prev`, `remove` it, and delete from map.

## Validate

```bash
make verify
```
