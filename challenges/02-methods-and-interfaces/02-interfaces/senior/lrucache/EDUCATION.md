# LRU Cache

## Intuition

The map answers "where is this entry?" in O(1); the linked list answers "which entry is coldest?" in O(1). Neither structure alone gives both, which is why LRU pairs them.

## Approach

1. `Get` looks the node up, unlinks it, pushes it to the front, and returns its value.
2. `Put` on an existing key updates the value and touches the node.
3. `Put` on a new key evicts `tail.prev` when at capacity, then inserts at the front.
4. Both paths keep the map and the list in sync; `Len` reads the map.

## Solution

```go
func (l *LRU) Get(key string) (string, bool) {
	n, ok := l.items[key]
	if !ok {
		return "", false
	}
	l.unlink(n)
	l.pushFront(n)
	return n.value, true
}

func (l *LRU) Put(key, value string) {
	if l.Cap <= 0 {
		return
	}
	if n, ok := l.items[key]; ok {
		n.value = value
		l.unlink(n)
		l.pushFront(n)
		return
	}
	if len(l.items) >= l.Cap {
		lru := l.tail.prev
		l.unlink(lru)
		delete(l.items, lru.key)
	}
	n := &node{key: key, value: value}
	l.items[key] = n
	l.pushFront(n)
}

func (l *LRU) Len() int { return len(l.items) }
```

## Walkthrough

`Put(a), Put(b), Get(a), Put(c)`: the `Get` moves `a` to the front, so the tail now holds `b`. Inserting `c` evicts `b`, not `a`.

## Pitfalls

- Not touching on `Get`, which silently degrades to insertion-order eviction.
- Deleting the evicted node from the map by the wrong key — store `key` in the node so eviction can find it.
- Scanning `items` to find the oldest, which makes every `Put` O(n).
