# A Buffer Per Shard, Padded Apart

## Intuition

Per-shard state removes contention in software; padding removes it in hardware. Both are needed — eight mutexes on one cache line contend almost as badly as one mutex.

## Approach

1. Reduce `id` modulo the shard count, correcting a negative result.
2. Lock the shard, call `fn(st.buf[:0])`, store the result back.
3. Return its length.

## Solution

```go
import (
	"sync"
	"unsafe"
)

// lineSize is the coherence granule the shards must not share.
const lineSize = 64

// stripe is one shard: a lock, a scratch buffer, and padding to a line.
type stripe struct {
	mu  sync.Mutex
	buf []byte
	_   [lineSize - unsafe.Sizeof(sync.Mutex{}) - unsafe.Sizeof([]byte(nil))]byte
}

// Striped hands out per-shard scratch buffers.
type Striped struct {
	stripes []stripe
}

// NewStriped returns a Striped with n shards, each holding a size-byte buffer.
func NewStriped(n, size int) *Striped {
	if n < 1 {
		n = 1
	}
	s := &Striped{stripes: make([]stripe, n)}
	for i := range s.stripes {
		s.stripes[i].buf = make([]byte, 0, size)
	}
	return s
}

// With runs fn on the shard's scratch buffer and returns the number of
// bytes fn left in it.
//
// Each shard has its own buffer and its own lock, padded so neighbouring
// shards do not share a cache line.
//
// Examples:
//
// 	s.With(0, func(b []byte) []byte { return append(b, 'x') }) => 1
func (s *Striped) With(id int, fn func(buf []byte) []byte) int {
	i := id % len(s.stripes)
	if i < 0 {
		i = -i
	}
	st := &s.stripes[i]
	st.mu.Lock()
	defer st.mu.Unlock()
	st.buf = fn(st.buf[:0])
	return len(st.buf)
}
```

## Walkthrough

Sixteen workers over eight shards contend in pairs rather than sixteen ways, and each shard's mutex and buffer header sit alone on their line.

## Pitfalls

- Discarding `fn`'s return value, so a buffer that grew is lost.
- `id % len` alone, which indexes negatively and panics.
- Padding the slice of stripes instead of the stripe itself.
