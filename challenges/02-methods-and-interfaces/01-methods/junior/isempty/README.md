# Is Empty

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

An undo system tracks operations in a stack. Before popping, it checks whether
the stack is empty.

## Task

Implement `IsEmpty` on `Stack` in [isempty.go](isempty.go):

1. Return `true` if the stack has no items.
2. Works for both `nil` and empty slices.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Stack{}.IsEmpty()
Output: true
```

**Example 2:**

```
Input:  Stack{items: []int{1}}.IsEmpty()
Output: false
```

**Example 3:**

```
Input:  Stack{items: []int{}}.IsEmpty()
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receiver** | Read-only check on the stack's length. |
| 2 | **len on nil slice** | `len(nil)` returns 0 in Go — no special case needed. |

## Hint

`return len(s.items) == 0` — works for both nil and empty slices.

## Validate

```bash
make verify
```
