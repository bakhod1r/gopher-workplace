# UTF-8 lead-byte classification

## Intuition

UTF-8 encodes the byte count in the lead byte's high bits:

| bytes | lead pattern | mask test |
|-------|--------------|-----------|
| 1 | `0xxxxxxx` | `c < 0x80` |
| 2 | `110xxxxx` | `c&0xE0 == 0xC0` |
| 3 | `1110xxxx` | `c&0xF0 == 0xE0` |
| 4 | `11110xxx` | `c&0xF8 == 0xF0` |

Continuation bytes are `10xxxxxx` (`c&0xC0 == 0x80`).

## Approach

1. Bug: 2-byte lead test `c&0xC0==0xC0` also matches 3/4-byte leads, assigning them n=1.
2. A 2-byte lead is 110xxxxx -> mask 0xE0 must equal 0xC0.
3. Fix: case c&0xE0 == 0xC0.

## Solution

```go
func Valid(b []byte) bool {
	i := 0
	for i < len(b) {
		c := b[i]
		var n int
		switch {
		case c < 0x80:
			n = 0
		case c&0xE0 == 0xC0:
			n = 1
		case c&0xF0 == 0xE0:
			n = 2
		case c&0xF8 == 0xF0:
			n = 3
		default:
			return false
		}
		for j := 1; j <= n; j++ {
			if i+j >= len(b) || b[i+j]&0xC0 != 0x80 {
				return false
			}
		}
		i += n + 1
	}
	return true
}
```

## Walkthrough

0xE2 (3-byte lead): 0xE2&0xE0=0xE0 != 0xC0, falls through to 0xF0-mask case -> n=2 correct.

## Pitfalls

- Each mask pairs a bit-width with the exact prefix value.
- Check that enough continuation bytes remain and each is `10xxxxxx`.
- Full validation also rejects overlong encodings and surrogate/range violations.
