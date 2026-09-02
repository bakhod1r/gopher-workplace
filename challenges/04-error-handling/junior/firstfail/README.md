# Index Of First Failure

**Level:** junior
**Topic:** 04-error-handling

## Context

A migration runs steps in order. The report names the position of the step that stopped it, so the operator can resume from there.

## Task

Implement `FirstFail` in [firstfail.go](firstfail.go):

1. Return the index of the first non-nil error and nil.
2. Return `-1` and `ErrNoFailure` when every entry is nil.
3. Treat an empty or nil slice as having no failure.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FirstFail([]error{nil, ErrStep})
Output: 1, nil
```

**Example 2:**

```
Input:  FirstFail([]error{nil, nil})
Output: -1, ErrNoFailure
```

**Example 3:**

```
Input:  FirstFail(nil)
Output: -1, ErrNoFailure
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Index from range** | `for i, err := range` gives position and value. |
| 2 | **Sentinel index** | `-1` marks "no position". |
| 3 | **Absence as an error** | Not finding anything can itself be reported. |

## Hint

"Nothing failed" is the success case for the migration but the error case for this lookup — return both parts consistently.

## Validate

```bash
make verify
```
