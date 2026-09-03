# Split Once, Copy Nothing

## Intuition

Strings are immutable, so a piece of one can point straight into it. Splitting is entirely a question of where the boundary is — nothing has to move.

## Approach

1. Scan for the first `sep`.
2. Return `s[:i]`, `s[i+1:]`, true.
3. After the loop, return `s`, `""`, false.

## Solution

```go
// Cut splits s around the first occurrence of sep.
//
// When sep is absent, before is s and after is empty. Both results are
// substrings, so nothing is copied.
//
// Examples:
//
// 	Cut("a=b", '=') => "a", "b", true
func Cut(s string, sep byte) (before, after string, found bool) {
	for i := 0; i < len(s); i++ {
		if s[i] == sep {
			return s[:i], s[i+1:], true
		}
	}
	return s, "", false
}
```

## Walkthrough

For "key=value", the separator is at index 3, so the results are `s[:3]` and `s[4:]` — two headers over the original bytes.

## Pitfalls

- Returning `s[i:]` for `after`, which keeps the separator.
- Building the halves with concatenation or `[]byte` round-trips.
