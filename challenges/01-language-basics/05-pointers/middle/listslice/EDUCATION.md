# Materialising a list

## Intuition

Traversing and appending converts a pointer chain to a contiguous slice.

## Approach

1. Walk the list.
2. Append each `Val` to a growing slice.
3. Return it.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func ToSlice(head *Node) []int {
	var out []int
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}
```

## Walkthrough

`ToSlice(1->2->3)` appends 1, 2, 3 in order, producing `[1 2 3]`.

## Pitfalls

- Return empty (len 0) for nil head.
- One pass, append each value.
