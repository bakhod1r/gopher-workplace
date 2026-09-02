# Multiset

## Intuition

The invariant "every stored key has a positive count" is what makes `Distinct` a simple `len`, and it must be restored by every removal.

## Approach

1. `Add`: increment.
2. `Remove`: report `false` when absent; delete at one, decrement otherwise.
3. `Count` and `Distinct`: read the map.

## Solution

```go
func NewBag[T comparable]() *Bag[T] {
	return &Bag[T]{counts: make(map[T]int)}
}

func (b *Bag[T]) Add(v T) {
	b.counts[v]++
}

func (b *Bag[T]) Remove(v T) bool {
	n, ok := b.counts[v]
	if !ok || n == 0 {
		return false
	}
	if n == 1 {
		delete(b.counts, v)
	} else {
		b.counts[v] = n - 1
	}
	return true
}

func (b *Bag[T]) Count(v T) int {
	return b.counts[v]
}

func (b *Bag[T]) Distinct() int {
	return len(b.counts)
}
```

## Walkthrough

`Add(a); Remove(a)` deletes the key entirely, so `Distinct()` is `0`, not `1`.

## Pitfalls

- Leaving zero-count keys behind.
- Returning `true` from `Remove` for an unknown value.
- Letting counts go negative.
