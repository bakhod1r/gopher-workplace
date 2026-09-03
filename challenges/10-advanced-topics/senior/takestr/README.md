# The String That Changed Underneath Its Owner

**Level:** senior
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A reader hands out strings built over its own reusable read buffer. Log lines come out interleaved, map keys stop matching themselves, and the strings in a cache change hours after they were stored.

## Task

Fix the single planted bug in [takestr.go](takestr.go):

1. Return the first `n` bytes of `buf` as a string, clamping `n` into `[0, len(buf)]`.
2. The result must survive later writes to `buf`.
3. Fix the single bug, and keep the cost at one allocation per call.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Take([]byte("hello"), 2)
Output: "he"
```

**Example 2:**

```
Input:  s := Take(buf, 5); copy(buf, "SECOND")
Output: s is unchanged
```

_Explanation:_ The string owns its bytes.

**Example 3:**

```
Input:  Take(nil, 3)
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Strings are assumed immutable** | Maps cache their hashes, and the compiler folds comparisons — a mutating string breaks both. |
| 2 | **unsafe.String does not copy** | It builds a header; ownership of the bytes is the caller's problem. |
| 3 | **Copy once, then wrap** | Allocating the bytes and wrapping them keeps the cost at one allocation. |

## Hint

The conversion is fine. The bytes it points at are the problem.

## Validate

```bash
make verify
```
