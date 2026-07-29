# Parsing integers by hand

## Intuition

Each digit shifts the running value up one decimal place: `n = n*10 + (c-'0')`.
`c - '0'` converts a digit character to its numeric value because ASCII digits
are contiguous.

## Approach

1. Empty string -> (0,false). 2. Detect a leading '-'; if the string is just '-' it's invalid. 3. Scan remaining bytes; each must be '0'..'9'. 4. Accumulate n = n*10 + (c-'0'). 5. Negate if a '-' was seen.

## Solution

```go
func Parse(s string) (int, bool) {
	if s == "" {
		return 0, false
	}
	neg := false
	i := 0
	if s[0] == '-' {
		neg = true
		i = 1
	}
	if i >= len(s) {
		return 0, false
	}
	n := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}
	if neg {
		n = -n
	}
	return n, true
}
```

## Walkthrough

Parse("-17"): neg=true, i=1. c='1' -> n=1. c='7' -> n=17. negate -> -17. return (-17,true).

## Pitfalls

- Range over **bytes** here (ASCII digits); ranging runes also works but bytes
  are simpler for digits.
- Handle the sign once, after (or before) the digit loop — not per digit.
- Real parsers also detect overflow; this simplified one does not.
