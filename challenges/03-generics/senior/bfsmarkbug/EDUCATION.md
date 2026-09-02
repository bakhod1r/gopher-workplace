# Breadth-First With Duplicates

## Intuition

Between enqueueing a node and popping it, another neighbour can enqueue it again, because it is not marked until it is dequeued.

## Approach

1. Mark the start as seen and enqueue it.
2. Pop, emit, then for each unseen neighbour mark it and enqueue it.

## Solution

```go
func BFS[T comparable](adj map[T][]T, start T) []T {
	seen := map[T]bool{start: true}
	queue := []T{start}
	out := make([]T, 0)
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		out = append(out, n)
		for _, m := range adj[n] {
			if !seen[m] {
				seen[m] = true
				queue = append(queue, m)
			}
		}
	}
	return out
}
```

## Walkthrough

In the diamond, both `b` and `c` see `d` as unseen and enqueue it, so `d` is emitted twice.

## Pitfalls

- Marking only on dequeue — correct output, quadratic queue growth, and duplicates.
- Forgetting to mark the start, which lets a cycle revisit it.
