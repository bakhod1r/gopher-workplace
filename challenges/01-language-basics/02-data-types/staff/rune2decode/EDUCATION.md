# Assembling a code point from UTF-8

## Intuition

A 2-byte UTF-8 rune splits its bits across the two bytes:

```
lead: 110xxxxx  (5 payload bits -> mask 0x1F)
cont: 10yyyyyy  (6 payload bits -> mask 0x3F)
rune = (x << 6) | y
```

## Approach

1. Bug: lead masked with 0x0F keeps only 4 bits; 2-byte lead payload is 5 bits.
2. Lead is 110xxxxx -> low 5 bits (0x1F), shifted left 6, OR cont low 6 bits.
3. Fix: rune(lead&0x1F)<<6 | rune(cont&0x3F).

## Solution

```go
func Decode2(lead, cont byte) rune {
	return rune(lead&0x1F)<<6 | rune(cont&0x3F)
}
```

## Walkthrough

0xD0&0x1F=0x10, <<6=0x400; 0x81&0x3F=0x01 -> 0x401 = U+0401.

## Pitfalls

- 2-byte lead → 5 bits (`0x1F`); 3-byte lead → 4 bits (`0x0F`); the mask depends
  on the length.
- Continuation bytes always contribute 6 bits (`0x3F`).
- Shift the lead payload left by 6 (the width of one continuation).
