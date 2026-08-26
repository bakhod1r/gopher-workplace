# Peek

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A syntax highlighter peeks at the top token without consuming it to decide
the colour.

## Task

Implement `Peek` on `Stack` in [peek.go](peek.go):

1. Return the top (last) element and `true`.
2. Return `(0, false)` if the stack is empty.
3. Do **not** modify the stack — value receiver.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Stack{Items: []int{1, 2, 3}}.Peek()
Output: (3, true)
```

**Example 2:**

```
Input:  Stack{Items: []int{42}}.Peek()
Output: (42, true)
```

**Example 3:**

```
Input:  Stack{}.Peek()
Output: (0, false)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receiver** | `Peek` is read-only — no mutation. |
| 2 | **Comma-ok idiom** | `(int, bool)` return signals success or empty. |

## Hint

Check length, then return `s.Items[n-1]` and `true`.

## Validate

```bash
make verify
```
