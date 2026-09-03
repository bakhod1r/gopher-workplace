# First Matching Failure

**Level:** middle
**Topic:** 04-error-handling

## Context

A fan-out call returns one error per replica. The caller wants the first result that was a genuine timeout, ignoring other kinds of failure.

## Task

Implement `FirstMatch` in [findsentinel.go](findsentinel.go):

1. Return the first entry whose chain matches `target`.
2. Return nil when no entry matches.
3. Skip nil entries.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FirstMatch([]error{ErrOther, wrapped}, ErrTimeout)
Output: wrapped
```

**Example 2:**

```
Input:  FirstMatch([]error{ErrOther}, ErrTimeout)
Output: nil
```

**Example 3:**

```
Input:  FirstMatch(nil, ErrTimeout)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Search with a predicate** | The predicate is `errors.Is`. |
| 2 | **Returning the match itself** | The wrapper is more informative than the sentinel. |
| 3 | **Nil entries** | `errors.Is(nil, target)` is false, so they skip themselves. |

## Hint

Return the element from the slice, not the target you were searching for.

## Validate

```bash
make verify
```
