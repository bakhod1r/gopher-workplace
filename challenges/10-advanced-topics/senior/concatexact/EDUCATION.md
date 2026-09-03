# Join With Exactly One Allocation

## Intuition

The rule for `unsafe.String` is that the bytes must never change. A buffer allocated inside the function, written once, and never referenced again satisfies that by construction — so the final copy buys nothing.

## Approach

1. Sum the parts' lengths; return `""` for zero.
2. `make([]byte, 0, n)` and append every part.
3. Wrap it with `unsafe.String`.

## Solution

```go
import "unsafe"

// Concat returns the parts joined end to end.
//
// The final length is the sum of the parts' lengths, so the result can be
// built in one allocation and handed out without a second copy.
//
// Examples:
//
// 	Concat([]string{"a", "bc"}) => "abc"
func Concat(parts []string) string {
	n := 0
	for _, p := range parts {
		n += len(p)
	}
	if n == 0 {
		return ""
	}
	buf := make([]byte, 0, n)
	for _, p := range parts {
		buf = append(buf, p...)
	}
	return unsafe.String(unsafe.SliceData(buf), len(buf))
}
```

## Walkthrough

Joining five parts totalling 27 bytes allocates one 27-byte buffer. `string(buf)` would allocate a second 27 bytes and copy — for bytes only this function has ever seen.

## Pitfalls

- Keeping a reference to `buf` after wrapping it; any later write breaks the string's immutability.
- Skipping the zero-length guard, which would wrap a nil data pointer.
