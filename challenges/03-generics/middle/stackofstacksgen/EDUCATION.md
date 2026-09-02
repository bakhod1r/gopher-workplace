# Nested Stacks

## Intuition

Both invariants must hold after every operation — dropping the emptied stack in `Pop` is what keeps `Stacks()` meaningful and `Push` correct.

## Approach

1. `Push`: start a new inner stack when there is none or the last is full, then append.
2. `Pop`: take from the last inner stack, then drop it if it emptied.
3. `Stacks`: report the count.

## Solution

```go
func (s *PlateStack[T]) Push(v T) {
	if s.cap < 1 {
		s.cap = 1
	}
	if len(s.stacks) == 0 || len(s.stacks[len(s.stacks)-1]) == s.cap {
		s.stacks = append(s.stacks, make([]T, 0, s.cap))
	}
	last := len(s.stacks) - 1
	s.stacks[last] = append(s.stacks[last], v)
}

func (s *PlateStack[T]) Pop() (T, bool) {
	if len(s.stacks) == 0 {
		var zero T
		return zero, false
	}
	last := len(s.stacks) - 1
	inner := s.stacks[last]
	v := inner[len(inner)-1]
	s.stacks[last] = inner[:len(inner)-1]
	if len(s.stacks[last]) == 0 {
		s.stacks = s.stacks[:last]
	}
	return v, true
}

func (s *PlateStack[T]) Stacks() int {
	return len(s.stacks)
}

func (s *PlateStack[T]) Cap(n int) {
	if n < 1 {
		n = 1
	}
	s.cap = n
}
```

## Walkthrough

With capacity 2, three pushes fill the first tray and start a second; popping empties the second tray and removes it.

## Pitfalls

- Starting a new inner stack on every push.
- Leaving empty inner stacks behind, so `Stacks` over-reports.
- Popping from the first inner stack, which breaks LIFO order.
