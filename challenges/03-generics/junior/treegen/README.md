# Binary Search Tree

**Level:** junior  
**Topic:** 03-generics

## Context

An autocomplete index keeps terms sorted as they arrive, without re-sorting the whole set on every insert.

## Task

Implement the stub(s) in [treegen.go](treegen.go):

1. Implement `Insert`, returning the root of the tree with `v` added.
2. Duplicates are ignored.
3. Implement `InOrder`, returning the values in ascending order; `InOrder(nil)` returns an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Insert(nil, 2) then Insert(root, 1); InOrder(root)
Output: []int{1, 2}
```

**Example 2:**

```
Input:  insert 2,1,3; InOrder
Output: []int{1, 2, 3}
```

**Example 3:**

```
Input:  InOrder(nil)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Recursive generic structures** | `TreeNode[T]` holds two `*TreeNode[T]` children. |
| 2 | **Ordered constraint on a structure** | `cmp.Ordered` is what lets the tree place values left or right. |
| 3 | **Returning the root** | Insert returns a node so the nil case can create one. |

## Hint

Return the (possibly new) root from `Insert` and assign it back into the child field.

## Validate

```bash
make verify
```
