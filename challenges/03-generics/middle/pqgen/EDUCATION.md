# Priority Queue

## Intuition

The tie rule is the whole design: without it, a burst of same-priority work would come out in an arbitrary order and old tickets could wait forever.

## Approach

1. `Push`: find the first entry with a strictly greater priority, and insert before it.
2. `Pop`: take the front entry.
3. `Len`: report the count.

## Solution

```go
func (q *PQ[T]) Push(v T, priority int) {
	e := entry[T]{value: v, priority: priority, seq: q.seq}
	q.seq++
	i := 0
	for i < len(q.items) && !(e.priority < q.items[i].priority) {
		i++
	}
	q.items = append(q.items, entry[T]{})
	copy(q.items[i+1:], q.items[i:])
	q.items[i] = e
}

func (q *PQ[T]) Pop() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	e := q.items[0]
	q.items = q.items[1:]
	return e.value, true
}

func (q *PQ[T]) Len() int {
	return len(q.items)
}
```

## Walkthrough

`Push(a,1); Push(b,1)` places `b` after `a`, so `Pop` returns `a` first.

## Pitfalls

- Inserting before equal priorities, which reverses arrival order.
- Popping from the back, which returns the lowest-priority item.
- Forgetting to shift the tail when inserting in the middle.
