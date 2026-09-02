# The Wrong Winner On A Tie

**Level:** senior  
**Topic:** 03-generics

## Context

A leaderboard shows a different winner on every deploy whenever two players are level, and support keeps reopening the ticket.

## Task

Fix the single planted bug in [tiebug.go](tiebug.go):

1. Find and fix the single bug so the earliest element wins a tie.
2. The chosen key must still be the maximum.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  MaxBy([{a 3} {b 3}], score)
Output: {a 3}
```

**Example 2:**

```
Input:  MaxBy([{a 1} {b 3}], score)
Output: {b 3}
```

**Example 3:**

```
Input:  MaxBy([], score)
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Tie-breaking** | Strict versus non-strict comparison decides whether the first or last equal element wins. |
| 2 | **Documented behaviour is a contract** | "First wins" is testable; "either" is a bug report waiting to happen. |
| 3 | **One character** | `>=` versus `>` is the entire difference. |

## Hint

Read the comparison in the update branch.

## Validate

```bash
make verify
```
