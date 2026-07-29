# Mutating a struct's slice field

## Intuition

A pointer-receiver method appends to and reassigns the struct's slice field, keeping the growth visible to callers.

## Approach

1. The receiver is `*Stack` so `data` grows persistently.
2. `Push` appends; `Len` returns `len(s.data)`.

## Solution

```go
type Stack struct{ data []int }

func (s *Stack) Len() int { return len(s.data) }

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}
```

## Walkthrough

After two `Push` calls the backing slice holds two elements, so `Len()` returns 2. A pointer receiver keeps the appended header.

## Pitfalls

- A value receiver would append to a copy's field and lose it.
- Reassign the field with the append result.
