# The Wraparound That Goes Negative

**Level:** staff  
**Topic:** 03-generics

## Context

A ring-buffer deque works perfectly as a queue. The first time a caller pushes to the front instead of the back, the process dies with an index out of range.

## Task

Fix the single planted bug in [dequewrapmodbug.go](dequewrapmodbug.go):

1. Find and fix the single bug so moving the head backwards wraps to the end of the buffer.
2. The existing back-of-the-deque behaviour must not change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  PushBack(2); PushFront(1); At(0)
Output: 1
```

**Example 2:**

```
Input:  PushBack(2); PushFront(1); At(1)
Output: 2
```

**Example 3:**

```
Input:  PushFront(1); PopFront()
Output: 1, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Go's % is remainder, not modulo** | `-1 % 3` is `-1` in Go, not `2`: the result carries the sign of the dividend. |
| 2 | **Ring arithmetic** | Add the modulus before taking the remainder to keep the value non-negative. |
| 3 | **Structural invariants** | Every operation must restore what the type promises about itself. |

## Hint

What does `(0 - 1) % 3` evaluate to in Go?

## Validate

```bash
make verify
```
