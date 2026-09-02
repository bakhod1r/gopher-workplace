# Nested Stacks

**Level:** middle  
**Topic:** 03-generics

## Context

A staging area holds items in trays of fixed size. When a tray fills, a new one starts; when a tray empties, it is taken away.

## Task

Implement the stub(s) in [stackofstacksgen.go](stackofstacksgen.go):

1. Implement `Push`, `Pop`, and `Stacks`.
2. A new inner stack starts only when the current one is full.
3. An inner stack that becomes empty is removed, so `Stacks` reflects reality.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  cap 2: Push x3; Stacks()
Output: 2
```

**Example 2:**

```
Input:  cap 2: Push x3; Pop()
Output: the third item
```

**Example 3:**

```
Input:  Pop() on empty
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Slices of slices** | `[][]T` with a capacity rule per inner slice. |
| 2 | **Two invariants** | No inner stack exceeds the capacity, and no empty inner stack is retained. |
| 3 | **Overall LIFO** | Pushes and pops always touch the last inner stack, so the whole structure stays LIFO. |

## Hint

Start a new inner stack only when the last one is full; drop it as soon as it empties.

## Validate

```bash
make verify
```
