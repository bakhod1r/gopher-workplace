# Canonical keys for grouping

## The idea

Two anagrams share the same multiset of letters, so their **sorted** letters form
a canonical key. Count distinct keys with a set:

```go
b := []byte(word)
sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
key := string(b)
set[key] = struct{}{}
```

## Why it matters

"Normalize to a canonical form, then group/compare" is a widely reusable pattern
(dedup by content hash, case-fold grouping, fingerprinting).

## Watch out

- Convert the sorted `[]byte` back to a `string` to use as a map key.
- This keys by bytes; for full Unicode, normalize runes.
- Sorting each word is O(k log k); a letter-count key is O(k).
