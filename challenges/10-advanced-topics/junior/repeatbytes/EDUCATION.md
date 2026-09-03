# Repeat With The Length You Already Know

## Intuition

The output length is completely determined by the inputs, so there is no reason to discover it by growing. Allocate it, then fill it.

## Approach

1. Handle the empty cases.
2. `make([]byte, len(b)*n)`.
3. `copy` `b` into the window starting at `i*len(b)` for each i.

## Solution

```go
// Repeat returns n concatenated copies of b in a freshly allocated
// slice that shares nothing with b.
//
// For n <= 0 the result is empty. The allocation must happen once, at the
// final size.
//
// Examples:
//
// 	Repeat([]byte("ab"), 2) => []byte("abab")
func Repeat(b []byte, n int) []byte {
	if n <= 0 || len(b) == 0 {
		return []byte{}
	}
	out := make([]byte, len(b)*n)
	for i := 0; i < n; i++ {
		copy(out[i*len(b):], b)
	}
	return out
}
```

## Walkthrough

`Repeat([]byte("ab"), 3)` allocates six bytes, then copies "ab" at offsets 0, 2 and 4.

## Pitfalls

- `append(out, b...)` in a loop — correct, but that is the reallocating version.
- Forgetting that `len(b) == 0` makes `len(b)*n` zero, which is fine, but `n <= 0` must be checked too.
