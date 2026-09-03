# Capture Frames

**Level:** senior
**Topic:** 04-error-handling

## Context

A crash reporter records the function names leading to a failure so a support ticket carries the path, not just the message.

## Task

Implement `Frames` in [stackframes.go](stackframes.go):

1. Return up to `max` function names starting from the caller of `Frames`.
2. Return an empty slice when `max` is zero or negative.
3. Include only function names, not files or lines.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Frames(1)
Output: ["…TestFrames.func1"]
```

**Example 2:**

```
Input:  Frames(0)
Output: []
```

**Example 3:**

```
Input:  len(Frames(3)) <= 3
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **runtime.Callers** | Fills a slice of program counters. |
| 2 | **runtime.CallersFrames** | Turns PCs into named frames. |
| 3 | **Skip counts** | Skip runtime.Callers and Frames itself. |

## Hint

`runtime.Callers(skip, pc)` with skip 2 starts at the caller of this function; iterate frames with `More()`.

## Validate

```bash
make verify
```
