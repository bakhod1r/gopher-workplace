# Pointer surgery on adjacent nodes

## Intuition

Swapping nodes rewires their Next links and promotes the second node to head — order the assignments so no link is lost.

## Approach

1. Guard lists with 0 or 1 node.
2. Let `second = head.Next`; relink `head.Next = second.Next` and `second.Next = head`.
3. Return `second` as the new head.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func SwapHead(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	second := head.Next
	head.Next = second.Next
	second.Next = head
	return second
}
```

## Walkthrough

`SwapHead(1->2->3)`: second is `2`; `1.Next` becomes `3`, `2.Next` becomes `1`, and `2` is returned → `2->1->3`.

## Pitfalls

- Save `second.Next` before overwriting.
- Return the new head (the old second node).
