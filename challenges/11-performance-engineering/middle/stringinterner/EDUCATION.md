# One Copy Of Every Label

## Intuition

A map from string to string looks pointless until you notice the key and the value are the same *value* but only one *allocation* — and that everyone who interns gets that one back.

## Approach

1. Lazily create the map.
2. On a hit, count it and return the stored instance.
3. On a miss, store and return the string, counting the miss.

## Solution

```go
func (in *Interner) Intern(s string) string {
	if in.seen == nil {
		in.seen = make(map[string]string)
	}
	if got, ok := in.seen[s]; ok {
		in.hits++
		return got
	}
	in.misses++
	in.seen[s] = s
	return s
}

func (in *Interner) InternBytes(b []byte) string {
	if in.seen == nil {
		in.seen = make(map[string]string)
	}
	if got, ok := in.seen[string(b)]; ok {
		in.hits++
		return got
	}
	in.misses++
	s := string(b)
	in.seen[s] = s
	return s
}

func (in *Interner) Stats() (int, int) { return in.hits, in.misses }

func (in *Interner) Len() int { return len(in.seen) }
```

## Walkthrough

`in.seen[string(b)]` is a compiler-recognised pattern: the conversion is used only to hash and compare, so no string is allocated. The miss path converts for real, because that copy is the one the map will keep.

## Pitfalls

- Hoisting `s := string(b)` above the lookup, which allocates on every call and destroys the point.
- Storing a string that aliases the caller's buffer, so the map key mutates underneath the map.
- Interning unbounded input — an interner over user-supplied values is a memory leak with a nice name.
