# Nil Receivers

## Intuition

In Go, you can call a method on a nil pointer — the call doesn't panic as long
as the method handles the nil case. This is powerful for linked structures.

## Approach

1. If `n == nil`, return 0.
2. Walk `Next`, incrementing a counter.

## Solution

```go
func (n *Node) Len() int {
	count := 0
	for cur := n; cur != nil; cur = cur.Next {
		count++
	}
	return count
}
```

## Walkthrough

For `&Node{1, &Node{2, &Node{3, nil}}}`:
- cur=node1 → count=1
- cur=node2 → count=2
- cur=node3 → count=3
- cur=nil → stop.

## Pitfalls

- Forgetting the nil check: `n.Next` on nil panics.
- Using a value receiver `(n Node)` — cannot handle nil.
