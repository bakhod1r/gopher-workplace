# Batcher

## Intuition

Transferring ownership rather than copying is safe precisely because the batcher drops its reference; keeping both would let the caller's writes reach the next batch.

## Approach

1. `NewBatcher`: clamp the size to at least 1.
2. `Add`: append, then hand over and reset when full.
3. `Flush`: hand over whatever remains.

## Solution

```go
func NewBatcher[T any](size int) *Batcher[T] {
	if size < 1 {
		size = 1
	}
	return &Batcher[T]{size: size}
}

func (b *Batcher[T]) Add(v T) ([]T, bool) {
	b.buf = append(b.buf, v)
	if len(b.buf) < b.size {
		return nil, false
	}
	batch := b.buf
	b.buf = nil
	return batch, true
}

func (b *Batcher[T]) Flush() ([]T, bool) {
	if len(b.buf) == 0 {
		return nil, false
	}
	batch := b.buf
	b.buf = nil
	return batch, true
}
```

## Walkthrough

With size 2, the second `Add` returns `[1 2]` and leaves the batcher empty for the next pair.

## Pitfalls

- Returning the buffer without clearing it, so the next batch repeats items.
- Returning a partial batch from `Add`.
- Losing the remainder by omitting `Flush` from the shutdown path.
