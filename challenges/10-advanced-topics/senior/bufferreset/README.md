# The Shared Buffer Nobody Emptied

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A renderer keeps one package-level buffer behind a mutex. The first response is correct and every response after it contains the whole history of the process.

## Task

Fix the single planted bug in [bufferreset.go](bufferreset.go):

1. Render `vals` as decimal numbers joined by `-`.
2. Fix the single bug so each call starts from an empty buffer.
3. The mutex must stay — the buffer is shared.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Render([]int{1,2,3})
Output: "1-2-3"
```

**Example 2:**

```
Input:  200 calls of Render([]int{7})
Output: every call returns "7"
```

**Example 3:**

```
Input:  Render(nil)
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Shared state needs resetting** | A package-level buffer keeps whatever the last caller left. |
| 2 | **Buffer.Reset** | Empties it while keeping the memory it has grown. |
| 3 | **Unbounded growth** | Without the reset the buffer grows forever. |
| 4 | **The mutex is not the bug** | Serialisation was correct; the state was not. |

## Hint

The lock is right and the loop is right. What is the buffer's length when the loop starts?

## Validate

```bash
make verify
```
