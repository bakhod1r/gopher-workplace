# Fluent Builder

## Intuition

The copy in `Build` matters more than it looks: without it, continuing to chain after building would rewrite the slice the caller already holds.

## Approach

1. `With` and `WithAll`: append, then `return b`.
2. `Build`: copy the collected values into a fresh slice.

## Solution

```go
func (b *Builder[T]) With(v T) *Builder[T] {
	b.items = append(b.items, v)
	return b
}

func (b *Builder[T]) WithAll(vs ...T) *Builder[T] {
	b.items = append(b.items, vs...)
	return b
}

func (b *Builder[T]) Build() []T {
	out := make([]T, len(b.items))
	copy(out, b.items)
	return out
}
```

## Walkthrough

`With(1).With(2)` mutates the same builder twice and `Build` hands back an independent `[1 2]`.

## Pitfalls

- Using a value receiver, so each call builds on a discarded copy.
- Returning `b.items` from `Build`, which aliases the builder.
- Returning a new builder per call, which silently drops earlier values.
