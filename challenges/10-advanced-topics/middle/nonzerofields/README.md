# Which Fields Were Actually Set

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A PATCH endpoint applies only the fields the caller sent. The struct cannot tell "omitted" from "set to zero" — but listing the non-zero ones is what the handler actually needs.

## Task

Implement [nonzerofields.go](nonzerofields.go):

1. Return the names of exported fields whose value is not the zero value.
2. Preserve declaration order; skip unexported fields.
3. Return nil for a non-struct or a nil interface.

Replace the stub body in [nonzerofields.go](nonzerofields.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  NonZero(patch{Name:"x", Count:3})
Output: [Name Count]
```

**Example 2:**

```
Input:  NonZero(patch{})
Output: <nil>
```

_Explanation:_ Everything is zero.

**Example 3:**

```
Input:  NonZero(patch{Tags: []string{}})
Output: [Tags]
```

_Explanation:_ An empty non-nil slice is not the zero slice.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value.IsZero** | Compares against the type's zero value, whatever the field's type. |
| 2 | **Type and Value in step** | Names come from the Type, values from the Value, at the same index. |
| 3 | **nil vs empty** | A nil slice is zero; an allocated empty one is not. |

## Hint

`rv.Field(i)` for the value, `rt.Field(i)` for the name.

## Validate

```bash
make verify
```
