# Nibble ordering in hex

## Intuition

Each byte is two hex digits: the **first** character is the high nibble (bits
4–7), the second the low nibble (bits 0–3):

```go
byte(hi<<4 | lo)
```

Swapping them (`lo<<4 | hi`) reverses each byte's nibbles — `"1a"` → `0xa1`.

## Approach

1. Bug: nibbles combined as `lo<<4|hi`, swapping high and low nibble.
2. Fix: `hi<<4|lo` — first hex digit is the high nibble.
3. Odd length and non-hex digits still return false.

## Solution

```go
func nib(c byte) (int, bool) {
	switch {
	case c >= '0' && c <= '9':
		return int(c - '0'), true
	case c >= 'a' && c <= 'f':
		return int(c-'a') + 10, true
	case c >= 'A' && c <= 'F':
		return int(c-'A') + 10, true
	}
	return 0, false
}

func Decode(s string) ([]byte, bool) {
	if len(s)%2 != 0 {
		return nil, false
	}
	out := make([]byte, 0, len(s)/2)
	for i := 0; i < len(s); i += 2 {
		hi, ok1 := nib(s[i])
		lo, ok2 := nib(s[i+1])
		if !ok1 || !ok2 {
			return nil, false
		}
		out = append(out, byte(hi<<4|lo))
	}
	return out, true
}
```

## Walkthrough

"476F": hi=4,lo=7 -> 0x47 'G'; hi=6,lo=F -> 0x6F 'o' -> "Go".

## Pitfalls

- High nibble shifts left by 4; low nibble is unshifted.
- Reject odd-length input up front.
- Accept both letter cases when decoding.
