# FNV-1a vs FNV-1

## Intuition

FNV-1a hashes each byte as **XOR then multiply**:

```go
h = (h ^ uint32(b)) * prime32
```

FNV-1 does the reverse (multiply then XOR). Same constants, different avalanche
behavior and different output — 1a mixes low bits better.

## Approach

1. Bug: does multiply-then-xor, which is FNV-1, not FNV-1a.
2. FNV-1a XORs the byte first, then multiplies by the prime.
3. Fix: h = (h ^ uint32(b)) * prime32.

## Solution

```go
const (
	offset32 = 2166136261
	prime32  = 16777619
)

func Hash(data []byte) uint32 {
	h := uint32(offset32)
	for _, b := range data {
		h = (h ^ uint32(b)) * prime32
	}
	return h
}
```

## Walkthrough

"a": h=offset, h^=0x61 then *prime32 -> 0xe40c292c.

## Pitfalls

- The 32-bit offset basis is `0x811c9dc5`; the prime is `16777619`.
- The `uint32` multiply wraps mod 2³² — that overflow is intended.
- Empty input hashes to the offset basis, a good sanity check.
