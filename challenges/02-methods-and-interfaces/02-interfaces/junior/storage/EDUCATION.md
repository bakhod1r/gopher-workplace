# Storage

## Intuition

Two methods describe a whole storage backend. `Copy` works between any pair of stores — memory to memory, or memory to something else entirely.

## Approach

1. `Put` writes `m.data[key] = value`.
2. `Get` returns the comma-ok result directly.
3. `Copy` reads each key from `src`; on `ok`, writes it to `dst` and increments the counter.

## Solution

```go
func (m *MemStore) Put(key, value string) { m.data[key] = value }

func (m *MemStore) Get(key string) (string, bool) {
	v, ok := m.data[key]
	return v, ok
}

func Copy(src, dst Store, keys []string) int {
	n := 0
	for _, k := range keys {
		if v, ok := src.Get(k); ok {
			dst.Put(k, v)
			n++
		}
	}
	return n
}
```

## Walkthrough

`Copy(src, dst, []string{"a", "zz"})`: `a` exists, so it is written to `dst` and counted; `zz` is missing, so it is skipped. Result 1.

## Pitfalls

- Building `MemStore{}` directly — `data` is nil and `Put` panics.
- Returning `m.data[key], true` unconditionally, which hides missing keys.
- Counting every requested key instead of only the copied ones.
