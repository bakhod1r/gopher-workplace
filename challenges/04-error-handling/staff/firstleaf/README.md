# The Original Failure

**Level:** staff
**Topic:** 04-error-handling

## Context

A postmortem wants the failure that actually happened first, with every annotation and aggregation layer stripped away.

## Task

Implement `Origin` in [firstleaf.go](firstleaf.go):

1. Return the leftmost leaf of the error tree.
2. Descend through both unwrap shapes.
3. Return nil for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Origin(fmt.Errorf("x: %w", ErrA))
Output: ErrA
```

**Example 2:**

```
Input:  Origin(errors.Join(ErrA, ErrB))
Output: ErrA
```

**Example 3:**

```
Input:  Origin(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Leftmost leaf** | Depth-first, first branch. |
| 2 | **Both shapes** | Wraps descend, joins pick the first branch. |
| 3 | **Root cause vs first cause** | In a tree they are not the same thing. |

## Hint

For a join, only the first branch matters — the rest of the tree is never visited.

## Validate

```bash
make verify
```
