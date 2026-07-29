# Returning the new head

## Intuition

Deleting the head has no predecessor to relink, so the function must return the successor for the caller to adopt.

## Approach

1. Removing the first node means the new head is `head.Next`.
2. The bug returns `head` (the node meant to be removed).
3. Return `head.Next` instead.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func RemoveFirst(head *Node) *Node {
	if head == nil {
		return nil
	}
	return head.Next
}
```

## Walkthrough

`RemoveFirst(1->2->3)` should drop node 1 and return node 2. Returning `head` keeps node 1; `head.Next` gives `2->3`.

## Pitfalls

- Returning the original `head` keeps the deleted node; RETURN head.Next.
- The caller must use the returned head.
