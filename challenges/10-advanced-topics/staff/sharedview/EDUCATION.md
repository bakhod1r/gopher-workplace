# Many Readers Over One Reinterpreted Buffer

## Intuition

Reinterpretation and parallelism compose cleanly because both are read-only: the view aliases the caller's bytes, and disjoint readers of shared memory never need to coordinate. The only synchronisation is the join.

## Approach

1. Reject an empty or non-multiple-of-eight length and a misaligned data pointer.
2. Build the view with `unsafe.Slice((*int64)(p), len(b)/8)`.
3. Clamp `workers`, split into ceiling-divided clamped chunks, sum into per-worker slots.
4. `Wait`, then fold the partials.

## Solution

```go
import (
	"sync"
	"unsafe"
)

// SumParallel reinterprets b as int64 values and totals them using
// workers goroutines over disjoint chunks.
//
// The view shares b's storage, so nothing may write through it. Concurrent
// reads of shared memory need no synchronisation at all.
//
// Examples:
//
// 	SumParallel(sixteenBytes, 2) => the two int64s summed, true
func SumParallel(b []byte, workers int) (int64, bool) {
	if len(b) == 0 || len(b)%8 != 0 {
		return 0, false
	}
	p := unsafe.Pointer(unsafe.SliceData(b))
	if uintptr(p)&7 != 0 {
		return 0, false
	}
	view := unsafe.Slice((*int64)(p), len(b)/8)
	if workers < 1 {
		workers = 1
	}
	if workers > len(view) {
		workers = len(view)
	}
	partial := make([]int64, workers)
	size := (len(view) + workers - 1) / workers
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		start := w * size
		end := start + size
		if start > len(view) {
			start = len(view)
		}
		if end > len(view) {
			end = len(view)
		}
		go func(w int, part []int64) {
			defer wg.Done()
			var sum int64
			for _, v := range part {
				sum += v
			}
			partial[w] = sum
		}(w, view[start:end])
	}
	wg.Wait()
	var total int64
	for _, v := range partial {
		total += v
	}
	return total, true
}
```

## Walkthrough

A 512 KiB buffer becomes a 65536-element view with no copy. Four workers each sum 16384 elements from memory they alone read, and the four partials are added after the join.

## Pitfalls

- Passing `len(b)` as the element count, which produces a view eight times too long.
- Writing through the view — the caller's bytes may be a read-only mapping, and other readers would see the change.
- Folding the partials before `Wait`.
