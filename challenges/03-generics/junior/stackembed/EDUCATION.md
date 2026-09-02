# Embedding A Generic Type

## Intuition

Embedding forwards every method you do not override, so `Len` comes for free. The moment you declare `Push` on the outer type, the inner one is only reachable through the field name.

## Approach

1. `Push`: delegate, then store `v` and increment the counter.
2. `Pushes`: return the counter.
3. `Last`: report `false` until something is pushed.

## Solution

```go
func (t *TracedStack[T]) Push(v T) {
	t.Stack.Push(v)
	t.last = v
	t.pushes++
}

func (t *TracedStack[T]) Pushes() int {
	return t.pushes
}

func (t *TracedStack[T]) Last() (T, bool) {
	if t.pushes == 0 {
		var zero T
		return zero, false
	}
	return t.last, true
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}
```

## Walkthrough

After `Push(1); Push(2)`, `Len()` reports 2 via promotion while `Pushes()` reports 2 from the counter you maintain.

## Pitfalls

- Calling `t.Push(v)` inside `Push`, which recurses forever.
- Re-implementing the append instead of delegating, so the two states drift apart.
- Embedding `Stack` without the type argument.
