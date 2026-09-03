# The Delete That Half-Unlinks

## Intuition

In a doubly linked list every node is referenced twice: by its predecessor's `next` and by its successor's `prev` (or by the list's `head`/`tail`). Repairing only the forward chain leaves stale back-pointers, and a stale `tail` means the next insertion is appended to a node that is no longer in the list.

## Approach

1. Look the node up and drop it from the index.
2. Repair the forward link: the predecessor's `next`, or `head`.
3. Repair the backward link: the successor's `prev`, or `tail`.
4. Clear the removed node's own pointers so it cannot be walked back into.

## Solution

```go
func (m *OrderedMap[K, V]) Delete(k K) bool {
	n, ok := m.idx[k]
	if !ok {
		return false
	}
	delete(m.idx, k)
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		m.head = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		m.tail = n.prev
	}
	n.prev, n.next = nil, nil
	return true
}

func (m *OrderedMap[K, V]) Set(k K, v V) {
	if m.idx == nil {
		m.idx = make(map[K]*onode[K, V])
	}
	if n, ok := m.idx[k]; ok {
		n.val = v
		return
	}
	n := &onode[K, V]{key: k, val: v, prev: m.tail}
	if m.tail != nil {
		m.tail.next = n
	} else {
		m.head = n
	}
	m.tail = n
	m.idx[k] = n
}

func (m *OrderedMap[K, V]) Get(k K) (V, bool) {
	if n, ok := m.idx[k]; ok {
		return n.val, true
	}
	var zero V
	return zero, false
}

func (m *OrderedMap[K, V]) Len() int {
	return len(m.idx)
}

func (m *OrderedMap[K, V]) Keys() []K {
	out := make([]K, 0, len(m.idx))
	for n := m.head; n != nil; n = n.next {
		out = append(out, n.key)
	}
	return out
}

func (m *OrderedMap[K, V]) RevKeys() []K {
	out := make([]K, 0, len(m.idx))
	for n := m.tail; n != nil; n = n.prev {
		out = append(out, n.key)
	}
	return out
}
```

## Walkthrough

`Set a,b; Delete b` leaves `tail` pointing at the removed `b`. `Set c` then hangs `c` off `b`, so `Keys()` returns `[a]` while `Len()` says 2 — the entry is in the index but unreachable from the list.

## Pitfalls

- Fixing `head` but not `tail`, which only shows when the last entry is removed.
- Leaving the removed node's pointers set, so a stale reference can walk into live data.
- Testing deletes only from the middle, where the forward chain alone looks correct.
