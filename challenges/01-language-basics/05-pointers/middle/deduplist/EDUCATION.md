# In-place list compaction

## Intuition

Relinking past equal successors removes duplicates without allocation; advance only when you didn't skip.

## Approach

1. Walk `cur`.
2. While `cur.Next` has the same value, unlink it via `cur.Next = cur.Next.Next`.
3. Otherwise advance `cur`.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func Dedup(head *Node) *Node {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
```

## Walkthrough

`Dedup(1->1->2->3->3)`: node 1 skips the duplicate 1; node 3 skips the duplicate 3, leaving `1->2->3`.

## Pitfalls

- Don't advance `cur` when you just skipped — there may be more dups.
- Guard `cur.Next != nil` before comparing.
