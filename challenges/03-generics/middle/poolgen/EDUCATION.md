# Object Pool

## Intuition

The factory parameter is what makes the pool generic: without it, the pool could not construct a `T` it knows nothing about.

## Approach

1. `Get`: build when empty, otherwise pop the last free value.
2. `Put`: append.
3. `Idle`: report the length.

## Solution

```go
func NewPool[T any](make func() T) *Pool[T] {
	return &Pool[T]{make: make}
}

func (p *Pool[T]) Get() T {
	if len(p.free) == 0 {
		return p.make()
	}
	v := p.free[len(p.free)-1]
	p.free = p.free[:len(p.free)-1]
	return v
}

func (p *Pool[T]) Put(v T) {
	p.free = append(p.free, v)
}

func (p *Pool[T]) Idle() int {
	return len(p.free)
}
```

## Walkthrough

`Put(a); Get()` pops `a` straight back out without calling the factory.

## Pitfalls

- Calling the factory even when a value is available.
- Handing the same value out twice by forgetting to shrink the free list.
- Claiming concurrency safety without synchronisation.
