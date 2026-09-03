# Compare Bytes Without Making A String

## Intuition

`string(b)` has to copy, because the resulting string must be immutable while `b` is not. But comparing does not need a string at all — both sides index to bytes.

## Approach

1. Return false when `prefix` is longer than `b`.
2. Compare byte by byte up to `len(prefix)`.
3. Return true if none differed.

## Solution

```go
// HasPrefix reports whether b begins with prefix.
//
// Converting b to a string copies it. The comparison can be done on the
// bytes that are already there.
//
// Examples:
//
// 	HasPrefix([]byte("hello"), "he") => true
func HasPrefix(b []byte, prefix string) bool {
	if len(prefix) > len(b) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if b[i] != prefix[i] {
			return false
		}
	}
	return true
}
```

## Walkthrough

For a 896-byte frame and a 7-byte prefix, the conversion would copy 896 bytes; the byte loop reads at most 7 and allocates nothing.

## Pitfalls

- `bytes.HasPrefix` is the real-world answer and also allocation-free — the point here is why.
- Indexing before the length check, which panics on a short `b`.
