# Positional base conversion

## Intuition

To write `n` in base `b`, repeatedly take `n % b` (the next digit) and `n /= b`,
until `n` is 0. Digits emerge least-significant first, so reverse:

```go
const digits = "0123456789abcdef"
for n > 0 { buf = append(buf, digits[n%base]); n /= base }
// reverse buf
```

## Approach

1. n==0 returns "0". 2. Repeatedly take n%base as the next digit (indexing "0123456789abcdef"). 3. Divide n by base. 4. Digits come out least-significant first, so reverse the buffer before returning.

## Solution

```go
func Format(n, base int) string {
	if n == 0 {
		return "0"
	}
	const digits = "0123456789abcdef"
	var buf []byte
	for n > 0 {
		buf = append(buf, digits[n%base])
		n /= base
	}
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
```

## Walkthrough

Format(255,16): 255%16=15 -> 'f', n=15; 15%16=15 -> 'f', n=0. buf=['f','f'], reversed still "ff".

## Pitfalls

- Handle `n == 0` specially, or you return an empty string.
- Indexing `digits[n%base]` gives a `byte`; build a `[]byte` and convert once.
- This version assumes non-negative `n`; a sign needs separate handling.
