# A View The Caller Cannot Corrupt

## Intuition

A slice's capacity is a licence to write. Handing out a view with capacity beyond its length hands out a licence over memory the caller was never given — and `append` will use it without a word.

## Approach

1. Validate `off`, `n` and `off+n` against `len(b)`.
2. Return nil for a zero-length window.
3. Build the view from the offset pointer and cap it with `[:n:n]`.

## Solution

```go
import "unsafe"

// Window returns the n bytes of b starting at off, as a view whose
// capacity is exactly n.
//
// The caller may append to the result, so the capacity must not let that
// append reach the bytes after the window.
//
// Examples:
//
// 	Window(buf, 2, 3) => buf[2:5] with capacity 3, true
func Window(b []byte, off, n int) ([]byte, bool) {
	if off < 0 || n < 0 || off+n > len(b) {
		return nil, false
	}
	if n == 0 {
		return nil, true
	}
	p := unsafe.Add(unsafe.Pointer(unsafe.SliceData(b)), off)
	return unsafe.Slice((*byte)(p), n)[:n:n], true
}
```

## Walkthrough

Without the capacity cap, `Window(b, 2, 3)` on a six-byte buffer yields capacity 4, so one append overwrites `b[5]`. With `[:n:n]` the append is forced to allocate and the buffer is untouched.

## Pitfalls

- `b[off : off+n]` alone — correct bytes, wrong capacity.
- Checking `off < len(b)` instead of `off+n <= len(b)`.
- Building the window with a nil data pointer when `n` is 0.
