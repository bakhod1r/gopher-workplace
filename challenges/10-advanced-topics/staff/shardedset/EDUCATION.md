# A Set That Many Goroutines Can Share

## Intuition

Striping is about which lock, never about whether. The whole correctness of `Add` rests on the presence check and the insert happening without a gap — that is what makes "newly added" meaningful.

## Approach

1. Route to the shard.
2. Lock it; return false if the key is present.
3. Insert and return true.

## Solution

```go
import (
	"hash/maphash"
	"sync"
)

// bucket is one shard of the set, padded to a cache line.
type bucket struct {
	mu sync.Mutex
	m  map[string]struct{}
	_  [48]byte
}

// Set is a striped concurrent string set.
type Set struct {
	seed    maphash.Seed
	buckets []bucket
}

// NewSet returns a set with n shards.
func NewSet(n int) *Set {
	if n < 1 {
		n = 1
	}
	s := &Set{seed: maphash.MakeSeed(), buckets: make([]bucket, n)}
	for i := range s.buckets {
		s.buckets[i].m = make(map[string]struct{})
	}
	return s
}

// bucketFor returns the shard owning key.
func (s *Set) bucketFor(key string) *bucket {
	h := maphash.String(s.seed, key)
	return &s.buckets[h%uint64(len(s.buckets))]
}

// Has reports whether key is present.
func (s *Set) Has(key string) bool {
	b := s.bucketFor(key)
	b.mu.Lock()
	defer b.mu.Unlock()
	_, ok := b.m[key]
	return ok
}

// Len reports the total number of keys.
func (s *Set) Len() int {
	n := 0
	for i := range s.buckets {
		b := &s.buckets[i]
		b.mu.Lock()
		n += len(b.m)
		b.mu.Unlock()
	}
	return n
}

// Add inserts key and reports whether it was newly added.
//
// The set is striped across shards so concurrent writers of different keys
// rarely contend, and each shard is padded onto its own cache line.
//
// Examples:
//
// 	s := NewSet(4); s.Add("a") => true, then false
func (s *Set) Add(key string) bool {
	b := s.bucketFor(key)
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, ok := b.m[key]; ok {
		return false
	}
	b.m[key] = struct{}{}
	return true
}
```

## Walkthrough

Sixteen workers racing on one key all reach the same shard. One holds the lock, finds it absent and inserts; the rest find it present and report false.

## Pitfalls

- Checking with `Has` and then calling a separate insert — two lock acquisitions with a race between them.
- Locking every shard, which is correct and no better than one mutex.
- Re-seeding the hash per call, which would scatter one key across shards.
