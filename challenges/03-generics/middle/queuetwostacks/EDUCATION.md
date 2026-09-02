# Queue From Two Stacks

## Intuition

Each element moves between stacks at most once, so `n` operations cost O(n) in total — amortised constant time even though one dequeue may be expensive.

## Approach

1. `Enqueue`: push onto `in`.
2. `Dequeue`: when `out` is empty, move everything from `in` in reverse, then pop `out`.
3. `Len`: sum both lengths.

## Solution

```go
func (q *SQueue[T]) Enqueue(v T) {
	q.in = append(q.in, v)
}

func (q *SQueue[T]) Dequeue() (T, bool) {
	if len(q.out) == 0 {
		for i := len(q.in) - 1; i >= 0; i-- {
			q.out = append(q.out, q.in[i])
		}
		q.in = q.in[:0]
	}
	if len(q.out) == 0 {
		var zero T
		return zero, false
	}
	v := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return v, true
}

func (q *SQueue[T]) Len() int {
	return len(q.in) + len(q.out)
}
```

## Walkthrough

`Enqueue(1); Enqueue(2); Dequeue()` moves `[1 2]` into `out` as `[2 1]`, so popping yields `1`.

## Pitfalls

- Transferring on every dequeue, which reorders elements.
- Copying `in` forwards, which turns the queue into a stack.
- Forgetting to clear `in` after the transfer, duplicating elements.
