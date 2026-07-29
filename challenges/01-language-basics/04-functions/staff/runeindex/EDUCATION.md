# Indexing strings by rune

## Intuition

String indexing is byte-based; to address the i-th character convert to `[]rune` (O(n)) or decode sequentially.

## Approach

1. `s[i]` indexes **bytes**, breaking on multi-byte runes.
2. Convert to `[]rune` and index that.

## Solution

```go
func CharAt(s string, i int) rune {
	return []rune(s)[i]
}
```

## Walkthrough

`"héllo"[2]` returns a byte in the middle of the multi-byte string, not the 3rd character. `[]rune(s)[2]` decodes and returns `'l'`.

## Pitfalls

- `s[i]` yields a byte; `[]rune(s)[i]` yields the i-th character.
- `[]rune(s)` costs a pass and allocation; fine for random access.
