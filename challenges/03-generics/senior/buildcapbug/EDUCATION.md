# The Builder That Hands Out Its Buffer

## Intuition

Returning `b.items` shares the backing array, so a caller's `append` can write into the builder's spare capacity and corrupt later builds.

## Approach

1. Allocate a slice of the same length.
2. Copy the items in.
3. Return the copy.

## Solution

```go
func (b *Builder[T]) Build() []T {
	out := make([]T, len(b.items))
	copy(out, b.items)
	return out
}

func (b *Builder[T]) Add(v T) *Builder[T] {
	b.items = append(b.items, v)
	return b
}
```

## Walkthrough

After `Build()`, a caller appending `99` writes into the builder's capacity; the next `Add` then overwrites it — or reads it back.

## Pitfalls

- Returning `b.items[:len(b.items):len(b.items)]`, which stops the overwrite but still shares reads.
- Copying with `out := b.items` — that copies the header, not the data.
