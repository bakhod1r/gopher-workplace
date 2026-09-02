# AST Visitor

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Before you can rewrite a tree you usually want to measure it. This visitor walks
an expression tree and counts identifier nodes, accumulating through a pointer
the caller owns.

## Task

Implement `Visit` on `*Node` in [astvisitor.go](astvisitor.go):

1. Return if the receiver is nil.
2. If `Type == "Ident"`, increment `*count`.
3. Recurse into `Left` and `Right`, passing `count` along.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a single Ident node
Output: count == 1
```

**Example 2:**

```
Input:  BinOp(Ident x, BinOp(Ident y, Num))
Output: count == 2
```

**Example 3:**

```
Input:  a tree with no Idents
Output: count == 0
```

_Explanation:_ non-matching node types are traversed but not counted.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Accumulator by pointer** | `*int` lets every recursive frame add to one shared total. |
| 2 | **Nil receiver base case** | Children are recursed unconditionally because nil is handled. |
| 3 | **Type-tagged nodes** | A `Type` string stands in for Go's real `ast.Node` interface. |

## Hint

`*count++` does not parse the way you might hope — write `*count = *count + 1`
or, more idiomatically, `(*count)++`.

## Validate

```bash
make verify
```
