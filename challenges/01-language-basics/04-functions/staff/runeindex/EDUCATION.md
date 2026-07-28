# Indexing strings by rune

## The idea

String indexing is byte-based; to address the i-th character convert to `[]rune` (O(n)) or decode sequentially.

## Why it matters

Byte-indexing multibyte text returns fragments of characters — a real internationalisation bug.

## Watch out

- `s[i]` yields a byte; `[]rune(s)[i]` yields the i-th character.
- `[]rune(s)` costs a pass and allocation; fine for random access.
