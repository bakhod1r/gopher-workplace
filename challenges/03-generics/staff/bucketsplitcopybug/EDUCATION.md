# The Split That Keeps Both Halves

## Intuition

A split partitions a bucket. The new right-hand bucket is built correctly, but the original is never truncated to its lower half, so every element at or above the midpoint now lives in two buckets at once.

## Approach

1. Copy the upper half into a fresh bucket.
2. Replace the original bucket with just its lower half.
3. Grow the bucket list by one and shift the tail right to make room.
4. Drop the new bucket in at bi+1.

## Solution

```go
func (ix *Index[K, V]) split(bi int) {
	b := ix.buckets[bi]
	mid := len(b) / 2
	right := make([]entry[K, V], len(b)-mid)
	copy(right, b[mid:])
	left := make([]entry[K, V], mid)
	copy(left, b[:mid])
	ix.buckets[bi] = left
	ix.buckets = append(ix.buckets, nil)
	copy(ix.buckets[bi+2:], ix.buckets[bi+1:])
	ix.buckets[bi+1] = right
}

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

func (ix *Index[K, V]) Keys() []K {
	out := make([]K, 0)
	for _, b := range ix.buckets {
		for _, e := range b {
			out = append(out, e.key)
		}
	}
	return out
}
```

## Walkthrough

Splitting `[1 2 3 4 5]` produces a right bucket `[3 4 5]` while the original still reads `[1 2 3 4 5]`, so `Keys()` returns `1 2 3 4 5 3 4 5`.

## Pitfalls

- Assigning `ix.buckets[bi] = b[:mid]` and then reusing `b` — correct here, but it keeps the upper half alive in the same array.
- Appending the new bucket at the end instead of inserting it at bi+1, which destroys the global ordering.
- Shifting the tail before growing the slice, which writes past the end.
