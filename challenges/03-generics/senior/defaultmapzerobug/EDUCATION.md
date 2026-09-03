# The Lookup That Grows The Map

## Intuition

Materialising the default on read makes every queried key permanent, so `Len` measures traffic rather than configuration.

## Approach

1. Look the key up with the comma-ok form.
2. Return the stored value on a hit.
3. Return the default on a miss, storing nothing.

## Solution

```go
func (d *DefaultMap[K, V]) Get(k K) V {
	if v, ok := d.m[k]; ok {
		return v
	}
	return d.Default
}

func (d *DefaultMap[K, V]) Set(k K, v V) {
	if d.m == nil {
		d.m = make(map[K]V)
	}
	d.m[k] = v
}

func (d *DefaultMap[K, V]) Len() int {
	return len(d.m)
}
```

## Walkthrough

One `Get` of an unknown flag leaves `Len() == 1` even though nothing was configured.

## Pitfalls

- Auto-vivification copied from other languages.
- Fixing it by clearing the map afterwards instead of never writing.
