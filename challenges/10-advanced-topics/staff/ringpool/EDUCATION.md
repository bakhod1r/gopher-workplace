# Hand Out Buffers Without A Lock On The Fast Path

## Intuition

A lock-free hand-off works only if claiming is indivisible. `Swap` returns the old value and installs nil in one step, so two goroutines racing on a slot cannot both come away with the buffer.

## Approach

1. Walk the slots, `Swap(nil)` on each.
2. Return the first non-nil buffer, resliced to length 0.
3. If every slot was empty, allocate a fresh one.

## Solution

```go
import "sync/atomic"

// BufPool hands out fixed-size buffers from a bounded ring.
type BufPool struct {
	size  int
	next  atomic.Int64
	slots []atomic.Pointer[[]byte]
}

// NewBufPool returns a pool of n slots holding size-byte buffers.
func NewBufPool(n, size int) *BufPool {
	if n < 1 {
		n = 1
	}
	if size < 1 {
		size = 1
	}
	return &BufPool{size: size, slots: make([]atomic.Pointer[[]byte], n)}
}

// Put returns a buffer to the ring, dropping it if the ring is full or the
// buffer is the wrong size.
func (p *BufPool) Put(b []byte) {
	if cap(b) != p.size {
		return
	}
	b = b[:0]
	i := int(p.next.Add(1)-1) % len(p.slots)
	if i < 0 {
		i = -i
	}
	p.slots[i].CompareAndSwap(nil, &b)
}

// Get returns a buffer from the ring, or a fresh one when the ring is
// empty.
//
// The ring is a fixed array of slots claimed with an atomic index, so
// concurrent callers never block each other and never receive the same
// buffer twice.
//
// Examples:
//
// 	p := NewBufPool(4, 64); p.Get() => a 64-byte buffer
func (p *BufPool) Get() []byte {
	for i := range p.slots {
		if b := p.slots[i].Swap(nil); b != nil {
			return (*b)[:0]
		}
	}
	return make([]byte, 0, p.size)
}
```

## Walkthrough

With 16 goroutines and 8 slots, the fast path is one atomic swap; when the ring runs dry a caller allocates rather than blocking, and `Put` refills the ring as buffers come back.

## Pitfalls

- `Load` then `Store(nil)` — two goroutines can both see the same buffer.
- Returning the buffer without `[:0]`, so the caller appends onto the last user's bytes.
- Assuming the ring is a cache; a buffer that never comes back is simply collected.
