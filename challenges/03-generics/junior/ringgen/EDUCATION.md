# Ring Buffer

## Intuition

The simple append-and-trim keeps the code readable; the copy in `Items` is what stops a caller's `append` from writing into the ring's own storage.

## Approach

1. `NewRing`: clamp a negative size to 0 and store it.
2. `Add`: ignore when size is 0; append, then drop the front element if over capacity.
3. `Items`: return a copy of the buffered slice.

## Solution

```go
func NewRing[T any](size int) *Ring[T] {
	if size < 0 {
		size = 0
	}
	return &Ring[T]{size: size}
}

func (r *Ring[T]) Add(v T) {
	if r.size == 0 {
		return
	}
	r.items = append(r.items, v)
	if len(r.items) > r.size {
		r.items = r.items[1:]
	}
}

func (r *Ring[T]) Items() []T {
	out := make([]T, len(r.items))
	copy(out, r.items)
	return out
}
```

## Walkthrough

`Add(3)` on a full ring of `[1 2]` appends to make `[1 2 3]`, then reslices to `[2 3]`.

## Pitfalls

- Returning `r.items` directly, so a caller's `append` mutates the ring.
- Trimming from the back, which drops the newest element instead of the oldest.
- Panicking on a zero-size ring.
