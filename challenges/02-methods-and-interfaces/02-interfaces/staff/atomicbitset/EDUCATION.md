# Atomic Bitset

## Intuition

There is no atomic single-bit instruction here — setting a bit is a read-modify-write of a 64-bit word. Two goroutines touching different bits in the same word will lose one update unless the write is a CAS.

## Approach

1. Locate the word (`i/64`) and the mask (`1 << (i%64)`).
2. Loop: load the word, return `false` if the bit is already in the desired state.
3. CAS the new word; retry on failure because another goroutine changed a different bit.
4. `Count` popcounts every word with `bits.OnesCount64`.

## Solution

```go
func (b *Bitset) Set(i int) bool {
	if !b.inRange(i) {
		return false
	}
	w := &b.words[i/64]
	mask := uint64(1) << (uint(i) % 64)

	for {
		old := w.Load()
		if old&mask != 0 {
			return false
		}
		if w.CompareAndSwap(old, old|mask) {
			return true
		}
	}
}

func (b *Bitset) Clear(i int) bool {
	if !b.inRange(i) {
		return false
	}
	w := &b.words[i/64]
	mask := uint64(1) << (uint(i) % 64)

	for {
		old := w.Load()
		if old&mask == 0 {
			return false
		}
		if w.CompareAndSwap(old, old&^mask) {
			return true
		}
	}
}
```

## Walkthrough

`TestConcurrentSetsInOneWord` puts all 64 bits in a single word and sets them from 64 goroutines. A non-atomic `|=` would lose most of them; the CAS retry loop keeps every one.

## Pitfalls

- `w.Store(w.Load() | mask)` — the read and the write are separate, so concurrent bits are dropped.
- Returning `true` from the CAS-failure path, which over-reports changes.
- `old & ^mask` with a signed shift or the wrong precedence; use `&^` for clearing.
