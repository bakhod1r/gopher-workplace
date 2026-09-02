# Last Index

**Level:** middle  
**Topic:** 03-generics

## Context

A path splitter needs the final separator, not the first one, to isolate the file name.

## Task

Implement the stub(s) in [lastindexgen.go](lastindexgen.go):

1. Implement `LastIndex`, returning the index of the last matching element.
2. Return `-1` when no element matches.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  LastIndex([]int{7,1,7}, 7)
Output: 2
```

**Example 2:**

```
Input:  LastIndex([]int{7}, 7)
Output: 0
```

**Example 3:**

```
Input:  LastIndex([]int{1}, 7)
Output: -1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Backwards scans** | Iterating from the end returns the last match without tracking state. |
| 2 | **Early return** | The first hit going backwards is the last hit going forwards. |
| 3 | **Sentinel `-1`** | Reused convention for "not found". |

## Hint

Loop from `len(s)-1` down to `0` and return on the first match.

## Validate

```bash
make verify
```
