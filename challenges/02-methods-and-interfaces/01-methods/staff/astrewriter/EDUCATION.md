# AST Rewriter

## Intuition

A rewriter is a visitor that is allowed to write. The traversal is identical to
a read-only walk; the only difference is that the node method holds a pointer,
so an assignment sticks. That is why `Rewrite` must be on `*Node` and not on
`Node`.

## Approach

1. Stop at nil.
2. Apply the rewrite rule to this node if it matches.
3. Recurse into both children.

## Solution

```go
func (n *Node) Rewrite() {
	if n == nil {
		return
	}
	if n.Type == "Ident" && n.Val == "foo" {
		n.Val = "bar"
	}
	n.Left.Rewrite()
	n.Right.Rewrite()
}
```

## Walkthrough

The test tree is a single `Ident` node holding `"foo"`. The guard passes, so
`n.Val` is assigned `"bar"` through the receiver pointer — the caller's `root`
sees the change. Both children are nil, and calling a pointer-receiver method
on a nil pointer is legal, so the recursion terminates on the first guard.

## Pitfalls

- **Value receiver.** `func (n Node) Rewrite()` compiles and mutates a copy;
  the test still sees `"foo"`.
- **Skipping the type check.** A `BinOp` node that happened to carry `Val ==
  "foo"` would be rewritten too.
- **Guarding at the call site** (`if n.Left != nil`) instead of at the top —
  correct, but it duplicates the check.

## Real Go AST rewriting

`go/ast` uses `ast.Inspect` or `astutil.Apply` for exactly this shape, with a
node interface and many concrete types instead of one struct with a `Type`
string. The traversal contract — visit, then descend, mutate through pointers —
is the same.
