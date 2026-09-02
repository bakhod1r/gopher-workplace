# Generic List

**Level:** junior  
**Topic:** 03-generics

## Context

A widget list is driven by user input, so index arguments are untrusted and must never panic.

## Task

Implement the stub(s) in [listgen.go](listgen.go):

1. Implement `Append`, `At`, and `Len`.
2. `At` returns the zero value and `false` for any out-of-range index, including negatives.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Append(1); At(0)
Output: 1, true
```

**Example 2:**

```
Input:  At(5) on a 1-element list
Output: 0, false
```

**Example 3:**

```
Input:  At(-1)
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Generic types** | `type Stack[T any] struct { ... }` parameterises the type itself, not just a function. |
| 2 | **Total functions** | Returning `(T, bool)` turns a panicking index into a checked one. |
| 3 | **Zero value of `T`** | `var zero T` names the zero value of an unknown type. |

## Hint

Guard both ends: `i < 0` as well as `i >= len`.

## Validate

```bash
make verify
```
