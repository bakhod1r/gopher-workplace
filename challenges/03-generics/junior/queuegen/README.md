# Generic Queue

**Level:** junior  
**Topic:** 03-generics

## Context

A job runner processes tasks in arrival order. The same queue type serves both string job IDs and struct payloads.

## Task

Implement the stub(s) in [queuegen.go](queuegen.go):

1. Implement `Enqueue`, adding to the back.
2. Implement `Dequeue`, removing and returning the front element with `true`, or the zero value and `false` when empty.
3. Implement `Len`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Enqueue(1); Enqueue(2); Dequeue()
Output: 1, true
```

**Example 2:**

```
Input:  empty queue Dequeue()
Output: 0, false
```

**Example 3:**

```
Input:  Enqueue(1); Len()
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Generic types** | `type Stack[T any] struct { ... }` parameterises the type itself, not just a function. |
| 2 | **Receivers repeat the parameter** | A method on `Stack[T]` writes the receiver as `(s *Stack[T])` — the parameter comes along. |
| 3 | **FIFO versus LIFO** | The only difference from a stack is which end `Dequeue` reads. |

## Hint

Take from the front with `q.items[0]`, then reslice with `q.items[1:]`.

## Validate

```bash
make verify
```
