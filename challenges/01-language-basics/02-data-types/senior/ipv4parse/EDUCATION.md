# Validate before narrowing

## Intuition

An IPv4 octet fits in a `byte` (0..255). Validation must happen **before** the
`byte(val)` conversion, because narrowing wraps: `byte(256) == 0`. A too-loose
bound (`> 999`) lets an invalid octet convert to a wrong byte.

## Approach

1. Bug: octet range checked `val > 999`, letting 256..999 through.
2. Fix: reject `val > 255`.
3. Field count and digit checks unchanged.

## Solution

```go
func Parse(s string) (out [4]byte, ok bool) {
	field, seen, val := 0, 0, 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == '.' {
			if seen == 0 {
				return [4]byte{}, false
			}
			if val > 255 {
				return [4]byte{}, false
			}
			if field > 3 {
				return [4]byte{}, false
			}
			out[field] = byte(val)
			field++
			val, seen = 0, 0
			continue
		}
		c := s[i]
		if c < '0' || c > '9' {
			return [4]byte{}, false
		}
		val = val*10 + int(c-'0')
		seen++
	}
	if field != 4 {
		return [4]byte{}, false
	}
	return out, true
}
```

## Walkthrough

"256.1.1.1": first octet val=256 -> 256>255 -> false.

## Pitfalls

- The bound is `> 255`, exactly the byte maximum.
- Narrowing conversions never panic; they truncate silently.
- Also guard field count and empty fields, as this parser does.
