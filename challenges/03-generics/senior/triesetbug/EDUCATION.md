# The Trie That Marks The Wrong Node

## Intuition

Marking before the descent flags the root, so the empty sequence looks inserted while the real endpoint is never marked.

## Approach

1. Start at the root.
2. Walk or create a child per element.
3. Mark the node you finish on.

## Solution

```go
func (t *Trie[T]) Insert(seq []T) {
	n := t.root()
	for _, e := range seq {
		next, ok := n.kids[e]
		if !ok {
			next = &node[T]{kids: make(map[T]*node[T])}
			n.kids[e] = next
		}
		n = next
	}
	n.end = true
}

func (t *Trie[T]) Contains(seq []T) bool {
	n := t.root()
	for _, e := range seq {
		next, ok := n.kids[e]
		if !ok {
			return false
		}
		n = next
	}
	return n.end
}
```

## Walkthrough

After `Insert([a,b])` the root carries `end`, so `Contains([])` is true and `Contains([a,b])` is false.

## Pitfalls

- Marking inside the loop, which flags every prefix.
- Forgetting to allocate the child map, which panics on the first write.
