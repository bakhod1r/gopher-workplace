# Floyd's cycle detection

## Intuition

Two pointers at speeds 1 and 2 must meet inside a cycle and never meet in an acyclic list — O(1) space.

## Approach

1. Move `slow` by 1 and `fast` by 2.
2. If they ever meet, there is a cycle.
3. If `fast` reaches nil, the list is acyclic.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func HasCycle(head *Node) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
```

## Walkthrough

On a cyclic list fast laps slow and they collide → true; on an acyclic list fast walks off the end → false.

## Pitfalls

- Guard `fast` and `fast.Next` before the double step.
- Reaching nil proves no cycle.
