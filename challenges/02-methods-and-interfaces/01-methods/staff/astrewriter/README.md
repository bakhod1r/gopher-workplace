# AST Rewriter

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Source-to-source tools walk a syntax tree and edit nodes in place. This is the
smallest possible version: rename every identifier `foo` to `bar`, everywhere in
the tree.

## Task

Implement `Rewrite` on `*Node` in [astrewriter.go](astrewriter.go):

1. Return if the receiver is nil.
2. If `Type == "Ident"` and `Val == "foo"`, set `Val = "bar"`.
3. Recurse into `Left` and `Right`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  {Type: "Ident", Val: "foo"}
Output: Val == "bar"
```

**Example 2:**

```
Input:  {Type: "Ident", Val: "baz"}
Output: Val == "baz"  (untouched)
```

**Example 3:**

```
Input:  a BinOp whose Left is Ident "foo"
Output: only the Left child is renamed
```

_Explanation:_ the type check keeps the rewrite from hitting operator nodes.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **In-place mutation via pointer receiver** | The tree is edited where it stands; nothing is rebuilt. |
| 2 | **Nil receiver as the base case** | Lets both recursive calls be unconditional. |
| 3 | **Guarded rewrite** | Two conditions must hold before writing — type *and* value. |

## Hint

Mutation before recursion or after makes no difference here, since a node's own
value and its children are independent. The nil guard is what keeps the two
recursive calls safe.

## Validate

```bash
make verify
```
