# Sum The Fields You Are Allowed To Read

**Level:** senior
**Topic:** 10-advanced-topics / 03-reflection

## Context

An audit helper totals the numeric fields of whatever record it is handed. It works in the unit tests and panics the moment a struct with a private field reaches it.

## Task

Fix the single planted bug in [skipunexported.go](skipunexported.go):

1. Total the exported int fields of `v`.
2. Skip unexported fields and fields of any other kind.
3. Return 0 for a nil interface or a non-struct.
4. Fix the single bug so no input can make the function panic.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  SumInts(mix{A:1, B:2, hidden:100})
Output: 3
```

_Explanation:_ `hidden` is unexported.

**Example 2:**

```
Input:  SumInts(mix{})
Output: 0
```

**Example 3:**

```
Input:  SumInts(&mix{})
Output: 0
```

_Explanation:_ A pointer is not a struct.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **CanInterface / IsExported** | An unexported field's Value cannot be boxed back into an `any`. |
| 2 | **Typed accessors** | `f.Int()` reads the value without going through `Interface`. |
| 3 | **Type and Value in step** | Export status lives on the `StructField`, not the `Value`. |

## Hint

Two things are wrong with one line. What does `Interface()` require, and what does an `int` field actually assert to?

## Validate

```bash
make verify
```
