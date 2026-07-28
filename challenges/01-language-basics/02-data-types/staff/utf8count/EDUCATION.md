# Counting runes without decoding

## The idea

Every rune has exactly one lead byte; continuation bytes are `10xxxxxx`
(`c&0xC0 == 0x80`). So the rune count equals the number of **non-continuation**
bytes:

```go
if c&0xC0 != 0x80 { n++ }
```

## Why it matters

This is how `utf8.RuneCount` works — O(n) time, zero allocation, no decoding. The
inverted test counts exactly the wrong bytes, so the result is nonsense for any
real text.

## Watch out

- ASCII bytes (`0xxxxxxx`) are non-continuation and count as one rune each.
- This assumes valid UTF-8; invalid bytes need a real decoder.
- `c&0xC0` isolates the top two bits — the continuation signature is `10`.
