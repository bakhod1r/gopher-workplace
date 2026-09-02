# Group Consecutive

**Level:** middle  
**Topic:** 03-generics

## Context

A run-length encoder needs the runs before it can encode them, and the input arrives unsorted.

## Task

Implement the stub(s) in [groupconsecutivegen.go](groupconsecutivegen.go):

1. Implement `GroupRuns`, returning runs of equal consecutive elements.
2. Return an empty result for an empty input; never emit an empty run.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  GroupRuns([]int{1,1,2,1})
Output: [][]int{{1,1},{2},{1}}
```

**Example 2:**

```
Input:  GroupRuns([]int{1,2})
Output: [][]int{{1},{2}}
```

**Example 3:**

```
Input:  GroupRuns([]int{})
Output: [][]int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Runs versus groups** | Non-adjacent equal values start a new run — this is not a grouping by value. |
| 2 | **Flush at the end** | The final run only lands after the loop. |
| 3 | **`comparable` for `==`** | Neighbour comparison is the only operation needed. |

## Hint

Same shape as `ChunkBy` with `==` as the boundary test.

## Validate

```bash
make verify
```
