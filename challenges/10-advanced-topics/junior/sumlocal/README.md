# A Function That Touches No Heap

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A hot metrics loop calls a sum helper millions of times. The helper looks trivial, yet the allocation profile blames it.

## Task

Implement [sumlocal.go](sumlocal.go):

1. Return the sum of the elements of `s`.
2. The function must make zero allocations.
3. A nil input sums to 0.

Replace the stub body in [sumlocal.go](sumlocal.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Sum([]int{1,2,3})
Output: 6
```

**Example 2:**

```
Input:  Sum(nil)
Output: 0
```

_Explanation:_ The empty sum.

**Example 3:**

```
Input:  Sum([]int{-3,3})
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Stack vs heap** | A local that never outlives the call stays on the stack. |
| 2 | **Reading a slice** | Ranging a slice copies elements into a loop variable — still no heap. |
| 3 | **AllocsPerRun** | Zero allocations is a testable property, not a hope. |

## Hint

An `int` accumulator. Nothing else.

## Validate

```bash
make verify
```
