# Named Unit Types

**Level:** middle  
**Topic:** 03-generics

## Context

The codebase wraps floats in `Meters` and `Seconds` so they cannot be mixed up. Helpers must accept them without stripping the type.

## Task

Implement the stub(s) in [unitclamp.go](unitclamp.go):

1. Implement `ClampUnit`, keeping the caller's named type in the result.
2. The `~float64` constraint is what admits `Meters` and `Seconds`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ClampUnit(Meters(5), 0, 3)
Output: Meters(3)
```

**Example 2:**

```
Input:  ClampUnit(Seconds(-1), 0, 3)
Output: Seconds(0)
```

**Example 3:**

```
Input:  ClampUnit(2.0, 0, 3)
Output: 2.0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The `~` token** | `~int` admits every type whose underlying type is `int`. |
| 2 | **Type safety survives** | `Meters` and `Seconds` cannot be mixed in one call — the single `T` forbids it. |
| 3 | **No conversion needed** | Returning `T` keeps the unit; converting to `float64` would lose it. |

## Hint

One type parameter for all three arguments keeps units from mixing.

## Validate

```bash
make verify
```
