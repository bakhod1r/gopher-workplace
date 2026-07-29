# Guarding multi-step pointer advances

## Intuition

Moving two steps requires both the current and next pointers to be non-nil; a single guard leaves the second hop unsafe.

## Approach

1. `fast = fast.Next.Next` needs both `fast` and `fast.Next` non-nil.
2. The bug checks only `fast != nil`, so `fast.Next.Next` panics at the end of an odd list.
3. Guard with `fast != nil && fast.Next != nil`.

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

On an acyclic odd-length list, `fast` reaches the last node where `fast.Next` is nil; the full guard stops the loop instead of dereferencing nil.

## Pitfalls

- `fast.Next.Next` needs `fast.Next != nil` too.
- Check every dereference in a compound advance.
