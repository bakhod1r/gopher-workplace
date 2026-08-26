# Pop Stack

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

The expression evaluator pops operands when it encounters an operator.
`Pop` must handle the empty-stack case gracefully.

## Task

Implement `Pop` on `*Stack` in [popstk.go](popstk.go):

1. Remove and return the last element of `Items`.
2. Return `(0, false)` if the stack is empty.
3. Return `(value, true)` on success.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  s := Stack{Items: []int{1, 2, 3}}; s.Pop()
Output: (3, true), s.Items == [1, 2]
```

**Example 2:**

```
Input:  s := Stack{Items: []int{42}}; s.Pop()
Output: (42, true), s.Items == []
```

**Example 3:**

```
Input:  s := Stack{}; s.Pop()
Output: (0, false)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer receiver** | `Pop` shrinks the slice — caller must see the change. |
| 2 | **Multiple return values** | `(int, bool)` — the comma-ok idiom. |
| 3 | **Slice reslicing** | `s.Items = s.Items[:n-1]` shrinks the slice. |

## Hint

Check length first. Get the last element, reslice to `[:n-1]`, return value and
`true`.

## Validate

```bash
make verify
```
