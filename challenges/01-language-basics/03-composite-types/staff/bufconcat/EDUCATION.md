# Writing into a buffer at an offset

## Intuition

Filling a preallocated buffer means copying each piece to its correct offset. The
second slice starts where the first ends:

```go
out := make([]byte, len(a)+len(b))
copy(out, a)
copy(out[len(a):], b) // offset by len(a)
```

## Approach

1. Bug: `copy(out, b)` writes b at offset 0, overwriting the copy of a.
2. Fix: `copy(out[len(a):], b)` writes b right after a.

## Solution

```go
func Concat(a, b []byte) []byte {
	out := make([]byte, len(a)+len(b))
	copy(out, a)
	copy(out[len(a):], b)
	return out
}
```

## Walkthrough

out is zeroed len 4. copy(out,a) -> [1 2 0 0]. Bug copy(out,b) -> [3 4 0 0]. Fixed copy(out[2:],b) -> [1 2 3 4].

## Pitfalls

- `copy` returns the count; the destination sub-slice sets where it lands.
- Size the buffer exactly, or trailing zeros/overflows appear.
- For many pieces, `bytes.Buffer` or `append` is simpler than manual offsets.
