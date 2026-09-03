# The Buffer Allocated Once Per Iteration

**Level:** middle
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A report renderer allocates a scratch buffer for every row it formats. With a million rows that is a million buffers the collector has to chase, all of them identical and all of them dead immediately.

## Task

Fix the single planted bug in [bufinloop.go](bufinloop.go):

1. Render each row as its values joined by `,`.
2. Fix the single bug so the scratch buffer is allocated once per call, not once per row.
3. The output must not change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Render([][]int{{1,2},{3}})
Output: ["1,2" "3"]
```

**Example 2:**

```
Input:  Render([][]int{{}})
Output: [""]
```

_Explanation:_ An empty row renders as an empty string.

**Example 3:**

```
Input:  64 rows
Output: about 66 allocations, not 130
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Hoisting allocations** | State that is the same every iteration belongs outside the loop. |
| 2 | **Resetting instead of reallocating** | `buf = buf[:0]` gives a clean buffer at no cost. |
| 3 | **string(buf) copies** | The per-row string is a real allocation and is supposed to be. |

## Hint

Only one line has to move. Something else has to be added where it used to be.

## Validate

```bash
make verify
```
