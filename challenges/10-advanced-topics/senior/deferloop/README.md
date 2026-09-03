# Deferred Cleanup That Waits For The Whole Loop

**Level:** senior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A batch importer opens a handle per row and defers the close. At row 60000 the process dies on "too many open files" — the deferred closes were all still waiting.

## Task

Fix the single planted bug in [deferloop.go](deferloop.go):

1. Return each item doubled, in order.
2. Call `release` for each item as soon as that item is processed, in input order.
3. Fix the single bug that keeps every item outstanding until the function returns.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Process([]int{1,2,3}, rel)
Output: [2 4 6]
```

**Example 2:**

```
Input:  release call order
Output: 1, 2, 3
```

_Explanation:_ Forward order, one per iteration.

**Example 3:**

```
Input:  outstanding items at any moment
Output: at most one
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **defer is function-scoped** | It runs when the function returns, not when the loop body ends. |
| 2 | **LIFO order** | Stacked defers run in reverse, which is the wrong order here. |
| 3 | **Resource lifetime** | Holding N resources instead of 1 is a memory and handle leak. |

## Hint

Where does a deferred call actually run? Not where you wrote it.

## Validate

```bash
make verify
```
