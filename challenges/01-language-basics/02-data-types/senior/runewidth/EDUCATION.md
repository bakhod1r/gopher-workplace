# Advancing by rune width

## Intuition

`utf8.DecodeRuneInString(s)` returns the first rune and the number of bytes it
occupies. A decoder advances by that `size`, not by the whole string length:

```go
_, size := utf8.DecodeRuneInString(s)
return size
```

## Approach

1. Bug: returned `len(s)`, the whole string's byte length.
2. Fix: return `size` from utf8.DecodeRuneInString, the first rune's byte width.
3. Empty string returns 0 early.

## Solution

```go
import "unicode/utf8"

func FirstWidth(s string) int {
	if s == "" {
		return 0
	}
	_, size := utf8.DecodeRuneInString(s)
	_ = size
	return size
}
```

## Walkthrough

FirstWidth("日本"): DecodeRuneInString gives size=3 for 日 -> return 3.

## Pitfalls

- `DecodeRuneInString` returns `(RuneError, 1)` for invalid bytes, so decoders
  still make progress.
- Width is 1 for ASCII, up to 4 for astral-plane runes.
- `for range` gives you this stepping for free; manual decode is for control.
