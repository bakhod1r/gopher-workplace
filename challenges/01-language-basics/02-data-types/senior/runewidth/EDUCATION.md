# Advancing by rune width

## The idea

`utf8.DecodeRuneInString(s)` returns the first rune and the number of bytes it
occupies. A decoder advances by that `size`, not by the whole string length:

```go
_, size := utf8.DecodeRuneInString(s)
return size
```

## Why it matters

Manual UTF-8 iteration (parsers, tokenizers, streaming decoders) steps one rune
at a time. Returning `len(s)` collapses the loop to a single step and skips the
rest of the input — a real bug in hand-written decoders.

## Watch out

- `DecodeRuneInString` returns `(RuneError, 1)` for invalid bytes, so decoders
  still make progress.
- Width is 1 for ASCII, up to 4 for astral-plane runes.
- `for range` gives you this stepping for free; manual decode is for control.
