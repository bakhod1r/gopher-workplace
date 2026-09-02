# Breadth-First Order

## Intuition

Marking at dequeue time lets a node be queued twice before it is first visited, which duplicates it in the output — the classic BFS bug.

## Approach

1. Seed the queue and the seen set with `start`.
2. Pop from the front, emit it, and enqueue every unseen neighbour.

## Solution

```go
func BFS[K comparable](edges map[K][]K, start K) []K {
	out := make([]K, 0)
	seen := map[K]bool{start: true}
	queue := []K{start}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		out = append(out, cur)
		for _, n := range edges[cur] {
			if seen[n] {
				continue
			}
			seen[n] = true
			queue = append(queue, n)
		}
	}
	return out
}
```

## Walkthrough

`BFS({a:[a]}, a)` enqueues `a` once; the self-edge is skipped because `a` is already marked.

## Pitfalls

- Marking at dequeue time, which duplicates nodes in cyclic graphs.
- Popping from the back, which produces depth-first order.
- Omitting `start` from the result.
