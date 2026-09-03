# Two Stacks, One Wrong Transfer

## Intuition

Transferring unconditionally stacks newer elements on top of older ones already waiting in `out`, so the newest is served first.

## Approach

1. Transfer only when `out` is empty.
2. Then pop from the top of `out`.

## Solution

```go
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			v := q.in[len(q.in)-1]
			q.in = q.in[:len(q.in)-1]
			q.out = append(q.out, v)
		}
	}
	if len(q.out) == 0 {
		var zero T
		return zero, false
	}
	v := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return v, true
}

func (q *Queue[T]) Enqueue(v T) {
	q.in = append(q.in, v)
}
```

## Walkthrough

`Enq 1,2; Deq` gives 1 and leaves `out = [2]`. `Enq 3` then transfers 3 on top, so the next `Deq` returns 3 before 2.

## Pitfalls

- Transferring on every enqueue instead — correct, but it loses the amortised bound.
- Popping from the front of `out`, which is O(n) per operation.
