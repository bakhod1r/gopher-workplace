# Modular letter shifting

## The idea

Map a letter to 0..25 by subtracting its case base, add the shift, take it mod 26
(normalized non-negative), and add the base back:

```go
base + rune(((int(r-base)+n)%26+26)%26)
```

Non-letters are left untouched.

## Why it matters

It combines rune ranges, character arithmetic, and the normalize-modulo idiom —
the same wraparound pattern as clock arithmetic, now over a 26-letter alphabet.

## Watch out

- Handle upper and lower case with their own base.
- Normalize the modulo (`(x%26+26)%26`) so negative shifts wrap correctly.
- ROT13 is `Shift(s,13)` and is its own inverse.
