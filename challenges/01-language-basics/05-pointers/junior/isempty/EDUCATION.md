# nil as the empty case

## Intuition

A nil pointer conventionally represents an empty linked structure; testing `== nil` needs no dereference.

## Approach

1. An empty list is represented by a nil head.
2. `return head == nil`.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func IsEmpty(head *Node) bool {
	return head == nil
}
```

## Walkthrough

`IsEmpty(nil)` → `true`. Any real `*Node`, even zero-valued, is non-nil → `false`.

## Pitfalls

- Comparing to nil is always safe; dereferencing nil is not.
- Many recursive list algorithms base-case on nil.
