# Iterative list reversal

## Intuition

Walking the list while re-pointing each node's Next to the previous node reverses it in O(n) time and O(1) space.

## Approach

1. Keep a `prev`, initially nil.
2. For each node, save `next`, point `cur.Next` back to `prev`, advance both.
3. When `cur` is nil, `prev` is the new head.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func Reverse(head *Node) *Node {
	var prev *Node
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
```

## Walkthrough

For `1->2->3`: node 1's Next flips to nil, node 2's to 1, node 3's to 2; the loop ends with `prev == node 3`, so the head is `3->2->1`.

## Pitfalls

- Save `cur.Next` before overwriting it, or you lose the rest.
- The final `prev` is the new head.
