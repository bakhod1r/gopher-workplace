# Recursion over pointer structures

## Intuition

Linked lists and trees recurse naturally: base-case on nil, combine the current node with the recursive result.

## Approach

1. Base case: a nil list sums to 0.
2. Otherwise return `head.Val + SumList(head.Next)`.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func SumList(head *Node) int {
	if head == nil {
		return 0
	}
	return head.Val + SumList(head.Next)
}
```

## Walkthrough

`SumList(1->2->3->4)` unfolds to `1 + (2 + (3 + (4 + 0)))` = `10`.

## Pitfalls

- Base-case nil BEFORE dereferencing.
- Each call handles one node plus the rest.
