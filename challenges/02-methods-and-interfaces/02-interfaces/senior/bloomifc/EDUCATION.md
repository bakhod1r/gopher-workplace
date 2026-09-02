# Bloom Filter

## Intuition

A Bloom filter answers "definitely not present" cheaply and "probably present" otherwise. Bits are only ever set, never cleared, which is exactly why a false negative is impossible.

## Approach

1. Reduce each hash modulo the bit count.
2. `setBit` ORs `1 << (pos%8)` into byte `pos/8`.
3. `Add` sets both hash positions.
4. `MayContain` requires both bits — a single missing bit proves absence.
5. `FilterMissing` keeps the keys the filter rejects.

## Solution

```go
func (b *Bloom) setBit(pos uint32) {
	pos %= b.m
	b.bits[pos/8] |= 1 << (pos % 8)
}

func (b *Bloom) hasBit(pos uint32) bool {
	pos %= b.m
	return b.bits[pos/8]&(1<<(pos%8)) != 0
}

func (b *Bloom) Add(key string) {
	b.setBit(hash1(key))
	b.setBit(hash2(key))
}

func (b *Bloom) MayContain(key string) bool {
	return b.hasBit(hash1(key)) && b.hasBit(hash2(key))
}
```

## Walkthrough

`TestNoFalseNegatives` adds 1000 keys and re-checks every one. Because `Add` only sets bits and `MayContain` only reads them, no later insertion can ever unset a bit an earlier key relies on.

## Pitfalls

- Using `||` in `MayContain`, which inflates the false-positive rate enormously.
- Clearing bits on some "remove" path — a plain Bloom filter cannot support deletion.
- Forgetting the modulo, which panics with an index out of range on the first large hash.
