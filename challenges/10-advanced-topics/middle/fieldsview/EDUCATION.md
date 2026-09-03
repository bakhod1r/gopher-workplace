# Split Without Copying The Text

## Intuition

Strings are immutable, so a substring can point straight into the original bytes with no risk. Slicing is free; only the slice of headers has to be allocated, and its size is knowable in advance.

## Approach

1. Count the separators to size the result exactly.
2. Walk the bytes, cutting a substring at each separator.
3. Append the final piece after the loop.

## Solution

```go
import "strings"

// Fields splits s on sep and returns the pieces.
//
// Substrings of a string share the original bytes, so the pieces cost
// nothing but their headers. Only the header slice may be allocated.
//
// Examples:
//
// 	Fields("a,b,c", ',') => []string{"a", "b", "c"}
func Fields(s string, sep byte) []string {
	n := strings.Count(s, string(sep)) + 1
	out := make([]string, 0, n)
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == sep {
			out = append(out, s[start:i])
			start = i + 1
		}
	}
	return append(out, s[start:])
}
```

## Walkthrough

For "a,b,c": cuts at index 1 and 3 give "a" and "b"; the tail "c" is appended after the loop. Three headers, zero bytes copied.

## Pitfalls

- `string([]byte(...))` anywhere in the loop — that is the copy you are avoiding.
- Dropping the trailing piece by forgetting the append after the loop.
