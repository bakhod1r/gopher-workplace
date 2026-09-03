# List A Struct's Exported Fields

**Level:** junior
**Topic:** 10-advanced-topics / 03-reflection

## Context

A config auditor wants to report which settings a struct exposes. Hand-maintaining the list drifts from the struct within a release.

## Task

Implement [fieldnames.go](fieldnames.go):

1. Return the exported field names of `v`, in declaration order.
2. Skip unexported fields.
3. Return nil for anything that is not a struct, including nil and a pointer to a struct.

Replace the stub body in [fieldnames.go](fieldnames.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  FieldNames(user{})
Output: [Name Age]
```

_Explanation:_ `admin` is unexported.

**Example 2:**

```
Input:  FieldNames(&user{})
Output: <nil>
```

_Explanation:_ A pointer is not a struct.

**Example 3:**

```
Input:  FieldNames(3)
Output: <nil>
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **reflect.Type.NumField / Field** | Struct layout is walked by index, in declaration order. |
| 2 | **StructField.IsExported** | Export status is part of the field's metadata. |
| 3 | **Guarding the kind** | `NumField` panics on a non-struct, so check first. |

## Hint

`reflect.TypeOf(nil)` is nil — check that before you check the kind.

## Validate

```bash
make verify
```
