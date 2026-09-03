# Match With A Predicate

**Level:** staff
**Topic:** 04-error-handling

## Context

A diagnostic tool answers arbitrary questions about a failure — "does anything here mention a timeout?" — without knowing the concrete types.

## Task

Implement `Any` in [matchfunc.go](matchfunc.go):

1. Return true when `pred` returns true for any error in the tree.
2. Search both unwrap shapes, including the outermost error.
3. Return false for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Any(err, isTimeout)
Output: true
```

**Example 2:**

```
Input:  Any(nil, pred)
Output: false
```

**Example 3:**

```
Input:  Any(ErrA, func(error) bool { return true })
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Generalised traversal** | `errors.Is` is one predicate among many. |
| 2 | **Short-circuit search** | Stop at the first match. |
| 3 | **Tree walking** | Wrapped and joined children alike. |

## Hint

Test the node itself before recursing — the outermost error counts.

## Validate

```bash
make verify
```
