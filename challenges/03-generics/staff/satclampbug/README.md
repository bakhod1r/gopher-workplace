# The Clamp That Runs Too Late

**Level:** staff  
**Topic:** 03-generics

## Context

A saturating accumulator over 8-bit audio samples is supposed to clip loud passages. Instead the loudest passages come back inverted, and the clipping branch never appears in coverage.

## Task

Fix the single planted bug in [satclampbug.go](satclampbug.go):

1. Find and fix the single bug so the limits are tested before the addition happens.
2. Additions that fit must return the exact sum for every member of the constraint.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  AddSat(int8(100), int8(100))
Output: 127
```

**Example 2:**

```
Input:  AddSat(int8(-100), int8(-100))
Output: -128
```

**Example 3:**

```
Input:  AddSat(3, 4)
Output: 7
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Overflow is not observable afterwards** | A wrapped sum is an ordinary value of `T`; nothing about it is out of range. |
| 2 | **Order of operations** | Doing the right steps in the wrong order is still a bug. |
| 3 | **Type sets** | A constraint names a *set* of types; the body must be correct for every member of it. |

## Hint

Can a value of type `T` ever be greater than the largest value of type `T`?

## Validate

```bash
make verify
```
