# Shallow Tree Copy Shares Nodes

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Reusing `t.Left`/`t.Right` shares the original subtrees; a deep copy must
recursively copy the children so the two trees share no nodes.

## Task

Fix [shallowcopy.go](shallowcopy.go) to deep-copy the children.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  c := Copy(tree); mutate c.Left
Output: original tree unchanged
```

**Example 2:**

```
Input:  Copy(nil)
Output: nil
```

**Example 3:**

```
Input:  Copy(single node)
Output: independent single node
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Deep vs shallow** | Reusing children shares memory. |
| 2 | **Recursive copy** | `Copy(t.Left)`, `Copy(t.Right)`. |
| 3 | **Independence** | Mutating the copy must not touch the original. |

## Hint

Recurse: `return &Tree{Val: t.Val, Left: Copy(t.Left), Right: Copy(t.Right)}`.

## Validate

```bash
make verify
```
