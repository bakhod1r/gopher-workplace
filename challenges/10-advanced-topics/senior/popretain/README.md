# The Popped Element That Never Left

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A worker pops jobs off a stack that is reused for the process's lifetime. Each job carries a kilobyte of payload, and the heap never comes back down after a burst.

## Task

Fix the single planted bug in [popretain.go](popretain.go):

1. Return the last element and the shortened slice.
2. An empty input returns nil and the slice unchanged.
3. Fix the single bug so the popped element stops being reachable through the array.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Pop([]*Job{a, b})
Output: b, [a]
```

**Example 2:**

```
Input:  s[1] after Pop(s)
Output: nil
```

_Explanation:_ The vacated slot is cleared.

**Example 3:**

```
Input:  Pop(nil)
Output: nil, []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Reslicing does not erase** | The element past the new length is still in the array. |
| 2 | **Long-lived containers leak** | A reused stack keeps every slot it ever filled. |
| 3 | **Clear before shortening** | Writing nil is the only way to release the reference. |

## Hint

The returned slice is right. What is still sitting at index `len(s)-1`?

## Validate

```bash
make verify
```
