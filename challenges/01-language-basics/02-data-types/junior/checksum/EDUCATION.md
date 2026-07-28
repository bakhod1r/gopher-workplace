# Sums and modular reduction

## The idea

A simple checksum accumulates byte (or digit) values and reduces the total with
`%` to a fixed range:

```go
sum := 0
for _, b := range data { sum += int(b) }
return sum % 256
```

## Why it matters

Checksums detect accidental corruption cheaply. The building blocks — iterate,
accumulate, reduce mod N — recur in hashing and error detection.

## Watch out

- The modulus defines the checksum width; pick it deliberately.
- Byte order and starting value matter for cross-implementation agreement.
- A plain sum catches single-byte errors but not reordering (that needs a
  position-weighted scheme like Luhn or Fletcher).
