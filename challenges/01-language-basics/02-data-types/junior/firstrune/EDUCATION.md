# Bytes vs runes

## Intuition

A Go string is a read-only sequence of **bytes** holding UTF-8. `s[i]` indexes a
byte. A **rune** is a Unicode code point (an `int32`), and one rune may occupy
1–4 bytes. Ranging over a string decodes UTF-8 and yields runes:

```go
for _, r := range s { return r } // first rune, whole character
```

## Approach

1. Range over the string, which decodes UTF-8 into runes.
2. Return the rune on the first iteration.
3. If the loop never runs (empty string), return 0.

## Solution

```go
func First(s string) rune {
	for _, r := range s {
		return r
	}
	return 0
}
```

## Walkthrough

First("hello"): first range iteration yields r='h' at byte 0, return it immediately.

## Pitfalls

- The range index is a **byte offset**, so it jumps by rune width (0, then the
  byte length of each rune).
- `len(s)` is the byte count, not the rune count; use `utf8.RuneCountInString`.
- Converting `[]rune(s)` gives random access to runes at the cost of a copy.
