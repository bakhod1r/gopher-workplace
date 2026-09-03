# Flatten The Tree

**Level:** staff
**Topic:** 04-error-handling

## Context

Nested aggregations produce joins inside joins. The reporter needs the flat list of leaf failures.

## Task

Implement `Leaves` in [joinflatten.go](joinflatten.go):

1. Return every leaf error, depth first, left to right.
2. Descend into both `Unwrap() []error` and `Unwrap() error`.
3. Return nil for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Leaves(errors.Join(ErrA, errors.Join(ErrB)))
Output: [ErrA ErrB]
```

**Example 2:**

```
Input:  Leaves(fmt.Errorf("x: %w", ErrA))
Output: [ErrA]
```

**Example 3:**

```
Input:  Leaves(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Error trees** | Joins branch, wraps descend. |
| 2 | **Recursive traversal** | One function handles both shapes. |
| 3 | **Leaf definition** | An error that unwraps to nothing. |

## Hint

An error is a leaf when it implements neither unwrap shape — check both before recursing.

## Validate

```bash
make verify
```
