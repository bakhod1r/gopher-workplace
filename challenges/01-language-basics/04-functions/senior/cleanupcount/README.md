# Defers Fire at Function Exit

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

Deferred cleanups do NOT run at the end of each loop iteration — they all run
when the function returns. Decrementing eagerly inside the loop (the bug) drops
the open count immediately, so the peak never rises above 1.

## Task

Fix [cleanupcount.go](cleanupcount.go) so the drain is deferred to function exit.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  PeakThenDrain(3)
Output: 3
```

**Example 2:**

```
Input:  PeakThenDrain(0)
Output: 0
```

**Example 3:**

```
Input:  PeakThenDrain(1)
Output: 1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Defer timing** | Deferred calls fire at function return, not per iteration. |
| 2 | **Peak before drain** | All resources are open simultaneously. |
| 3 | **Loop-scheduled defers** | Each iteration schedules one, all run at exit. |

## Hint

Defer the decrement instead of doing it inline: replace `open--` with `defer func(){ open-- }()`.

## Validate

```bash
make verify
```
