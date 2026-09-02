# Generic Queue

## Intuition

Reslicing from the front is O(1) and keeps the code short; it does leave the popped element reachable by the backing array, which is a memory concern for long-lived queues rather than a correctness one.

## Approach

1. `Enqueue`: append.
2. `Dequeue`: guard empty, read index 0, reslice from 1.
3. `Len`: return `len(q.items)`.

## Solution

```go
func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	front := q.items[0]
	q.items = q.items[1:]
	return front, true
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}
```

## Walkthrough

`Enqueue(1); Enqueue(2); Dequeue()` returns `1`, leaving `items` as `[2]`.

## Pitfalls

- Reading the last element, which turns the queue into a stack.
- Using a value receiver, so the dequeue is not visible to the caller.
- Reslicing before reading the front element.
