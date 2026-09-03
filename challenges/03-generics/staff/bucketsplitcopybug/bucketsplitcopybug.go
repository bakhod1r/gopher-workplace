// Package bucketsplitcopybug — Gopher Workplace challenge.
package bucketsplitcopybug

import (
	"cmp"
)

// entry is one key/value pair inside a bucket.
type entry[K cmp.Ordered, V any] struct {
	key K
	val V
}

// Index is an ordered index made of sorted buckets, each holding at most Max entries.
type Index[K cmp.Ordered, V any] struct {
	buckets [][]entry[K, V]
	Max     int
}

// split divides an overfull bucket in two.
// The lower half stays at bi; the upper half becomes a new bucket at bi+1.
func (ix *Index[K, V]) split(bi int) {
	// CHANGE CODE BELOW THIS LINE
	b := ix.buckets[bi]
	mid := len(b) / 2
	right := make([]entry[K, V], len(b)-mid)
	copy(right, b[mid:])
	ix.buckets = append(ix.buckets, nil)
	copy(ix.buckets[bi+2:], ix.buckets[bi+1:])
	ix.buckets[bi+1] = right
	// CHANGE CODE ABOVE THIS LINE
}

// Insert stores v under k, splitting a bucket that grows past Max.
func (ix *Index[K, V]) Insert(k K, v V) {
	if ix.Max <= 0 {
		ix.Max = 4
	}
	if len(ix.buckets) == 0 {
		ix.buckets = [][]entry[K, V]{{}}
	}
	bi := 0
	for bi < len(ix.buckets)-1 {
		b := ix.buckets[bi]
		if len(b) > 0 && k > b[len(b)-1].key {
			bi++
			continue
		}
		break
	}
	b := ix.buckets[bi]
	i := 0
	for i < len(b) && b[i].key < k {
		i++
	}
	if i < len(b) && b[i].key == k {
		b[i].val = v
		return
	}
	b = append(b, entry[K, V]{})
	copy(b[i+1:], b[i:])
	b[i] = entry[K, V]{key: k, val: v}
	ix.buckets[bi] = b
	if len(b) > ix.Max {
		ix.split(bi)
	}
}

// Get returns the value stored under k.
func (ix *Index[K, V]) Get(k K) (V, bool) {
	for _, b := range ix.buckets {
		for _, e := range b {
			if e.key == k {
				return e.val, true
			}
		}
	}
	var zero V
	return zero, false
}

// Keys lists every key in ascending order.
func (ix *Index[K, V]) Keys() []K {
	out := make([]K, 0)
	for _, b := range ix.buckets {
		for _, e := range b {
			out = append(out, e.key)
		}
	}
	return out
}
