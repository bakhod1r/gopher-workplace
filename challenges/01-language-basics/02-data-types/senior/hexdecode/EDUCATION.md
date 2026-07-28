# Nibble ordering in hex

## The idea

Each byte is two hex digits: the **first** character is the high nibble (bits
4–7), the second the low nibble (bits 0–3):

```go
byte(hi<<4 | lo)
```

Swapping them (`lo<<4 | hi`) reverses each byte's nibbles — `"1a"` → `0xa1`.

## Why it matters

Hex/byte codecs sit under hashing, crypto, and binary protocols. A nibble-order
bug corrupts every byte yet still "decodes", so it passes shape checks and fails
only on values — exactly the kind of bug that ships.

## Watch out

- High nibble shifts left by 4; low nibble is unshifted.
- Reject odd-length input up front.
- Accept both letter cases when decoding.
