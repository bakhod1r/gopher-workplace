# The tortoise-and-hare technique

## Intuition

Two pointers at different speeds locate the middle (or detect cycles) in a single pass with O(1) extra space.

## Approach

1. Run `slow` by one and `fast` by two.
2. When `fast` reaches the end, `slow` sits at the middle.
3. One pass, no length count needed.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func Middle(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
```

## Walkthrough

For `1..5`: slow goes 1,2,3 while fast goes 1,3,5; fast stops, slow is at node 3.

## Pitfalls

- Guard both `fast != nil` and `fast.Next != nil` before advancing.
- The even-length midpoint convention depends on the loop condition.
