# A Pointer That Outlives Its Frame

**Level:** junior
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A new Go developer is told "pointers are fast, values are copies" and starts returning pointers everywhere. The allocation profile disagrees.

## Task

Implement [retptr.go](retptr.go):

1. Return a pointer to a new int holding `v`.
2. Each call must return a distinct pointer.
3. Exactly one allocation per call.

Replace the stub body in [retptr.go](retptr.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  *New(7)
Output: 7
```

**Example 2:**

```
Input:  a, b := New(1), New(1)
Output: a != b
```

_Explanation:_ Every call gets its own int.

**Example 3:**

```
Input:  allocations per call
Output: 1
```

_Explanation:_ The int must escape to the heap.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Escape analysis** | The compiler decides stack or heap by asking whether the value outlives the frame. |
| 2 | **Returning &local** | Legal in Go — the value is simply moved to the heap. |
| 3 | **Pointer identity** | Distinct allocations mean distinct addresses. |

## Hint

In C this would be a dangling pointer. In Go it is an allocation.

## Validate

```bash
make verify
```
