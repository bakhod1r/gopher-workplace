# Ring-based list rotation

## Intuition

Closing the list into a ring and breaking it at the right offset rotates in O(n); the modulo handles k larger than the length.

## Approach

1. Find the length and link the tail to the head to form a ring.
2. Walk `len - k%len` steps to the new tail.
3. Break the ring there; the next node is the new head.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func Rotate(head *Node, k int) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	k = k % n
	if k == 0 {
		return head
	}
	tail.Next = head
	steps := n - k
	newTail := head
	for i := 1; i < steps; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}
```

## Walkthrough

`Rotate(1..5, 2)`: form a ring, step `5-2=3` nodes to node 3 (new tail), break after it → head becomes node 4: `4->5->1->2->3`.

## Pitfalls

- Reduce k with `k % len` (guard len 0).
- Breaking the ring one node early sets the new head.
