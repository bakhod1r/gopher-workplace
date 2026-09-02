# Comparable Interface

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A ranking widget orders items it knows nothing about, using only a comparison method.

## Task

Implement the stub(s) in [compareifc.go](compareifc.go):

1. Implement `CompareTo` on `Score` — return a negative number, zero, or a positive number.
2. Implement `Max`, which returns the larger of two comparables (the first one when they are equal).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Score(5).CompareTo(Score(3))
Output: positive
```

**Example 2:**

```
Input:  Score(2).CompareTo(Score(2))
Output: 0
```

**Example 3:**

```
Input:  Max(Score(1), Score(9)).(Score)
Output: 9
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method on a defined non-struct type** | `type Score int` can satisfy an interface. |
| 2 | **Three-way comparison** | A single int encodes less / equal / greater. |
| 3 | **Type assertion** | Reused: the caller asserts the interface result back to `Score`. |

## Hint

`int(s) - int(o)` already has the right sign for small values, but write the explicit three-way form.

## Validate

```bash
make verify
```
