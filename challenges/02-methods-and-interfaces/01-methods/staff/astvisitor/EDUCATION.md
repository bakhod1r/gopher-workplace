# AST Visitor

## Intuition

Counting over a tree is a fold, and Go's usual way to write a fold without
returning values at every level is a pointer accumulator: one `int` owned by the
caller, incremented by whichever frames find a match.

## Approach

1. Handle nil so recursion is unconditional.
2. Test this node and update the accumulator.
3. Descend into both children with the same pointer.

## Solution

```go
func (n *Node) Visit(count *int) {
	if n == nil {
		return
	}
	if n.Type == "Ident" {
		*count++
	}
	n.Left.Visit(count)
	n.Right.Visit(count)
}
```

## Walkthrough

The root is a `BinOp`, so nothing is counted; it descends left to `Ident x`
(count 1), then right into another `BinOp`, whose left child `Ident y` brings
the count to 2. The `Num` node and every nil child contribute nothing. The
caller's `count` variable, addressed once at the top call, holds 2.

## Pitfalls

- **`*count++`.** Go parses this as `*(count++)`, which is not valid — `++` is a
  statement, not an expression, and pointer arithmetic does not exist. Use
  `(*count)++`.
- **Passing `count` by value.** Each frame then increments its own copy and the
  caller sees 0.
- **Returning a count instead.** Perfectly good design — `func (n *Node) Count()
  int` summing children — but it is not the signature this puzzle fixes.

## Why not a method value or closure?

A closure over a local `count` would avoid the pointer entirely, and that is what
the `visitorpatt` puzzle does. The pointer form shown here is what `go/ast`-style
walkers use when the visitor must also be usable from another package.
