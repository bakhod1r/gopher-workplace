# Report Which Fields Differ

**Level:** staff
**Topic:** 10-advanced-topics / 03-reflection

## Context

A config reload logs "settings changed" and nothing else. On-call needs to know which setting, and the struct has forty fields across four nested blocks.

## Task

Implement [fielddiff.go](fielddiff.go):

1. Return the dotted paths of the exported fields where `a` and `b` differ, in declaration order.
2. Descend into nested structs, joining names with `.`.
3. Skip unexported fields.
4. Return nil when the types differ or either value is a nil interface.

Replace the stub body in [fielddiff.go](fielddiff.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Diff(settings{Name:"x"}, settings{Name:"y"})
Output: [Name]
```

**Example 2:**

```
Input:  Diff(a, b) with only Limits.Soft differing
Output: [Limits.Soft]
```

_Explanation:_ Nested paths are dotted.

**Example 3:**

```
Input:  Diff(settings{}, limits{})
Output: <nil>
```

_Explanation:_ Different types are not comparable.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Parallel traversal** | Two Values of one type have identical field indices, so they walk in lockstep. |
| 2 | **Value.Equal** | Compares two Values of the same type without boxing them. |
| 3 | **Path accumulation** | The prefix is built on the way down, so leaves know their full name. |
| 4 | **Type identity** | `av.Type() != bv.Type()` is the only sound precondition for a field-by-field walk. |

## Hint

The recursive helper is given. Validate the pair, then let it walk.

## Validate

```bash
make verify
```
