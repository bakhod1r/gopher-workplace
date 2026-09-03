# Let The Caller Own The Buffer

## Intuition

An API that returns freshly allocated memory forces an allocation on every call. An API that fills a buffer lets the caller allocate once and reuse forever — the same work, none of the garbage.

## Approach

1. Range over `dst` by index and assign `v`.
2. Return `len(dst)`.

## Solution

```go
// Fill writes v into every byte of dst and returns how many bytes were
// written.
//
// The buffer belongs to the caller, so the function allocates nothing at
// all.
//
// Examples:
//
// 	Fill(make([]byte, 3), 'x') => 3, buffer is "xxx"
func Fill(dst []byte, v byte) int {
	for i := range dst {
		dst[i] = v
	}
	return len(dst)
}
```

## Walkthrough

`Fill(buf, 1)` on a 256-byte buffer writes 256 bytes and returns. No reference to `dst` survives the call, so the escape analyser has nothing to move.

## Pitfalls

- Allocating a temporary and copying it into `dst`.
- Writing past `len(dst)` into the capacity — that memory is not yours.
