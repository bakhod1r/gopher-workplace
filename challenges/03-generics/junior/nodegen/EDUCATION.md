# Linked List

## Intuition

Because the node's `Next` field is `*Node[T]` — the same instantiation — the structure is self-referential without any interface or `any` in sight.

## Approach

1. `Prepend`: return `&Node[T]{Value: v, Next: head}`.
2. `ToSlice`: walk from `head` until nil, appending each value.

## Solution

```go
func Prepend[T any](head *Node[T], v T) *Node[T] {
	return &Node[T]{Value: v, Next: head}
}

func ToSlice[T any](head *Node[T]) []T {
	out := make([]T, 0)
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Value)
	}
	return out
}
```

## Walkthrough

`ToSlice(Prepend(Prepend(nil, 2), 1))` walks the node holding `1` first, then the one holding `2`.

## Pitfalls

- Writing `Next *Node` without the type argument — that does not compile.
- Returning `nil` from `ToSlice(nil)` when an empty slice is expected.
- Mutating `head.Next` in `Prepend`, which breaks the older versions.
