# Depth Of An Error Tree

**Level:** staff
**Topic:** 04-error-handling

## Context

A guard rejects pathological error values built by untrusted input, so the tree's depth must be measurable.

## Task

Implement `Depth` in [treedepth.go](treedepth.go):

1. Return `0` for a nil error and `1` for a leaf.
2. Return `1 + max(child depths)` for both unwrap shapes.
3. Treat a joined error's branches independently.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Depth(nil)
Output: 0
```

**Example 2:**

```
Input:  Depth(ErrA)
Output: 1
```

**Example 3:**

```
Input:  Depth(errors.Join(ErrA, fmt.Errorf("x: %w", ErrB)))
Output: 3
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Tree recursion** | Depth is the maximum over children. |
| 2 | **Two child shapes** | Both must be handled. |
| 3 | **Resource guards** | Unbounded structures are an attack surface. |

## Hint

A join adds a level of its own: the join node, then its deepest branch.

## Validate

```bash
make verify
```
