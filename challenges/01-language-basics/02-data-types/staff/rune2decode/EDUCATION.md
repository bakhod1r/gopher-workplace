# Assembling a code point from UTF-8

## The idea

A 2-byte UTF-8 rune splits its bits across the two bytes:

```
lead: 110xxxxx  (5 payload bits -> mask 0x1F)
cont: 10yyyyyy  (6 payload bits -> mask 0x3F)
rune = (x << 6) | y
```

## Why it matters

The mask must match the number of payload bits the encoding defines. `0x0F`
keeps only 4 bits and silently drops the 5th, corrupting every code point from
U+0100 upward — a decoder bug that passes for low-range accents and fails higher
up.

## Watch out

- 2-byte lead → 5 bits (`0x1F`); 3-byte lead → 4 bits (`0x0F`); the mask depends
  on the length.
- Continuation bytes always contribute 6 bits (`0x3F`).
- Shift the lead payload left by 6 (the width of one continuation).
