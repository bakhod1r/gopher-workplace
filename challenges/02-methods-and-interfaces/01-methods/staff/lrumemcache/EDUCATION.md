# LRU Cache

## Intuition

Two data structures, one invariant: the map and the list always describe the
same set of entries. The map answers "where is this key"; the list answers "what
should die next". Every operation must keep both in step — the classic bug is
evicting from the list and forgetting the map, which leaks memory and, worse,
leaves the map pointing at an unlinked node.

## Approach

1. `Get`: look up, and on a hit re-insert at the front to record the access.
2. `Put`: update-in-place if present, otherwise insert new.
3. After a new insert, check the size and evict from the tail if over capacity.

## Solution

```go
func (l *LRU) Get(key string) (int, bool) {
	n, ok := l.cache[key]
	if !ok {
		return 0, false
	}
	l.remove(n)
	l.insert(n)
	return n.val, true
}

func (l *LRU) Put(key string, val int) {
	if n, ok := l.cache[key]; ok {
		n.val = val
		l.remove(n)
		l.insert(n)
		return
	}

	n := &node{key: key, val: val}
	l.insert(n)
	l.cache[key] = n

	if len(l.cache) > l.cap {
		lru := l.tail.prev
		l.remove(lru)
		delete(l.cache, lru.key)
	}
}
```

## Walkthrough

With `cap == 2`:

| step | list (front → back) | map |
|------|---------------------|-----|
| `Put("a",1)` | a | a |
| `Put("b",2)` | b, a | a, b |
| `Get("a")` | a, b | a, b |
| `Put("c",3)` | c, a | a, c |

`Put("c",3)` inserts first, making the map size 3, so the tail neighbour — `b`,
untouched since it was inserted — is unlinked and deleted. The `Get("a")` is
what reordered the list; without it `a` would have been the victim.

This is exactly why the node stores its own `key`: eviction starts from the
*list* and needs the key to reach back into the map.

## Pitfalls

- **Evicting from the list without deleting from the map.** The next `Get` for
  that key returns a node that is no longer in the list; the two structures have
  diverged and the cache silently corrupts.
- **`Get` that does not reorder.** It becomes a FIFO cache — still functional,
  still wrong, and the test's `Get("a")` step is designed to catch it.
- **Checking capacity before inserting.** Off by one: the cache holds `cap-1`
  entries.
- **Evicting on an update.** An overwrite does not grow the map, so the guard
  must be tied to insertion.
- **Evicting `l.tail` itself.** That is the sentinel; the victim is `l.tail.prev`.

## Why sentinels

Without dummy head and tail nodes, `remove` and `insert` each need branches for
"is this the first node" and "is this the last". The sentinels make every real
node have a non-nil `prev` and `next` forever, which collapses four cases into
the two-line helpers this puzzle hands you.
