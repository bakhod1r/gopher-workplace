# Annotate Every Failure

**Level:** middle
**Topic:** 04-error-handling

## Context

A worker pool returns one error per job. The aggregator tags each with its job number before handing them on.

## Task

Implement `WrapAll` in [wrapall.go](wrapall.go):

1. Return a slice with each non-nil error wrapped as `"job <i>: <err>"`, using its original index.
2. Skip nil entries entirely.
3. Return nil when nothing failed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  WrapAll([]error{nil, ErrJob})
Output: ["job 1: job failed"]
```

**Example 2:**

```
Input:  WrapAll([]error{ErrJob})
Output: ["job 0: job failed"]
```

**Example 3:**

```
Input:  WrapAll(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Index-aware annotation** | The original position must survive filtering. |
| 2 | **Filtering and mapping** | One pass does both. |
| 3 | **errors.Is after wrapping** | Each wrapper still matches its cause. |

## Hint

The index comes from `range` over the input, not from the length of the output.

## Validate

```bash
make verify
```
