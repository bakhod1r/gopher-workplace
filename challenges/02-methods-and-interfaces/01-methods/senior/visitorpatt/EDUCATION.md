# Visitor Pattern

## Intuition

Two things change at different rates: the tree's shape (rarely) and what you
want to compute over it (constantly). Visitor freezes the traversal in the
structure's own method and lets the computation arrive from outside — sum,
print, validate, all with the same `Accept`.

## Approach

1. Handle the nil node first, so callers never need a guard.
2. Visit the current value.
3. Recurse into both children.

## Solution

```go
func (n *Node) Accept(visitor func(int)) {
	if n == nil {
		return
	}
	visitor(n.Val)
	n.Left.Accept(visitor)
	n.Right.Accept(visitor)
}
```

## Walkthrough

`root.Accept` visits 1, then descends left. That node visits 2, then calls
`Accept` on its nil `Left` and nil `Right` — legal, because a method with a
pointer receiver can run with `n == nil`; only *dereferencing* nil panics, and
the guard returns before any field access. The right subtree contributes 3, so
`sum == 6`.

## Pitfalls

- **Guarding at the call site instead.** `if n.Left != nil { n.Left.Accept(...) }`
  works but duplicates the check at every recursion point.
- **Dereferencing before the nil check.** `visitor(n.Val)` first panics on the
  leaf children.
- **Visiting after both recursions.** That is post-order — a different traversal,
  and a different result for order-sensitive visitors.

## Why nil-receiver methods work

A method call on a pointer is a function call with the pointer as the first
argument. Passing nil is fine; the panic comes only when the body reads through
it. Go's own `container` and tree packages lean on this to keep recursive code
short.
