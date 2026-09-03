# The Concatenation That Copied Everything Each Time

## Intuition

Each `+=` allocates a string as long as everything so far and copies it. By the last part you have copied the whole report once per fragment — the cost is invisible at three parts and fatal at three hundred.

## Approach

1. Sum the parts' lengths.
2. `Grow` a `strings.Builder` to that size.
3. Write every part and return `b.String()`.

## Solution

```go
import "strings"

// Join concatenates parts end to end.
//
// Strings are immutable, so += allocates a new string and copies both sides
// every round — quadratic in the total length.
//
// Examples:
//
// 	Join([]string{"a", "bc"}) => "abc"
func Join(parts []string) string {
	n := 0
	for _, p := range parts {
		n += len(p)
	}
	var b strings.Builder
	b.Grow(n)
	for _, p := range parts {
		b.WriteString(p)
	}
	return b.String()
}
```

## Walkthrough

128 five-byte parts total 640 bytes. `+=` allocates 128 strings and copies about 41,000 bytes; the builder allocates once and copies 640.

## Pitfalls

- Using a builder without `Grow`, which is linear but still reallocates a few times.
- `strings.Join(parts, "")` is the real-world answer — the point is why `+=` is not.
