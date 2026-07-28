# FNV-1a vs FNV-1

## The idea

FNV-1a hashes each byte as **XOR then multiply**:

```go
h = (h ^ uint32(b)) * prime32
```

FNV-1 does the reverse (multiply then XOR). Same constants, different avalanche
behavior and different output — 1a mixes low bits better.

## Why it matters

Hashes are all-or-nothing across systems. Swapping the two operations yields a
valid-looking but incompatible hash, so anything sharding or deduplicating by
hash silently disagrees. Order matters as much as the constants.

## Watch out

- The 32-bit offset basis is `0x811c9dc5`; the prime is `16777619`.
- The `uint32` multiply wraps mod 2³² — that overflow is intended.
- Empty input hashes to the offset basis, a good sanity check.
