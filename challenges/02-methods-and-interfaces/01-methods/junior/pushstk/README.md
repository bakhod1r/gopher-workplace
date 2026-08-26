# Push Stack

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

An expression evaluator uses a stack. `Push` adds operands as they are parsed.

## Task

Implement `Push` on `*Stack` in [pushstk.go](pushstk.go):

1. Append `v` to `Items`.
2. Pointer receiver — the slice header must be updated for the caller.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  s := Stack{}; s.Push(5)
Output: s.Items == [5]
```

**Example 2:**

```
Input:  s := Stack{}; s.Push(5); s.Push(3)
Output: s.Items == [5, 3]
```

**Example 3:**

```
Input:  s := Stack{Items: []int{1, 2}}; s.Push(3)
Output: s.Items == [1, 2, 3]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer receiver** | `append` may reallocate — the caller needs the new slice header. |
| 2 | **append** | `s.Items = append(s.Items, v)` grows the slice. |

## Hint

`s.Items = append(s.Items, v)` — pointer receiver ensures the new header
propagates.

## Validate

```bash
make verify
```
