# Canonical keys for grouping

## Intuition

Two anagrams share the same multiset of letters, so their **sorted** letters form
a canonical key. Count distinct keys with a set:

```go
b := []byte(word)
sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
key := string(b)
set[key] = struct{}{}
```

## Approach

1. For each word, sort its bytes to form a canonical key.
2. Insert the key into a set (map[string]bool).
3. Distinct anagram groups = number of keys in the set.
4. Return len(set).

## Solution

```go
import "sort"

func Count(words []string) int {
	seen := map[string]bool{}
	for _, w := range words {
		bs := []byte(w)
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
		seen[string(bs)] = true
	}
	return len(seen)
}
```

## Walkthrough

"eat" and "tea" both sort to "aet" -> one key; "tan","nat" -> "ant"; "bat" -> "abt". Set = {aet,ant,abt}, len 3.

## Pitfalls

- Convert the sorted `[]byte` back to a `string` to use as a map key.
- This keys by bytes; for full Unicode, normalize runes.
- Sorting each word is O(k log k); a letter-count key is O(k).
