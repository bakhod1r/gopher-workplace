# Generic Stack

**Level:** junior  
**Topic:** 03-generics

## Context

An undo feature needs a last-in-first-out store. Two screens use it with different element types.

## Task

Implement the stub(s) in [stackgen.go](stackgen.go):

1. Implement `Push`, adding an element to the top.
2. Implement `Pop`, removing and returning the top element with `true`, or the zero value and `false` when empty.
3. Implement `Len`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Push(1); Push(2); Pop()
Output: 2, true
```

**Example 2:**

```
Input:  empty stack Pop()
Output: 0, false
```

**Example 3:**

```
Input:  Push(1); Len()
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Generic types** | `type Stack[T any] struct { ... }` parameterises the type itself, not just a function. |
| 2 | **Receivers repeat the parameter** | A method on `Stack[T]` writes the receiver as `(s *Stack[T])` — the parameter comes along. |
| 3 | **Pointer receivers mutate** | Reused from methods: a value receiver gets a copy, so mutation needs `*T`. |

## Hint

The zero value `Stack[int]{}` must already be usable — no constructor required.

## Validate

```bash
make verify
```
