# How Many Fields Does This Have

**Level:** junior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A serialiser reports how much of each struct it can actually see. Unexported fields are invisible to it, and the counts explain a lot of confused bug reports.

## Task

Implement [fieldcount.go](fieldcount.go):

1. Return the total field count and the exported field count of `v`'s struct type.
2. Report 0, 0 for a non-struct, a pointer to a struct, or a nil interface.

Replace the stub body in [fieldcount.go](fieldcount.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  FieldCount(rec{})
Output: 3, 2
```

_Explanation:_ One unexported field.

**Example 2:**

```
Input:  FieldCount(struct{}{})
Output: 0, 0
```

**Example 3:**

```
Input:  FieldCount(&rec{})
Output: 0, 0
```

_Explanation:_ A pointer is not a struct.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type.NumField** | The field count is part of the type, not the value. |
| 2 | **StructField.IsExported** | Export status is metadata, not something to infer from the name. |
| 3 | **Guard the kind first** | `NumField` panics on anything but a struct. |

## Hint

Named results let you count in place without extra variables.

## Validate

```bash
make verify
```
