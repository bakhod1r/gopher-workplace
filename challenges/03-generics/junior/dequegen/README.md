# Deque

**Level:** junior  
**Topic:** 03-generics

## Context

A task scheduler adds urgent work at the front and routine work at the back, always taking from the front.

## Task

Implement the stub(s) in [dequegen.go](dequegen.go):

1. Implement `PushFront`, `PushBack`, and `PopFront`.
2. `PopFront` returns the zero value and `false` when empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  PushBack(1); PushFront(0); PopFront()
Output: 0, true
```

**Example 2:**

```
Input:  PushBack(1); PushBack(2); PopFront()
Output: 1, true
```

**Example 3:**

```
Input:  empty PopFront()
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Both ends of a slice** | Appending at the back is cheap; inserting at the front costs a copy. |
| 2 | **Variadic splice** | `append([]T{v}, d.items...)` builds the new front-loaded slice in one call. |
| 3 | **Pointer receivers mutate** | Reused from methods: a value receiver gets a copy, so mutation needs `*T`. |

## Hint

`append([]T{v}, d.items...)` is the front insert.

## Validate

```bash
make verify
```
