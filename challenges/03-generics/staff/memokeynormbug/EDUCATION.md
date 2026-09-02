# The Cache Whose Keys Never Match

## Intuition

The map does its job perfectly: two structs whose `Host` fields differ by a single capital letter are different keys. Without the normalisation step, each spelling gets its own entry and `fn` runs again for every one.

## Approach

1. Lazily create the map.
2. Normalise the key before touching the map.
3. Return the hit, or compute, store and return.

## Solution

```go
func (m *Memo[V]) Get(k Key, fn func(Key) V) V {
	if m.m == nil {
		m.m = make(map[Key]V)
	}
	k = Norm(k)
	if v, ok := m.m[k]; ok {
		return v
	}
	v := fn(k)
	m.m[k] = v
	return v
}

func Norm(k Key) Key {
	k.Host = strings.ToLower(strings.TrimSpace(k.Host))
	return k
}
```

## Walkthrough

Alternating `API.example.com` and `api.example.com` across a hundred thousand lookups calls `fn` a hundred thousand times instead of once.

## Pitfalls

- Normalising on write but not on read (or the reverse), which halves the problem instead of fixing it.
- Normalising after the lookup, so the stored key is canonical but the probe is not.
