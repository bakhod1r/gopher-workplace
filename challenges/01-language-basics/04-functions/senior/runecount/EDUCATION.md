# String length in bytes vs runes

## Intuition

A Go string is UTF-8 bytes; `len` returns bytes while ranging or `utf8.RuneCountInString` returns code points.

## Approach

1. `len(s)` counts **bytes**, not characters.
2. Range over the string (which yields runes) and count.

## Solution

```go
func CharCount(s string) int {
	n := 0
	for range s {
		n++
	}
	return n
}
```

## Walkthrough

`"héllo"` is 6 bytes but 5 runes; `len` returns 6. Ranging decodes UTF-8 and counts 5 runes.

## Pitfalls

- `len("é")` is 2, not 1.
- Range indices over a string are BYTE offsets, but the value is a rune.
