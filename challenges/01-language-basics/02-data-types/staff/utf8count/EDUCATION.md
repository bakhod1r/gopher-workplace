# Counting runes without decoding

## Intuition

Every rune has exactly one lead byte; continuation bytes are `10xxxxxx`
(`c&0xC0 == 0x80`). So the rune count equals the number of **non-continuation**
bytes:

```go
if c&0xC0 != 0x80 { n++ }
```

## Approach

1. Bug: counts bytes that ARE continuations (==0x80), inverting the logic.
2. Each rune has exactly one non-continuation lead byte; count those.
3. Fix: if c&0xC0 != 0x80.

## Solution

```go
func Count(b []byte) int {
	n := 0
	for _, c := range b {
		if c&0xC0 != 0x80 {
			n++
		}
	}
	return n
}
```

## Walkthrough

"日本語": 9 bytes, 6 are continuations, 3 lead bytes -> count 3.

## Pitfalls

- ASCII bytes (`0xxxxxxx`) are non-continuation and count as one rune each.
- This assumes valid UTF-8; invalid bytes need a real decoder.
- `c&0xC0` isolates the top two bits — the continuation signature is `10`.
