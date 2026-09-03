# The String That Changed Underneath Its Owner

## Intuition

`unsafe.String` is safe exactly when nothing will ever write to the bytes again. Over a buffer the caller is about to reuse, that condition is false by construction — and nothing will report the violation.

## Approach

1. Clamp `n` and handle the empty cases.
2. Allocate an `n`-byte slice and copy the prefix into it.
3. Wrap the copy with `unsafe.String`, which adds no second copy.

## Solution

```go
import "unsafe"

// Take returns the first n bytes of buf as a string the caller keeps.
//
// buf is a scratch buffer the caller reuses, so the result must not be a
// view of it — a string that changes is a contradiction the rest of the
// program is not prepared for.
//
// Examples:
//
// 	Take([]byte("hello"), 2) => "he", independent of buf
func Take(buf []byte, n int) string {
	if n <= 0 || len(buf) == 0 {
		return ""
	}
	if n > len(buf) {
		n = len(buf)
	}
	out := make([]byte, n)
	copy(out, buf)
	return unsafe.String(unsafe.SliceData(out), n)
}
```

## Walkthrough

Before the fix, `Take(buf, 5)` returns a header pointing into `buf`; the next `copy(buf, "SECOND")` rewrites the string in place. After it, the string points at a private array nothing else can reach.

## Pitfalls

- `string(buf[:n])` is also correct here — the point is that the copy is what was missing, not the conversion.
- Copying and then converting with `string(out)`, which copies a second time.
