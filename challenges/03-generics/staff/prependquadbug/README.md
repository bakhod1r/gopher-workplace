# Reverse That Costs A Square

**Level:** staff  
**Topic:** 03-generics

## Context

A replay tool reverses an event log before rendering it. It is instant on a developer's 200-event fixture and times out on a production log of a hundred thousand events.

## Task

Fix the single planted bug in [prependquadbug.go](prependquadbug.go):

1. Find and fix the single bug so reversal is linear in the length of the input.
2. The returned order and contents must not change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Reversed([]int{1,2,3})
Output: [3 2 1]
```

**Example 2:**

```
Input:  Reversed([]string{})
Output: []
```

**Example 3:**

```
Input:  Reversed of 120000 elements
Output: well under the time budget
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Asymptotic cost** | A per-operation cost that grows with the container size is a production outage waiting to happen. |
| 2 | **Prepending is not appending** | `append([]T{v}, out...)` allocates a new array and copies everything already collected. |
| 3 | **Index arithmetic beats shuffling** | Writing straight to the mirrored index needs one allocation and one pass. |

## Hint

Count how many elements are copied on the `i`-th iteration.

## Validate

```bash
make verify
```
