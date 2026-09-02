# Expression Evaluator

## Intuition

A tree-walking interpreter is the simplest evaluator there is: the tree already
carries precedence and associativity, so `Eval` never has to think about
operators — only about what this one node means once its children are numbers.

## Approach

1. Handle nil, which is both a guard and the base case.
2. Switch on the node kind.
3. Recurse for operands, combine, return.

## Solution

```go
func (e *Expr) Eval() int {
	if e == nil {
		return 0
	}
	switch e.Kind {
	case Num:
		return e.Val
	case Add:
		return e.Left.Eval() + e.Right.Eval()
	case Mul:
		return e.Left.Eval() * e.Right.Eval()
	case Neg:
		return -e.Left.Eval()
	}
	return 0
}
```

## Walkthrough

For `(2 + 3) * 4` the root is `Mul`. It asks its left child, an `Add`, which
asks its two `Num` leaves and returns 5. The right leaf returns 4. The root
multiplies: 20.

For `-(7 * 2)` the `Neg` node evaluates only `Left`; its nil `Right` is never
touched — and even if the code did touch it, the nil guard returns 0 rather than
panicking.

The evaluation order is depth-first, left to right, which is exactly the order
the recursive calls appear in the source.

## Pitfalls

- **Missing nil guard.** `Neg` (and any malformed tree) dereferences nil and
  panics.
- **Handling `Neg` as `Left - Right`.** Works by accident because `Right` is nil
  and evaluates to 0, but it hides the intent and breaks if `Right` is ever set.
- **Value receiver.** `func (e Expr) Eval()` cannot be called on a nil `*Expr` at
  all — the automatic dereference panics before the body runs.
- **A `default` that panics.** Tempting for exhaustiveness, but it turns an
  unknown kind into a crash; returning 0 keeps the evaluator total.

## Kind tags vs interfaces

Go's own `go/ast` uses an interface with one type per node and a type switch at
each visit. The `Kind` field here is the compact alternative: fewer types, but
the compiler can no longer tell you a case is missing. Both shapes appear in
real Go code; the interface version scales better once nodes carry different
fields.
