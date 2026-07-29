# Node deletion by relinking

## Intuition

Removing a node re-points its predecessor's Next to its successor; deleting the head yields a new head.

## Approach

1. If the head matches, return `head.Next`.
2. Else walk with a `prev`; when `prev.Next.Val == v`, splice it out via `prev.Next = prev.Next.Next`.
3. Return the (possibly new) head.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func Delete(head *Node, v int) *Node {
	if head == nil {
		return nil
	}
	if head.Val == v {
		return head.Next
	}
	prev := head
	for prev.Next != nil {
		if prev.Next.Val == v {
			prev.Next = prev.Next.Next
			break
		}
		prev = prev.Next
	}
	return head
}
```

## Walkthrough

`Delete(1->2->3, 2)`: head is `1`, not a match; at node 1 the next is `2`, so `1.Next = 3`, yielding `1->3`.

## Pitfalls

- Handle the head separately — it has no predecessor.
- Only the first match is removed.
