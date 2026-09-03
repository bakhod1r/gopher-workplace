# Look The Field Up Once, Not Once Per Row

**Level:** senior
**Topic:** 10-advanced-topics / 03-reflection

## Context

A reporting layer sums a column by name over a few million rows. The profile is dominated by `FieldByName`, which is doing a string comparison per field per row.

## Task

Implement [fieldindex.go](fieldindex.go):

1. Total the named int field over the slice of structs `rows`.
2. Resolve the field against the element type once, before the loop.
3. Return `ErrShape` for a non-slice, a non-struct element, or a field that is missing, unexported, or not an int.
4. An empty slice totals 0.

Replace the stub body in [fieldindex.go](fieldindex.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  SumColumn([]rec{{N:1},{N:2}}, "N")
Output: 3, nil
```

**Example 2:**

```
Input:  SumColumn([]rec{{}}, "Label")
Output: ErrShape
```

_Explanation:_ Not an int field.

**Example 3:**

```
Input:  SumColumn([]rec{}, "N")
Output: 0, nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type metadata is per type, not per value** | The field's position is the same for every row. |
| 2 | **StructField.Index / FieldByIndex** | The resolved position is an index path, usable without another name search. |
| 3 | **Validate once** | Shape errors are properties of the type, so they can be decided before the loop. |
| 4 | **Reflection's real cost** | Name resolution, not field access, is what dominates. |

## Hint

`FieldByName` returns a `StructField`. What is in its `Index`?

## Validate

```bash
make verify
```
