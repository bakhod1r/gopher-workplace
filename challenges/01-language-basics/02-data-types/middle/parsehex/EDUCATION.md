# Parsing hex

## Intuition

Each hex digit contributes 4 bits: `n = n*16 + d`. The digit value `d` depends on
which of three contiguous ranges the character falls in:

```go
switch {
case c >= '0' && c <= '9': d = int(c - '0')
case c >= 'a' && c <= 'f': d = int(c-'a') + 10
case c >= 'A' && c <= 'F': d = int(c-'A') + 10
default: return 0, false
}
```

## Approach

1. Empty -> (0,false). 2. Map each char: '0'-'9' -> 0-9, 'a'-'f'/'A'-'F' -> 10-15, else invalid. 3. Accumulate n = n*16 + v.

## Solution

```go
func Parse(s string) (int, bool) {
	if s == "" {
		return 0, false
	}
	n := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		var v int
		switch {
		case c >= '0' && c <= '9':
			v = int(c - '0')
		case c >= 'a' && c <= 'f':
			v = int(c-'a') + 10
		case c >= 'A' && c <= 'F':
			v = int(c-'A') + 10
		default:
			return 0, false
		}
		n = n*16 + v
	}
	return n, true
}
```

## Walkthrough

Parse("ff"): 'f'=15 -> n=15; 'f'=15 -> n=15*16+15=255. (255,true).

## Pitfalls

- Handle both letter cases.
- `n*16` can overflow for long inputs; this simplified version ignores that.
- Reject empty input explicitly, or it returns `(0, true)`.
