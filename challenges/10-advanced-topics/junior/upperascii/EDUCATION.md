# Upper-Case The Bytes You Were Given

## Intuition

`[]byte(s)` and `string(b)` each copy. When the data is already a mutable byte slice, transform it where it is.

## Approach

1. Range over `b` by index and value.
2. When the byte is in `a`..`z`, write back the upper-case byte.
3. Return `b`.

## Solution

```go
// Upper upper-cases the ASCII letters of b in place and returns b.
//
// Non-letters and non-ASCII bytes are left alone. Nothing is allocated —
// the caller's buffer is the working buffer.
//
// Examples:
//
// 	Upper([]byte("go1")) => []byte("GO1")
func Upper(b []byte) []byte {
	for i, c := range b {
		if c >= 'a' && c <= 'z' {
			b[i] = c - 'a' + 'A'
		}
	}
	return b
}
```

## Walkthrough

"go1 x": 'g' and 'o' are shifted by -32, '1' and ' ' fail the range test, 'x' is shifted. The array now reads "GO1 X" and no new memory was touched.

## Pitfalls

- Ranging a `[]byte` with `for i, c := range` gives bytes, not runes — which is what is wanted here, but the same loop over a `string` gives runes.
- Touching bytes above 0x7F; multi-byte UTF-8 must be left alone.
