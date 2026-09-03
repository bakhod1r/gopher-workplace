# The Buffer That An Interface Sent To The Heap

**Level:** senior
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A hot encoder keeps its scratch buffer as a fixed-size local, exactly as the style guide says. The allocation profile still shows the buffer on the heap, once per call.

## Task

Fix the single planted bug in [ifaceescape.go](ifaceescape.go):

1. Render each value as decimal digits into the local scratch buffer.
2. Return the sum of the bytes written.
3. Fix the single bug so the function allocates nothing.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Checksum([]int{1})
Output: 49
```

_Explanation:_ '1' is byte 49.

**Example 2:**

```
Input:  Checksum([]int{12})
Output: 97
```

_Explanation:_ '1' + '2'.

**Example 3:**

```
Input:  Checksum(nil)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface arguments escape** | A value passed as an interface may be stored by the callee, so it goes to the heap. |
| 2 | **The receiver escapes too** | `&c` behind an interface makes `c` heap-allocated as well. |
| 3 | **Concrete calls keep the frame** | Direct code on a local slice keeps everything in the frame. |

## Hint

The buffer is a fixed-size local. Which line takes its address out of the function's sight?

## Validate

```bash
make verify
```
