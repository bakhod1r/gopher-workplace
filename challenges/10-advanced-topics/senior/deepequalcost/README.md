# A Comparison That Should Not Reflect

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A config watcher compares the old and new settings on every poll. `reflect.DeepEqual` was the obvious spelling, and the watcher now allocates twice a second forever.

## Task

Fix the single planted bug in [deepequalcost.go](deepequalcost.go):

1. Report whether the two configs differ.
2. Fix the single bug so the comparison allocates nothing.
3. Every field must participate in the comparison.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Changed(Config{Retries:1}, Config{Retries:2})
Output: true
```

**Example 2:**

```
Input:  Changed(Config{}, Config{})
Output: false
```

**Example 3:**

```
Input:  allocations per call
Output: 0
```

_Explanation:_ `==` compiles to a direct comparison.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Comparable structs** | A struct of comparable fields supports `==` field by field. |
| 2 | **DeepEqual boxes its arguments** | Both operands become `any`, which allocates. |
| 3 | **Reflection is for unknown types** | Here the type is known at compile time. |
| 4 | **When DeepEqual is still right** | Slices, maps and functions are not comparable with `==`. |

## Hint

The struct has no slices, maps or functions in it. What does that make it?

## Validate

```bash
make verify
```
