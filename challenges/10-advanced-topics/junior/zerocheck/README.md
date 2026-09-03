# Is This Value The Zero Value

**Level:** junior
**Topic:** 10-advanced-topics / 03-reflection

## Context

A patch endpoint must tell "field omitted" from "field set to its zero value". The first cut compared against `nil` and treated every explicit zero as absent.

## Task

Implement [zerocheck.go](zerocheck.go):

1. Report whether `v` holds the zero value for its type.
2. A nil interface is zero.
3. A nil pointer is zero; a pointer to a zero struct is not.

Replace the stub body in [zerocheck.go](zerocheck.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  IsZero(0)
Output: true
```

**Example 2:**

```
Input:  IsZero(&pair{})
Output: false
```

_Explanation:_ The pointer itself is not nil.

**Example 3:**

```
Input:  IsZero([]int{})
Output: false
```

_Explanation:_ An empty non-nil slice is not the zero slice.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value.IsZero** | Compares against the type's zero value, whatever the type is. |
| 2 | **Value.IsValid** | False exactly when the interface held nothing. |
| 3 | **nil vs empty** | A nil slice is the zero value; an allocated empty one is not. |

## Hint

Calling a method on an invalid Value panics. Check validity first.

## Validate

```bash
make verify
```
