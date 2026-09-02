# Insertion-Ordered Map

## Intuition

The duplicate check is what keeps `len(keys)` equal to `len(items)`; without it, updates would emit the same key twice.

## Approach

1. `Set`: append the key when new, then assign.
2. `Get`: comma-ok lookup.
3. `Keys`: return a copy of the order list.

## Solution

```go
func NewOrdered[K comparable, V any]() *Ordered[K, V] {
	return &Ordered[K, V]{items: make(map[K]V), keys: make([]K, 0)}
}

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
```

## Walkthrough

`Set(a,1); Set(a,2)` updates the value in place, leaving `Keys()` as `[a]`.

## Pitfalls

- Appending on every `Set`, producing duplicate keys.
- Returning `o.keys` directly, letting callers reorder the structure.
- Relying on map range order instead of the key list.
