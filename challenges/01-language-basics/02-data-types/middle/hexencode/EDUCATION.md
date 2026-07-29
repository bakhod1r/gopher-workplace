# Hex encoding

## Intuition

A byte is two 4-bit nibbles. Split them and map each through a digit table:

```go
const hexd = "0123456789abcdef"
hi := hexd[b>>4]
lo := hexd[b&0x0f]
```

## Approach

1. For each byte, take the high nibble (c>>4) and low nibble (c&0x0f). 2. Index "0123456789abcdef" for each. 3. Append both chars.

## Solution

```go
func Encode(b []byte) string {
	const digits = "0123456789abcdef"
	out := make([]byte, 0, len(b)*2)
	for _, c := range b {
		out = append(out, digits[c>>4], digits[c&0x0f])
	}
	return string(out)
}
```

## Walkthrough

Encode([0x1a]): 0x1a>>4=1->'1', 0x1a&0xf=10->'a'. "1a".

## Pitfalls

- Mask the low nibble with `& 0x0f`; `>>4` already drops the low bits for the
  high one.
- Build with a `[]byte`/`strings.Builder` and convert once — not `+=` per byte.
- Output is fixed two chars per byte, including leading zeros (`0x00` → "00").
