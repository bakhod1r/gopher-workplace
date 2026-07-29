# Counting runes, not bytes

## Intuition

`len(s)` is the **byte** count. The number of characters is the rune count:

```go
utf8.RuneCountInString(s) // characters
len(s)                    // bytes (>= rune count)
```

Ranging over a string also yields runes, so counting iterations works too.

## Approach

1. Convert the string to []rune, which decodes UTF-8.
2. Return len of that rune slice (a range-count loop works equally).

## Solution

```go
func Count(s string) int {
	return len([]rune(s))
}
```

## Walkthrough

Count of a 5-character accented word: []rune decodes to 5 code points, length 5.

## Pitfalls

- `len(s)` equals the rune count only for pure ASCII.
- `for range s` decodes UTF-8; the index is a byte offset that jumps by rune
  width.
- A rune is a code point, not necessarily a full grapheme (combining marks).
