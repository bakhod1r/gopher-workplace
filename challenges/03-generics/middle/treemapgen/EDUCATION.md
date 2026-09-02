# Sorted Map

## Intuition

Maintaining order incrementally trades a linear shift per insert for a free `Keys()` — the right call when reads are frequent.

## Approach

1. `Set`: for a new key, find the first larger key and splice it in; then store the value.
2. `Get`: comma-ok lookup.
3. `Keys`: copy the list.

## Solution

```go
func NewSorted[K cmp.Ordered, V any]() *SortedMap[K, V] {
	return &SortedMap[K, V]{items: make(map[K]V), keys: make([]K, 0)}
}

func (m *SortedMap[K, V]) Set(k K, v V) {
	if _, ok := m.items[k]; !ok {
		i := 0
		for i < len(m.keys) && m.keys[i] < k {
			i++
		}
		m.keys = append(m.keys, k)
		copy(m.keys[i+1:], m.keys[i:])
		m.keys[i] = k
	}
	m.items[k] = v
}

func (m *SortedMap[K, V]) Get(k K) (V, bool) {
	v, ok := m.items[k]
	return v, ok
}

func (m *SortedMap[K, V]) Keys() []K {
	out := make([]K, len(m.keys))
	copy(out, m.keys)
	return out
}
```

## Walkthrough

`Set(b,1); Set(a,2)` inserts `a` before `b`, so `Keys()` reports `[a b]`.

## Pitfalls

- Appending and sorting on every `Set`.
- Inserting duplicate keys on update.
- Returning the internal key slice.
