# Search Bytes Without Building A String

## Intuition

Searching needs to compare bytes, and both operands already are bytes. The conversion exists only to satisfy a function signature — and it costs a full copy of the larger side.

## Approach

1. Handle the empty needle and the too-long needle.
2. For each start position, skip quickly unless the first byte matches.
3. Compare the rest; return true on a full match.

## Solution

```go
// Contains reports whether needle appears in haystack.
//
// An empty needle is always present. Neither operand may be converted:
// the search runs over the bytes that are already there.
//
// Examples:
//
// 	Contains([]byte("hello"), "ell") => true
func Contains(haystack []byte, needle string) bool {
	if len(needle) == 0 {
		return true
	}
	if len(needle) > len(haystack) {
		return false
	}
	for i := 0; i+len(needle) <= len(haystack); i++ {
		if haystack[i] != needle[0] {
			continue
		}
		match := true
		for j := 1; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
```

## Walkthrough

Searching "aab" in "aaa": positions 0 and 1 match the first byte and fail at the third; position 2 is too close to the end. The result is false.

## Pitfalls

- Looping `i` to `len(haystack)`, which reads past the end during the inner comparison.
- `bytes.Contains` is the real answer and is also allocation-free — the point here is why the conversion was the cost.
