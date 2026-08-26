# LRU Cache

## Solution

```go
func (l *LRU) Get(key string) (int, bool) {
	if n, ok := l.cache[key]; ok {
		l.remove(n)
		l.insert(n)
		return n.val, true
	}
	return 0, false
}

func (l *LRU) Put(key string, val int) {
	if n, ok := l.cache[key]; ok {
		n.val = val
		l.remove(n)
		l.insert(n)
		return
	}
	n := &node{key: key, val: val}
	l.cache[key] = n
	l.insert(n)
	if len(l.cache) > l.cap {
		evict := l.tail.prev
		l.remove(evict)
		delete(l.cache, evict.key)
	}
}
```
