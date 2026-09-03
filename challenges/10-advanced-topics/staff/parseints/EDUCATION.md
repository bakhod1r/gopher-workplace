# Parse Without Making A Single String

## Intuition

Splitting and converting exist to make the data convenient, not to make it parseable. Digits are already digits: one pass with an accumulator and a field start index does the whole job in the caller's memory.

## Approach

1. Walk `i` from 0 to `len(line)` inclusive, treating the end as a separator.
2. For each field, handle an optional sign, then fold digits into an accumulator.
3. Reject empty fields, a lone sign, and any non-digit byte with `ErrSyntax`.
4. Add to the total and count the field.

## Solution

```go
import "errors"

// ErrSyntax is returned for a field that is not a decimal integer.
var ErrSyntax = errors.New("invalid integer")

// ParseInts sums the decimal integers in line, which are separated by
// sep, and returns the total, the count parsed, and any error.
//
// No part of line may be converted to a string: the parse works on the
// bytes and allocates nothing.
//
// Examples:
//
// 	ParseInts([]byte("1,2,3"), ',') => 6, 3, nil
func ParseInts(line []byte, sep byte) (int64, int, error) {
	if len(line) == 0 {
		return 0, 0, nil
	}
	var total int64
	count := 0
	start := 0
	for i := 0; i <= len(line); i++ {
		if i < len(line) && line[i] != sep {
			continue
		}
		field := line[start:i]
		start = i + 1
		if len(field) == 0 {
			return 0, 0, ErrSyntax
		}
		neg := false
		j := 0
		if field[0] == '-' || field[0] == '+' {
			neg = field[0] == '-'
			j = 1
		}
		if j == len(field) {
			return 0, 0, ErrSyntax
		}
		var v int64
		for ; j < len(field); j++ {
			c := field[j]
			if c < '0' || c > '9' {
				return 0, 0, ErrSyntax
			}
			v = v*10 + int64(c-'0')
		}
		if neg {
			v = -v
		}
		total += v
		count++
	}
	return total, count, nil
}
```

## Walkthrough

"1,2,3" closes fields at indices 1, 3 and 5 (the virtual separator at the end). Each field folds to 1, 2 and 3, giving 6 over three fields, with no allocation at any point.

## Pitfalls

- Stopping the loop at `len(line)-1`, which silently drops the last field.
- Returning a formatted error, which reintroduces an allocation on the failure path.
- Accepting an empty field as zero — the spec says it is a syntax error.
