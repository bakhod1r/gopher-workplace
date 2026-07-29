# Guarding container operations

## Intuition

Removing from an empty pointer-backed container must check for nil before dereferencing the head.

## Approach

1. Popping an empty queue dereferences a nil head → panic.
2. Guard: `if q.head == nil { return 0, false }` before reading.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

type Queue struct{ head *Node }

func (q *Queue) Pop() (int, bool) {
	if q.head == nil {
		return 0, false
	}
	v := q.head.Val
	q.head = q.head.Next
	return v, true
}
```

## Walkthrough

With no guard, `q.head.Val` on an empty queue panics. The nil check returns the zero value and `false` safely.

## Pitfalls

- `q.head.Val` on an empty queue panics.
- Return the zero/not-ok pair before dereferencing.
