# Priority Queue

**Level:** middle  
**Topic:** 03-generics

## Context

A support desk works tickets by severity, and two tickets of the same severity must be handled oldest first.

## Task

Implement the stub(s) in [pqgen.go](pqgen.go):

1. Implement `Push`, `Pop`, and `Len`.
2. Lower priority numbers come out first.
3. Items with equal priority must come out in the order they were pushed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Push(a,2); Push(b,1); Pop()
Output: b
```

**Example 2:**

```
Input:  Push(a,1); Push(b,1); Pop()
Output: a (FIFO within a priority)
```

**Example 3:**

```
Input:  Pop() on empty
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Stability in a queue** | Insertion order must break priority ties, or the queue starves older work. |
| 2 | **Insert position** | Advancing while the stored priority is not greater keeps equal items behind. |
| 3 | **Generic types** | The type parameter belongs to the type; methods reuse it, never add to it. |

## Hint

Insert *after* every item of equal or lower priority number — that is what makes ties FIFO.

## Validate

```bash
make verify
```
