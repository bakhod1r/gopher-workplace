# Expression Evaluator

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

An expression tree is the output of every parser. Evaluating it is a recursive
method: each node asks its children for their values and combines them
according to its own kind.

## Task

Implement `Eval` on `*Expr` in [exprparser.go](exprparser.go):

1. A nil receiver evaluates to `0`.
2. `Num` returns `e.Val`.
3. `Add` returns `Left.Eval() + Right.Eval()`.
4. `Mul` returns `Left.Eval() * Right.Eval()`.
5. `Neg` returns `-Left.Eval()`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  {Kind: Num, Val: 5}
Output: 5
```

**Example 2:**

```
Input:  Mul(Add(2, 3), 4)
Output: 20
```

**Example 3:**

```
Input:  Neg(Mul(7, 2))
Output: -14
```

_Explanation:_ the tree already encodes precedence — evaluation just follows its shape.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Recursive method on a pointer type** | `e.Left.Eval()` is the same method, one level down. |
| 2 | **Nil receiver base case** | `var e *Expr; e.Eval()` must return 0 rather than panic. |
| 3 | **`switch` over an `iota` kind** | One arm per node type keeps the evaluator exhaustive and readable. |

## Hint

Put `if e == nil { return 0 }` first; it doubles as the base case for `Neg`,
whose `Right` is always nil, and for any malformed leaf.

## Validate

```bash
make verify
```
