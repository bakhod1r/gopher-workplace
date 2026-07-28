# Positional base conversion

## The idea

To write `n` in base `b`, repeatedly take `n % b` (the next digit) and `n /= b`,
until `n` is 0. Digits emerge least-significant first, so reverse:

```go
const digits = "0123456789abcdef"
for n > 0 { buf = append(buf, digits[n%base]); n /= base }
// reverse buf
```

## Why it matters

This is how every integer-to-string formatter works (`strconv.FormatInt`
included). It ties together `%`, `/`, and indexing a digit table by value.

## Watch out

- Handle `n == 0` specially, or you return an empty string.
- Indexing `digits[n%base]` gives a `byte`; build a `[]byte` and convert once.
- This version assumes non-negative `n`; a sign needs separate handling.
