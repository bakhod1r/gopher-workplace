# Join Without The Intermediate Strings

## Intuition

Every `+=` on a string is an allocate-and-copy of everything so far, so joining n parts costs O(n²) bytes copied. A builder writes into one buffer that only grows — and if you pre-size it, it never grows at all.

## Approach

1. Return early for an empty input.
2. Sum the part lengths plus `len(sep) * (len(parts)-1)`.
3. `b.Grow(n)`, then write parts with separators.
4. Return `b.String()`.

## Solution

```go
import "strings"

// Join concatenates parts separated by sep.
//
// Strings are immutable, so `s += p` allocates a new string every round.
// Build the result in one growing buffer instead.
//
// Examples:
//
// 	Join([]string{"a", "b"}, "-") => "a-b"
func Join(parts []string, sep string) string {
	if len(parts) == 0 {
		return ""
	}
	n := len(sep) * (len(parts) - 1)
	for _, p := range parts {
		n += len(p)
	}
	var b strings.Builder
	b.Grow(n)
	for i, p := range parts {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(p)
	}
	return b.String()
}
```

## Walkthrough

Joining 64 five-byte parts with a two-byte separator needs 64*5 + 63*2 = 446 bytes. One `Grow(446)` allocates once; `+=` would have allocated 64 times and copied about 14 KB.

## Pitfalls

- Adding the separator after every part and trimming the tail — an extra copy for nothing.
- Forgetting that `len` on a string is bytes, which is exactly what the buffer needs.
