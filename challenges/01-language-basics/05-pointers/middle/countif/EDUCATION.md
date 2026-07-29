# Predicate over a pointer structure

## Intuition

Threading a `func(int) bool` through a list traversal generalises counting/filtering to any condition.

## Approach

1. Walk the list.
2. Apply `pred` to each `Val`, counting the trues.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func CountIf(head *Node, pred func(int) bool) int {
	c := 0
	for n := head; n != nil; n = n.Next {
		if pred(n.Val) {
			c++
		}
	}
	return c
}
```

## Walkthrough

`CountIf(1->2->3->4, even)` tests each value; 2 and 4 pass, so the count is 2.

## Pitfalls

- Stop at nil; test `n.Val` each step.
- The predicate decouples the condition from the walk.
