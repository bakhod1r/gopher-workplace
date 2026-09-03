# Aggregate A Stream With No Garbage At All

## Intuition

Every convenience in a parser — split, convert, format — is an allocation. At a million lines a second the only viable design is one that reads the caller's bytes and produces nothing but numbers.

## Approach

1. For each non-empty line, walk `i` to `len(line)` inclusive, treating the end as a separator.
2. Per field: optional sign, then fold digits; reject empties, lone signs and non-digits.
3. Accumulate the total and the count.

## Solution

```go
import "errors"

// ErrSyntax marks a field that is not a decimal integer.
var ErrSyntax = errors.New("invalid integer")

// Aggregate sums the decimal integers across every line and reports the
// total and the field count.
//
// The whole aggregation must run without allocating: no conversions, no
// split slices, no formatted errors.
//
// Examples:
//
// 	Aggregate([][]byte{[]byte("1,2")}, ',') => 3, 2, nil
func Aggregate(lines [][]byte, sep byte) (int64, int, error) {
	var total int64
	count := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
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
	}
	return total, count, nil
}
```

## Walkthrough

For "-4,+5": the first field folds to 4 and is negated; the second folds to 5. The total gains 1 and the count gains 2, with no memory touched beyond the caller's slices.

## Pitfalls

- Returning `fmt.Errorf` for a bad field, which allocates on precisely the input an attacker controls.
- Stopping the inner loop at `len(line)-1`, which drops the last field.
- Treating an empty line as a syntax error; the spec skips it.
