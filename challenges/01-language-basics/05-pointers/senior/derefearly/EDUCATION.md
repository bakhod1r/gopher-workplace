# Nil checks precede dereferences

## Intuition

A dereference evaluates immediately; guarding after it is too late. The nil test must run before touching the pointee.

## Approach

1. The bug reads `head.Val` **before** the nil check, panicking on nil.
2. Move the nil check first, then dereference.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func FirstOr(head *Node, def int) int {
	if head == nil {
		return def
	}
	return head.Val
}
```

## Walkthrough

`FirstOr(nil, 7)` should return 7, but reading `head.Val` up front dereferences nil and panics. Checking `head == nil` first is safe.

## Pitfalls

- `head.Val` panics on nil regardless of a later check.
- Put `if head == nil` first.
