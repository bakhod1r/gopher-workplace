# Duplicate Keys In The Order List

## Intuition

Appending unconditionally records every write, so an updated key appears once per assignment while the map still holds one entry.

## Approach

1. Look the key up first.
2. Append it to the order list only when it is absent.
3. Then assign the value.

## Solution

```go
func (o *Ordered[K, V]) Set(k K, v V) {
	if _, ok := o.items[k]; !ok {
		o.keys = append(o.keys, k)
	}
	o.items[k] = v
}

func (o *Ordered[K, V]) Get(k K) (V, bool) {
	v, ok := o.items[k]
	return v, ok
}

func (o *Ordered[K, V]) Keys() []K {
	out := make([]K, len(o.keys))
	copy(out, o.keys)
	return out
}

func NewOrdered[K comparable, V any]() *Ordered[K, V] {
	return &Ordered[K, V]{items: make(map[K]V), keys: make([]K, 0)}
}
```

## Walkthrough

`Set(a,1); Set(a,2)` leaves `keys` as `[a a]` while `items` has one entry — the renderer prints the field twice.

## Pitfalls

- Appending on every write.
- Deduplicating in `Keys` instead of maintaining the invariant.
- Rebuilding the key list from the map, which loses insertion order.
