# Bloom Filter

## Intuition

A Bloom filter trades certainty for space. It never stores the items — only the
bits their hashes touch — so it cannot tell you an item is present, just that
all of its bits happen to be set. The asymmetry is the point: a negative answer
is always trustworthy, which is enough to skip an expensive lookup.

## Approach

1. `Add`: turn on one bit per hash function.
2. `MightContain`: require every one of those bits to be on.

## Solution

```go
func (f *Filter) Add(item string) {
	if len(item) == 0 {
		return
	}
	f.bits[hash1(item)] = true
	f.bits[hash2(item)] = true
}

func (f *Filter) MightContain(item string) bool {
	if len(item) == 0 {
		return false
	}
	return f.bits[hash1(item)] && f.bits[hash2(item)]
}
```

## Walkthrough

`hash1` is the first byte, `hash2` the last. `Add("hello")` sets bit `'h'` (104)
and bit `'o'` (111). `MightContain("hello")` checks the same two bits: true.
`MightContain("world")` checks `'w'` and `'d'`, neither set: false.

`"ho"` hashes to `'h'` and `'o'` — the very bits `"hello"` set — so the filter
says "might contain" for an item never added. That is the false positive the
test logs rather than fails on.

## Pitfalls

- **`||` instead of `&&`.** Turns the filter into "any bit matches", which
  makes false positives overwhelmingly likely.
- **Adding a `Remove` that clears bits.** It would clear bits shared with other
  items and create false *negatives*, destroying the only guarantee the
  structure offers. Real deletions need a counting Bloom filter.
- **Value receiver on `Add`.** The bit array is a fixed-size `[256]bool` — an
  array, not a slice — so it is copied wholesale and the write is lost.

## Why the array type matters

`[256]bool` is a value type. Every copy of `Filter` copies all 256 bytes. Had it
been a `[]bool`, a value receiver would have mutated the shared backing array
and the bug above would silently not happen — which is exactly the kind of
difference that makes arrays vs slices worth knowing cold.
