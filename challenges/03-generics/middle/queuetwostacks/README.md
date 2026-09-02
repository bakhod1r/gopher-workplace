# Queue From Two Stacks

**Level:** middle  
**Topic:** 03-generics

## Context

A queue built on slice reslicing keeps the whole backing array alive. Two stacks avoid that while staying O(1) amortised.

## Task

Implement the stub(s) in [queuetwostacks.go](queuetwostacks.go):

1. Implement `Enqueue`, `Dequeue`, and `Len`.
2. Move elements from the inbound to the outbound stack only when the outbound stack is empty.
3. `Dequeue` returns the zero value and `false` when both stacks are empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Enqueue(1); Enqueue(2); Dequeue()
Output: 1, true
```

**Example 2:**

```
Input:  Dequeue() on empty
Output: 0, false
```

**Example 3:**

```
Input:  Enqueue(1); Dequeue(); Enqueue(2); Dequeue()
Output: 1 then 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Amortised cost** | An occasional expensive step can still average out to O(1) per operation. |
| 2 | **Reversal on transfer** | Copying the inbound stack backwards puts the oldest element on top. |
| 3 | **Transfer only when empty** | Refilling early would interleave old and new elements out of order. |

## Hint

Transfer only when `out` is empty, and copy `in` backwards.

## Validate

```bash
make verify
```
