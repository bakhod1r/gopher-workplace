# Most Recent Failure

**Level:** junior
**Topic:** 04-error-handling

## Context

A retry log keeps one entry per attempt. The alert shows the failure from the final attempt, not the first.

## Task

Implement `LastError` in [lasterr.go](lasterr.go):

1. Return the last non-nil error in the slice.
2. Return nil when every entry is nil.
3. Return nil for an empty or nil slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  LastError([]error{ErrA, nil, ErrB})
Output: ErrB
```

**Example 2:**

```
Input:  LastError([]error{ErrA, nil})
Output: ErrA
```

**Example 3:**

```
Input:  LastError(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Accumulator vs early return** | Keeping the newest match means no early exit. |
| 2 | **Loop variable scope** | The kept value must outlive the loop. |
| 3 | **Order of traversal** | Later entries overwrite earlier ones. |

## Hint

This is the mirror image of finding the first failure — do not return inside the loop.

## Validate

```bash
make verify
```
