# First Failure

**Level:** junior
**Topic:** 04-error-handling

## Context

A deployment pipeline runs several checks and collects their results. The report shows the first thing that went wrong.

## Task

Implement `FirstError` in [firsterr.go](firsterr.go):

1. Return the first non-nil error in the slice.
2. Return nil when every entry is nil.
3. Return nil for an empty or nil slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FirstError([]error{nil, ErrB, ErrC})
Output: ErrB
```

**Example 2:**

```
Input:  FirstError([]error{nil, nil})
Output: nil
```

**Example 3:**

```
Input:  FirstError(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Slice iteration** | `for range` walks the slice in order. |
| 2 | **Early return** | Return as soon as the first failure is found. |
| 3 | **Nil slice** | Ranging over a nil slice runs zero iterations. |

## Hint

Order matters: return inside the loop, not after it.

## Validate

```bash
make verify
```
