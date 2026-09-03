# A Sink That Sinks Nothing

**Level:** senior  
**Topic:** 11-performance-engineering

## Context

The benchmark has a package-level sink, the body assigns to it, and the numbers still look impossibly fast. Every piece of the pattern is present except the part that makes it work.

## Task

Fix the single planted bug in [sinkmissingbug.go](sinkmissingbug.go):

1. Find and fix the one bug so `Consume` actually stores its argument in `Sink`.
2. `Consume` must still return the value it replaced.
3. `SumTo` already works and must keep working.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Sink is 0; Consume(7)
Output: returns 0, Sink is now 7
```

**Example 2:**

```
Input:  then Consume(9)
Output: returns 7, Sink is now 9
```

**Example 3:**

```
Input:  Consume(SumTo(i)) for i in 0..4
Output: Sink is SumTo(4)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`_ = v` is not a store** | It is the idiom for deliberately ignoring a value, and it protects nothing. |
| 2 | **Dead-code elimination** | A result nothing observes can be compiled away along with the work that produced it. |
| 3 | **Observable means package-level** | The store has to be visible outside the function for the optimiser to keep it. |

## Hint

The line between the markers ignores the argument instead of storing it.

## Validate

```bash
make verify
```
