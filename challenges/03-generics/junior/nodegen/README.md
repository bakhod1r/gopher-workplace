# Linked List

**Level:** junior  
**Topic:** 03-generics

## Context

An immutable history is built by chaining nodes: each new entry points at the previous head, so old versions stay valid.

## Task

Implement the stub(s) in [nodegen.go](nodegen.go):

1. Implement `Prepend`, returning a new node whose `Next` is the old head.
2. Implement `ToSlice`, returning the values in order from head to tail.
3. `ToSlice(nil)` returns an empty, non-nil slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Prepend(nil, 1)
Output: list [1]
```

**Example 2:**

```
Input:  ToSlice(Prepend(Prepend(nil, 2), 1))
Output: []int{1, 2}
```

**Example 3:**

```
Input:  ToSlice(nil)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Recursive generic types** | `Node[T]` has a `*Node[T]` field — a type may refer to its own instantiation. |
| 2 | **Nil as the empty list** | A nil `*Node[T]` is a perfectly good empty list. |
| 3 | **Pointer walking** | Reused from language basics: `for n := head; n != nil; n = n.Next`. |

## Hint

A nil head is the empty list — no special case needed beyond the loop condition.

## Validate

```bash
make verify
```
