# Pay For It Once

## Intuition

The map is derived data. Derive it the first time someone asks, store it on the struct, and every later call is a hash lookup.

## Approach

1. On a nil `byWord`, build it from `Words`, keeping the first index for duplicates.
2. Read the map and return the comma-ok result.

## Solution

```go
func (ix *Index) Lookup(w string) (int, bool) {
	if ix.byWord == nil {
		ix.byWord = make(map[string]int, len(ix.Words))
		for i, word := range ix.Words {
			if _, ok := ix.byWord[word]; !ok {
				ix.byWord[word] = i
			}
		}
	}
	i, ok := ix.byWord[w]
	return i, ok
}
```

## Walkthrough

The receiver must be `*Index`: with a value receiver the assignment to `ix.byWord` would land on a copy, the map would be rebuilt on every call, and the allocation test would fail.

## Pitfalls

- `len(ix.byWord) == 0` as the build check, which rebuilds forever when the word list is empty.
- Scanning `Words` linearly on every call — correct, but that is the hot spot you were removing.
- Mutating `Words` after the first lookup; the cached map goes stale.
