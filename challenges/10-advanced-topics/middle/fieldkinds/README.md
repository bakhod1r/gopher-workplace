# Describe A Struct's Shape

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A schema documentation generator prints each config struct's fields and their shapes. Keeping the document in step with the code by hand lasted exactly one release.

## Task

Implement [fieldkinds.go](fieldkinds.go):

1. Return `"Name:kind"` for each exported field, in declaration order.
2. Report the field's kind, not its declared type name.
3. Return nil for a non-struct or a nil interface.

Replace the stub body in [fieldkinds.go](fieldkinds.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  FieldKinds(row{})
Output: [ID:int Name:string Tags:slice Ratio:float64]
```

_Explanation:_ `hidden` is unexported.

**Example 2:**

```
Input:  a field of type *row
Output: Ptr:ptr
```

_Explanation:_ The kind, not the type name.

**Example 3:**

```
Input:  FieldKinds(3)
Output: <nil>
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **StructField.Type** | Each field carries its own `reflect.Type`. |
| 2 | **Kind vs Type name** | `[]string` has kind slice; its type string is "[]string". |
| 3 | **Declaration order** | `Field(i)` walks the struct in source order. |

## Hint

`f.Type.Kind().String()` is the right-hand half of each entry.

## Validate

```bash
make verify
```
