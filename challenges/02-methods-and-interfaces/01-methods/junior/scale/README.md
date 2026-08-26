# Scale

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A physics engine scales velocity vectors by a time-step factor. The
operation mutates the vector in place.

## Task

Implement `Scale` on `*Vector` in [scale.go](scale.go):

1. Multiply both `X` and `Y` by `factor`.
2. Mutate the receiver — this must be a **pointer receiver**.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  v := Vector{3, 4}; v.Scale(2)
Output: v == Vector{6, 8}
```

**Example 2:**

```
Input:  v := Vector{1, -1}; v.Scale(0)
Output: v == Vector{0, 0}
```

**Example 3:**

```
Input:  v := Vector{10, 6}; v.Scale(0.5)
Output: v == Vector{5, 3}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer receiver** | `(v *Vector)` lets the method mutate the original struct. |
| 2 | **Mutation in-place** | Without `*`, changes are lost — the method works on a copy. |

## Hint

A pointer receiver `*Vector` means `v` is a pointer to the original struct.
Assign `v.X = v.X * factor` (Go auto-dereferences).

## Validate

```bash
make verify
```
