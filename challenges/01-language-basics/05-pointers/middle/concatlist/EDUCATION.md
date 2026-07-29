# Joining linked lists

## Intuition

Linking the first list's tail to the second's head concatenates them in O(len a); the empty first-list case returns the second directly.

## Approach

1. If `a` is nil, `b` is the whole result.
2. Walk to `a`'s tail and link `tail.Next = b`.
3. Return `a`.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func Concat(a, b *Node) *Node {
	if a == nil {
		return b
	}
	t := a
	for t.Next != nil {
		t = t.Next
	}
	t.Next = b
	return a
}
```

## Walkthrough

`Concat(1->2, 3->4)`: the tail of `a` is node 2; setting `2.Next = b` joins the lists into `1->2->3->4`.

## Pitfalls

- Handle a nil first list.
- Only a's tail's Next changes.
